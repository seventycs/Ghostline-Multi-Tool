package cmd

import (
    "ghostline/internal/roblox"
    "ghostline/internal/ui"
)

func init() {
    ui.RegisterCategory(ui.Category{
        Name: "Roblox",
        Items: []ui.Item{
            {"User Info", "retrieve detailed roblox account info", roblox.UserInfo},
            {"Cookie Info", "validate and check cookie data", roblox.CookieInfo},
            {"Cookie Login", "login via roblox cookie", roblox.CookieLogin},
            {"Group Info", "analyze details of target groups", roblox.GroupInfo},
            {"Asset Download", "download game assets and textures", roblox.AssetDownload},
            {"Name History", "track and display past usernames", roblox.NameHistory},
            {"Username Checker", "check availability of usernames", roblox.UsernameChecker},
            {"Cookie Refresher", "generate new cookie from existing", roblox.CookieRefresher},
            {"Game Info", "fetch game details and stats", roblox.GameInfo},
            {"Inventory Dump", "dump user inventory items", roblox.InventoryDump},
            {"Trade Scanner", "scan user trades for value", roblox.TradeScanner},
            {"Limited Sniper", "snipe limited items on sale", roblox.LimitedSniper},
            {"Group Cloner", "clone a roblox group structure", roblox.GroupCloner},
            {"Bot Follower", "mass follow a user with bots", roblox.BotFollower},
        },
    })
}