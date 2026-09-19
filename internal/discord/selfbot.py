#!/usr/bin/env python3
"""
ghostline selfbot v2
usage: python selfbot.py
commands are triggered with + prefix
"""
import discord
from discord.ext import commands
import asyncio
import aiohttp
import json
import os
import random
import string
import time
from datetime import datetime

TOKEN_FILE = os.path.join(os.environ.get("TEMP", "/tmp"), ".selfbot_token")
PREFIX = "+"

# no intents needed for selfbots with discord.py-self
bot = commands.Bot(command_prefix=PREFIX, self_bot=True)

def load_token():
    if os.path.exists(TOKEN_FILE):
        with open(TOKEN_FILE) as f:
            return f.read().strip()
    return None

def save_token(t):
    with open(TOKEN_FILE, "w") as f:
        f.write(t)

def embed_color():
    return 0x00ffcc

def clr():
    import colorama
    colorama.init()

@bot.event
async def on_ready():
    print(f"[ghostline selfbot] logged in as {bot.user}")
    print(f"  id:     {bot.user.id}")
    print(f"  guilds: {len(bot.guilds)}")
    try:
        await bot.change_presence(
            status=discord.Status.online,
            activity=discord.Game(name="ghostline · +help")
        )
    except Exception:
        pass

@bot.event
async def on_command_error(ctx, error):
    try:
        await ctx.message.delete()
    except Exception:
        pass
    try:
        await ctx.send(f"`error: {error}`", delete_after=5)
    except Exception:
        pass

async def ghost_reply(ctx, text):
    """send a reply that auto-deletes after 5s, delete the invoking command"""
    try:
        await ctx.message.delete()
    except Exception:
        pass
    try:
        return await ctx.send(text, delete_after=5)
    except Exception:
        return None

def make_embed(title, description=None, color=None):
    e = discord.Embed(
        title=title,
        description=description,
        color=color if color is not None else embed_color(),
        timestamp=datetime.utcnow(),
    )
    e.set_footer(text="ghostline selfbot v2 · +help")
    return e

# ==================================================================
# GENERAL
# ==================================================================

@bot.command(name="help")
async def help_cmd(ctx):
    fields = [
        ("⚙  utility",
         "`+ping` `+userinfo [@user]` `+avatar [@user]` `+banner [@user]` "
         "`+serverinfo` `+roles` `+tokeninfo` `+emojis` `+invites`"),
        ("💬  messages",
         "`+say <text>` `+embed <title> | <desc>` `+purge <n>` `+spam <n> <msg>` "
         "`+massdm <msg>` `+pin` `+unpin` `+react <emoji>`"),
        ("👑  moderation",
         "`+kick <@user>` `+ban <@user>` `+mute <@user>` `+unmute <@user>` "
         "`+nuke` `+clone`"),
        ("🎭  status",
         "`+status <online|idle|dnd|invisible>` `+playing <text>` `+streaming <text>` "
         "`+listening <text>` `+watching <text>` `+autonick <text>`"),
        ("🔊  voice",
         "`+joinvc <channel_id>` `+leavevc` `+vcspam <channel_id>`"),
        ("👻  ghost",
         "`+ghost <secs>` — deletes your messages in the last n seconds "
         "`+ghostdm <@user> <msg>` `+clear`"),
        ("🎲  fun",
         "`+8ball <question>` `+coinflip` `+dice [n]` `+meme`"),
    ]
    e = make_embed("👻 ghostline selfbot · commands", "prefix: `+`")
    for n, v in fields:
        e.add_field(name=n, value=v, inline=False)
    try:
        await ctx.message.delete()
    except Exception:
        pass
    await ctx.send(embed=e)

@bot.command(name="ping")
async def ping_cmd(ctx):
    await ghost_reply(ctx, f"`{round(bot.latency * 1000)}ms`")

@bot.command(name="userinfo")
async def userinfo_cmd(ctx, user: discord.User = None):
    user = user or ctx.author
    e = make_embed(f"👤 {user}")
    e.add_field(name="id", value=str(user.id), inline=True)
    e.add_field(name="bot", value=str(user.bot), inline=True)
    e.add_field(name="created", value=user.created_at.strftime("%Y-%m-%d"), inline=True)
    e.set_thumbnail(url=user.display_avatar.url)
    try:
        await ctx.message.delete()
    except Exception:
        pass
    await ctx.send(embed=e)

@bot.command(name="avatar")
async def avatar_cmd(ctx, user: discord.User = None):
    user = user or ctx.author
    try:
        await ctx.message.delete()
    except Exception:
        pass
    await ctx.send(user.display_avatar.url)

@bot.command(name="banner")
async def banner_cmd(ctx, user: discord.User = None):
    user = user or ctx.author
    try:
        u = await bot.fetch_user(user.id)
    except Exception:
        u = user
    if not getattr(u, "banner", None):
        await ghost_reply(ctx, "`no banner`")
        return
    try:
        await ctx.message.delete()
    except Exception:
        pass
    await ctx.send(u.banner.url)

@bot.command(name="serverinfo")
async def serverinfo_cmd(ctx):
    g = ctx.guild
    if not g:
        await ghost_reply(ctx, "`not in a guild`")
        return
    e = make_embed(f"🏠 {g.name}")
    e.add_field(name="id", value=str(g.id), inline=True)
    e.add_field(name="owner", value=str(g.owner), inline=True)
    e.add_field(name="members", value=str(g.member_count), inline=True)
    e.add_field(name="channels", value=str(len(g.channels)), inline=True)
    e.add_field(name="roles", value=str(len(g.roles)), inline=True)
    e.add_field(name="created", value=g.created_at.strftime("%Y-%m-%d"), inline=True)
    if g.icon:
        e.set_thumbnail(url=g.icon.url)
    try:
        await ctx.message.delete()
    except Exception:
        pass
    await ctx.send(embed=e)

@bot.command(name="roles")
async def roles_cmd(ctx):
    if not ctx.guild:
        return
    names = [r.name for r in ctx.guild.roles]
    txt = ", ".join(names[:50])
    e = make_embed(f"🎭 roles ({len(names)})", txt[:4000])
    try:
        await ctx.message.delete()
    except Exception:
        pass
    await ctx.send(embed=e)

@bot.command(name="emojis")
async def emojis_cmd(ctx):
    if not ctx.guild:
        return
    txt = " ".join(str(e) for e in ctx.guild.emojis[:50])
    try:
        await ctx.message.delete()
    except Exception:
        pass
    await ctx.send(txt or "`no emojis`")

@bot.command(name="invites")
async def invites_cmd(ctx):
    if not ctx.guild:
        return
    try:
        invites = await ctx.guild.invites()
    except Exception:
        invites = []
    lines = [f"`{i.code}` — {i.uses} uses" for i in invites[:20]]
    try:
        await ctx.message.delete()
    except Exception:
        pass
    await ctx.send(embed=make_embed(f"✉ invites ({len(invites)})", "\n".join(lines) or "none"))

@bot.command(name="tokeninfo")
async def tokeninfo_cmd(ctx):
    async with aiohttp.ClientSession() as s:
        async with s.get("https://discord.com/api/v10/users/@me",
                         headers={"Authorization": bot.http.token}) as r:
            data = await r.json()
    e = make_embed("🔑 token info")
    for k, v in data.items():
        e.add_field(name=k, value=str(v)[:1024], inline=False)
    try:
        await ctx.message.delete()
    except Exception:
        pass
    await ctx.send(embed=e)

# ==================================================================
# MESSAGES
# ==================================================================

@bot.command(name="say")
async def say_cmd(ctx, *, text: str = ""):
    try:
        await ctx.message.delete()
    except Exception:
        pass
    if text:
        await ctx.send(text)

@bot.command(name="embed")
async def embed_cmd(ctx, *, args: str = ""):
    if "|" in args:
        title, desc = args.split("|", 1)
    else:
        title, desc = args, None
    try:
        await ctx.message.delete()
    except Exception:
        pass
    await ctx.send(embed=make_embed(title.strip(), desc.strip() if desc else None))

@bot.command(name="purge")
async def purge_cmd(ctx, n: int = 10):
    try:
        await ctx.message.delete()
    except Exception:
        pass
    try:
        deleted = await ctx.channel.purge(limit=n)
        m = await ctx.send(f"purged {len(deleted)}")
        await asyncio.sleep(3)
        await m.delete()
    except Exception:
        pass

@bot.command(name="spam")
async def spam_cmd(ctx, n: int = 5, *, msg: str = "ghostline"):
    try:
        await ctx.message.delete()
    except Exception:
        pass
    for _ in range(n):
        try:
            await ctx.send(msg)
        except Exception:
            break
        await asyncio.sleep(0.6)

@bot.command(name="massdm")
async def massdm_cmd(ctx, *, msg: str = "ghostline"):
    try:
        await ctx.message.delete()
    except Exception:
        pass
    if not ctx.guild:
        return
    sent = 0
    for m in ctx.guild.members:
        if m.bot or m == bot.user:
            continue
        try:
            await m.send(msg)
            sent += 1
            await asyncio.sleep(1.2)
        except Exception:
            pass
    m = await ctx.send(f"sent to {sent}")
    await asyncio.sleep(3)
    await m.delete()

@bot.command(name="pin")
async def pin_cmd(ctx):
    try:
        if ctx.message.reference:
            msg = await ctx.channel.fetch_message(ctx.message.reference.message_id)
            await msg.pin()
    except Exception:
        pass
    try:
        await ctx.message.delete()
    except Exception:
        pass

@bot.command(name="unpin")
async def unpin_cmd(ctx):
    try:
        if ctx.message.reference:
            msg = await ctx.channel.fetch_message(ctx.message.reference.message_id)
            await msg.unpin()
    except Exception:
        pass
    try:
        await ctx.message.delete()
    except Exception:
        pass

@bot.command(name="react")
async def react_cmd(ctx, emoji: str = "👻"):
    try:
        if ctx.message.reference:
            msg = await ctx.channel.fetch_message(ctx.message.reference.message_id)
            await msg.add_reaction(emoji)
    except Exception:
        pass
    try:
        await ctx.message.delete()
    except Exception:
        pass

# ==================================================================
# MODERATION
# ==================================================================

@bot.command(name="kick")
async def kick_cmd(ctx, member: discord.Member = None):
    if not member or not ctx.guild:
        return
    try:
        await member.kick()
        await ghost_reply(ctx, f"kicked {member}")
    except Exception as e:
        await ghost_reply(ctx, f"`error: {e}`")

@bot.command(name="ban")
async def ban_cmd(ctx, member: discord.Member = None):
    if not member or not ctx.guild:
        return
    try:
        await member.ban()
        await ghost_reply(ctx, f"banned {member}")
    except Exception as e:
        await ghost_reply(ctx, f"`error: {e}`")

@bot.command(name="mute")
async def mute_cmd(ctx, member: discord.Member = None):
    if not member or not ctx.guild:
        return
    try:
        await member.edit(mute=True)
        await ghost_reply(ctx, f"muted {member}")
    except Exception as e:
        await ghost_reply(ctx, f"`error: {e}`")

@bot.command(name="unmute")
async def unmute_cmd(ctx, member: discord.Member = None):
    if not member or not ctx.guild:
        return
    try:
        await member.edit(mute=False)
        await ghost_reply(ctx, f"unmuted {member}")
    except Exception as e:
        await ghost_reply(ctx, f"`error: {e}`")

@bot.command(name="nuke")
async def nuke_cmd(ctx):
    if not ctx.guild:
        return
    try:
        await ctx.message.delete()
    except Exception:
        pass
    for ch in list(ctx.guild.channels):
        try:
            await ch.delete()
        except Exception:
            pass
    for _ in range(20):
        try:
            await ctx.guild.create_text_channel(name="ghostline-nuked")
        except Exception:
            pass

@bot.command(name="clone")
async def clone_cmd(ctx):
    """clones the current guild's channels+roles to another guild id stored in the message"""
    await ghost_reply(ctx, "`use +clone <target_guild_id>` (not implemented in v2 — manual)`")

# ==================================================================
# STATUS
# ==================================================================

@bot.command(name="status")
async def status_cmd(ctx, status: str = "online"):
    statuses = {
        "online":    discord.Status.online,
        "idle":      discord.Status.idle,
        "dnd":       discord.Status.dnd,
        "invisible": discord.Status.invisible,
    }
    await bot.change_presence(status=statuses.get(status, discord.Status.online))
    await ghost_reply(ctx, f"`status -> {status}`")

@bot.command(name="playing")
async def playing_cmd(ctx, *, text: str = "ghostline"):
    await bot.change_presence(activity=discord.Game(name=text))
    await ghost_reply(ctx, f"`playing {text}`")

@bot.command(name="streaming")
async def streaming_cmd(ctx, *, text: str = "ghostline"):
    await bot.change_presence(activity=discord.Streaming(name=text, url="https://twitch.tv/ghostline"))
    await ghost_reply(ctx, "`streaming set`")

@bot.command(name="listening")
async def listening_cmd(ctx, *, text: str = "ghostline"):
    await bot.change_presence(activity=discord.Activity(type=discord.ActivityType.listening, name=text))
    await ghost_reply(ctx, f"`listening {text}`")

@bot.command(name="watching")
async def watching_cmd(ctx, *, text: str = "ghostline"):
    await bot.change_presence(activity=discord.Activity(type=discord.ActivityType.watching, name=text))
    await ghost_reply(ctx, f"`watching {text}`")

@bot.command(name="autonick")
async def autonick_cmd(ctx, *, name: str = "ghostline"):
    if not ctx.guild:
        return
    try:
        await ctx.guild.me.edit(nick=name)
        await ghost_reply(ctx, f"`nick -> {name}`")
    except Exception as e:
        await ghost_reply(ctx, f"`error: {e}`")

# ==================================================================
# VOICE
# ==================================================================

@bot.command(name="joinvc")
async def joinvc_cmd(ctx, channel_id: int = 0):
    ch = bot.get_channel(channel_id) if channel_id else (ctx.author.voice.channel if ctx.author.voice else None)
    if not ch:
        await ghost_reply(ctx, "`no channel`")
        return
    try:
        await ch.connect()
        await ghost_reply(ctx, "`joined vc`")
    except Exception as e:
        await ghost_reply(ctx, f"`error: {e}`")

@bot.command(name="leavevc")
async def leavevc_cmd(ctx):
    try:
        if ctx.voice_client:
            await ctx.voice_client.disconnect()
        await ghost_reply(ctx, "`left vc`")
    except Exception as e:
        await ghost_reply(ctx, f"`error: {e}`")

@bot.command(name="vcspam")
async def vcspam_cmd(ctx, channel_id: int = 0):
    if not channel_id or not ctx.guild:
        return
    try:
        for _ in range(10):
            vc = await bot.get_channel(channel_id).connect()
            await vc.disconnect()
            await asyncio.sleep(0.4)
        await ghost_reply(ctx, "`vc spam done`")
    except Exception as e:
        await ghost_reply(ctx, f"`error: {e}`")

# ==================================================================
# GHOST
# ==================================================================

@bot.command(name="ghost")
async def ghost_cmd(ctx, seconds: int = 10):
    """deletes your own messages in the current channel from the last N seconds"""
    try:
        await ctx.message.delete()
    except Exception:
        pass
    cutoff = datetime.utcnow().timestamp() - seconds
    deleted = 0
    try:
        async for m in ctx.channel.history(limit=200):
            if m.author == bot.user and m.created_at.timestamp() > cutoff:
                try:
                    await m.delete()
                    deleted += 1
                except Exception:
                    pass
    except Exception:
        pass
    try:
        x = await ctx.send(f"`ghosted {deleted}`")
        await asyncio.sleep(3)
        await x.delete()
    except Exception:
        pass

@bot.command(name="ghostdm")
async def ghostdm_cmd(ctx, user: discord.User = None, *, msg: str = ""):
    if not user:
        return
    try:
        m = await user.send(msg)
        await asyncio.sleep(5)
        await m.delete()
    except Exception:
        pass
    try:
        await ctx.message.delete()
    except Exception:
        pass

@bot.command(name="clear")
async def clear_cmd(ctx):
    """deletes ALL your messages in the current channel"""
    deleted = 0
    try:
        async for m in ctx.channel.history(limit=500):
            if m.author == bot.user:
                try:
                    await m.delete()
                    deleted += 1
                except Exception:
                    pass
    except Exception:
        pass
    try:
        x = await ctx.send(f"`cleared {deleted}`")
        await asyncio.sleep(3)
        await x.delete()
    except Exception:
        pass

# ==================================================================
# FUN
# ==================================================================

@bot.command(name="8ball")
async def eightball_cmd(ctx, *, question: str = ""):
    answers = [
        "yes", "no", "maybe", "definitely", "absolutely not",
        "ask again later", "signs point to yes", "my sources say no",
        "without a doubt", "don't count on it",
    ]
    await ghost_reply(ctx, f"`{random.choice(answers)}`")

@bot.command(name="coinflip")
async def coinflip_cmd(ctx):
    await ghost_reply(ctx, f"`{random.choice(['heads', 'tails'])}`")

@bot.command(name="dice")
async def dice_cmd(ctx, sides: int = 6):
    await ghost_reply(ctx, f"`{random.randint(1, sides)}`")

@bot.command(name="meme")
async def meme_cmd(ctx):
    try:
        async with aiohttp.ClientSession() as s:
            async with s.get("https://meme-api.com/gimme") as r:
                data = await r.json()
        await ctx.send(data.get("url", "no meme"))
        try:
            await ctx.message.delete()
        except Exception:
            pass
    except Exception as e:
        await ghost_reply(ctx, f"`error: {e}`")

# ==================================================================
# ENTRY
# ==================================================================

if __name__ == "__main__":
    tok = load_token()
    if not tok:
        try:
            tok = input("enter discord token: ").strip()
            save_token(tok)
        except (KeyboardInterrupt, EOFError):
            print("cancelled.")
            raise SystemExit
    try:
        bot.run(tok, bot=False)
    except Exception as e:
        print(f"selfbot crashed: {e}")
        input("press enter...")