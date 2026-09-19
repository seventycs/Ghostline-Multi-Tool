package cmd

import (

    "ghostline/internal/discord"
    "ghostline/internal/ui"
)

func init() {
    ui.RegisterCategory(ui.Category{
        Name: "Discord",
        Items: []ui.Item{
            {"Webhook Tools", "spam or delete discord webhooks", discord.WebhookTools},
            {"Token Tools", "account nuker, token login, info", discord.TokenTools},
            {"Server Info", "retrieve guild information", discord.ServerInfo},
            {"Bot Invite Gen", "generate admin bot invite links", discord.BotInviteGen},
            {"Selfbot", "launch advanced selfbot (python)", discord.Selfbot},
            {"Server Cloner", "clone a discord server using a token", discord.ServerCloner},
            {"Nuke Bot", "advanced server destruction console", discord.NukeBot},
            {"Username Checker", "check availability of usernames", discord.UsernameChecker},
            {"Token Checker", "validate tokens in bulk", discord.TokenChecker},
            {"Guild Backup", "backup full server structure", discord.GuildBackup},
            {"Message Logger", "log messages from a channel", discord.MessageLogger},
            {"Voice Spam", "join and spam voice channels", discord.VoiceSpam},
        },
    })
}