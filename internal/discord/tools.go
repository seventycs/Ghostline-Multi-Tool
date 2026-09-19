package discord

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "os/exec"
    "strings"
    "sync"
    "time"

    "ghostline/internal/ui"
)

const apiBase = "https://discord.com/api/v10"

func WebhookTools() {
    ui.Cyan("enter webhook URL:")
    var hook string
    fmt.Scanln(&hook)
    hook = strings.TrimSpace(hook)

    for {
        ui.Clear()
        ui.Green("webhook tools")
        fmt.Println("  [1] send message")
        fmt.Println("  [2] spam messages")
        fmt.Println("  [3] delete webhook")
        fmt.Println("  [4] get webhook info")
        fmt.Println("  [99] back")
        fmt.Print("\nchoice: ")
        var c string
        fmt.Scanln(&c)
        switch c {
        case "1":
            fmt.Print("message: ")
            var m string
            fmt.Scanln(&m)
            sendWebhook(hook, m, 1, 0)
        case "2":
            fmt.Print("message: ")
            var m string
            fmt.Scanln(&m)
            fmt.Print("count: ")
            var n int
            fmt.Scanln(&n)
            fmt.Print("delay ms: ")
            var d int
            fmt.Scanln(&d)
            sendWebhook(hook, m, n, time.Duration(d)*time.Millisecond)
        case "3":
            deleteWebhook(hook)
        case "4":
            getWebhookInfo(hook)
        case "99":
            return
        }
        fmt.Print("\nenter to continue...")
        fmt.Scanln()
    }
}

func sendWebhook(hook, msg string, count int, delay time.Duration) {
    var wg sync.WaitGroup
    for i := 0; i < count; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            body, _ := json.Marshal(map[string]string{"content": msg})
            resp, err := http.Post(hook, "application/json", bytes.NewReader(body))
            if err != nil {
                ui.Red("fail: " + err.Error())
                return
            }
            resp.Body.Close()
            fmt.Printf("  sent %d/%d [%d]\n", i+1, count, resp.StatusCode)
        }()
        if delay > 0 {
            time.Sleep(delay)
        }
    }
    wg.Wait()
}

func deleteWebhook(hook string) {
    req, _ := http.NewRequest("DELETE", hook, nil)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        ui.Red(err.Error())
        return
    }
    resp.Body.Close()
    ui.Green("webhook deleted.")
}

func getWebhookInfo(hook string) {
    resp, err := http.Get(hook)
    if err != nil {
        ui.Red(err.Error())
        return
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
}

func TokenTools() {
    ui.Cyan("enter discord token:")
    var tok string
    fmt.Scanln(&tok)
    tok = strings.TrimSpace(tok)

    for {
        ui.Clear()
        ui.Green("token tools")
        fmt.Println("  [1] validate token")
        fmt.Println("  [2] get account info")
        fmt.Println("  [3] list guilds")
        fmt.Println("  [4] nuke account (delete DMs, leave servers)")
        fmt.Println("  [5] change status")
        fmt.Println("  [6] rotate token")
        fmt.Println("  [99] back")
        fmt.Print("\nchoice: ")
        var c string
        fmt.Scanln(&c)
        switch c {
        case "1":
            validateToken(tok)
        case "2":
            accountInfo(tok)
        case "3":
            listGuilds(tok)
        case "4":
            nukeAccount(tok)
        case "5":
            setStatus(tok)
        case "6":
            rotateToken(tok)
        case "99":
            return
        }
        fmt.Print("\nenter to continue...")
        fmt.Scanln()
    }
}

func discordReq(method, path, token string, body interface{}) (*http.Response, error) {
    var buf io.Reader
    if body != nil {
        b, _ := json.Marshal(body)
        buf = bytes.NewReader(b)
    }
    req, _ := http.NewRequest(method, apiBase+path, buf)
    req.Header.Set("Authorization", token)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
    return http.DefaultClient.Do(req)
}

func validateToken(tok string) {
    resp, err := discordReq("GET", "/users/@me", tok, nil)
    if err != nil {
        ui.Red(err.Error())
        return
    }
    defer resp.Body.Close()
    if resp.StatusCode == 200 {
        ui.Green("token valid.")
    } else {
        ui.Red(fmt.Sprintf("invalid token (status %d)", resp.StatusCode))
    }
}

func accountInfo(tok string) {
    resp, err := discordReq("GET", "/users/@me", tok, nil)
    if err != nil {
        ui.Red(err.Error())
        return
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    var data map[string]interface{}
    json.Unmarshal(body, &data)
    for k, v := range data {
        fmt.Printf("  %s: %v\n", ui.Green(k), v)
    }
}

func listGuilds(tok string) {
    resp, err := discordReq("GET", "/users/@me/guilds", tok, nil)
    if err != nil {
        ui.Red(err.Error())
        return
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    var guilds []map[string]interface{}
    json.Unmarshal(body, &guilds)
    for _, g := range guilds {
        fmt.Printf("  %v — %v\n", ui.Green(fmt.Sprintf("%v", g["id"])), g["name"])
    }
}

func nukeAccount(tok string) {
    ui.Yellow("warning: this will leave all servers and delete DMs.")
    fmt.Print("confirm (yes): ")
    var c string
    fmt.Scanln(&c)
    if c != "yes" {
        return
    }
    resp, _ := discordReq("GET", "/users/@me/guilds", tok, nil)
    var guilds []map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&guilds)
    resp.Body.Close()
    for _, g := range guilds {
        id := fmt.Sprintf("%v", g["id"])
        if id == "" { continue }
        discordReq("DELETE", "/users/@me/guilds/"+id, tok, nil)
        fmt.Println("  left " + id)
    }
}

func setStatus(tok string) {
    statuses := []string{"online", "idle", "dnd", "invisible"}
    fmt.Println("statuses: " + strings.Join(statuses, ", "))
    fmt.Print("pick: ")
    var s string
    fmt.Scanln(&s)
    body := map[string]interface{}{"status": s}
    discordReq("PATCH", "/users/@me/settings", tok, body)
}

func rotateToken(tok string) {
    resp, _ := discordReq("POST", "/users/@me/change_password", tok, nil)
    if resp != nil {
        resp.Body.Close()
    }
    ui.Yellow("token rotation usually requires password confirmation")
}

func ServerInfo() {
    ui.Cyan("enter invite code or server ID:")
    var q string
    fmt.Scanln(&q)
    q = strings.TrimSpace(q)
    resp, err := http.Get("https://discord.com/api/v10/invites/" + q + "?with_counts=true")
    if err != nil {
        ui.Red(err.Error())
        return
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    var data map[string]interface{}
    json.Unmarshal(body, &data)
    fmt.Println()
    for k, v := range data {
        fmt.Printf("  %s: %v\n", ui.Green(k), v)
    }
}

func BotInviteGen() {
    ui.Cyan("enter bot client ID:")
    var id string
    fmt.Scanln(&id)
    perms := "8" // admin
    fmt.Print("permissions (default 8=admin): ")
    var p string
    fmt.Scanln(&p)
    if p != "" {
        perms = p
    }
    url := fmt.Sprintf("https://discord.com/api/oauth2/authorize?client_id=%s&permissions=%s&scope=bot%%20applications.commands", id, perms)
    fmt.Println(ui.Green(url))
}

func ServerCloner() {
    ui.Cyan("enter source guild ID:")
    var src string
    fmt.Scanln(&src)
    ui.Cyan("enter target guild ID:")
    var dst string
    fmt.Scanln(&dst)
    ui.Cyan("enter your token:")
    var tok string
    fmt.Scanln(&tok)
    ui.Yellow("cloning roles, channels, categories...")

    // fetch source guild
    resp, err := discordReq("GET", "/guilds/"+src, tok, nil)
    if err != nil {
        ui.Red(err.Error())
        return
    }
    var guild map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&guild)
    resp.Body.Close()

    // create roles
    if roles, ok := guild["roles"].([]interface{}); ok {
        for _, r := range roles {
            role := r.(map[string]interface{})
            if role["name"] == "@everyone" {
                continue
            }
            body := map[string]interface{}{
                "name":  role["name"],
                "color": role["color"],
                "hoist": role["hoist"],
                "mentionable": role["mentionable"],
            }
            discordReq("POST", "/guilds/"+dst+"/roles", tok, body)
            fmt.Println("  role cloned: " + fmt.Sprintf("%v", role["name"]))
        }
    }

    // create channels
    if channels, ok := guild["channels"].([]interface{}); ok {
        for _, c := range channels {
            ch := c.(map[string]interface{})
            body := map[string]interface{}{
                "name": ch["name"],
                "type": ch["type"],
            }
            if p, ok := ch["parent_id"]; ok {
                body["parent_id"] = p
            }
            discordReq("POST", "/guilds/"+dst+"/channels", tok, body)
            fmt.Println("  channel cloned: " + fmt.Sprintf("%v", ch["name"]))
        }
    }
    ui.Green("clone complete.")
}

func NukeBot() {
    ui.Cyan("enter your token:")
    var tok string
    fmt.Scanln(&tok)
    ui.Cyan("enter guild ID:")
    var gid string
    fmt.Scanln(&gid)
    ui.Yellow("this will delete ALL channels in the guild. confirm (yes):")
    var c string
    fmt.Scanln(&c)
    if c != "yes" {
        return
    }
    resp, _ := discordReq("GET", "/guilds/"+gid+"/channels", tok, nil)
    var chans []map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&chans)
    resp.Body.Close()
    for _, ch := range chans {
        id := fmt.Sprintf("%v", ch["id"])
        discordReq("DELETE", "/channels/"+id, tok, nil)
        fmt.Println("  deleted " + id)
    }
    // spam new channels
    for i := 0; i < 50; i++ {
        body := map[string]interface{}{"name": "ghostline-nuked", "type": 0}
        discordReq("POST", "/guilds/"+gid+"/channels", tok, body)
    }
    ui.Red("nuke complete.")
}

func UsernameChecker() {
    ui.Cyan("enter usernames (comma separated):")
    var line string
    fmt.Scanln(&line)
    names := strings.Split(line, ",")
    for _, n := range names {
        n = strings.TrimSpace(n)
        body := map[string]interface{}{
            "username": n,
            "password": "Ghostline!Check123",
            "email":    "",
        }
        // note: discord API requires captcha for register now
        b, _ := json.Marshal(body)
        resp, err := http.Post("https://discord.com/api/v9/auth/register", "application/json", bytes.NewReader(b))
        if err != nil {
            ui.Red(n + ": " + err.Error())
            continue
        }
        resp.Body.Close()
        fmt.Printf("  %s: status %d\n", n, resp.StatusCode)
    }
}

func TokenChecker() {
    ui.Cyan("enter tokens (one per line, blank to finish):")
    for {
        var t string
        fmt.Scanln(&t)
        if t == "" {
            break
        }
        resp, err := discordReq("GET", "/users/@me", t, nil)
        if err != nil {
            ui.Red(t[:20] + "...: " + err.Error())
            continue
        }
        resp.Body.Close()
        if resp.StatusCode == 200 {
            ui.Green(t[:20] + "...: VALID")
        } else {
            ui.Red(t[:20] + "...: INVALID")
        }
    }
}

func GuildBackup() {
    ui.Cyan("enter guild ID:")
    var gid string
    fmt.Scanln(&gid)
    ui.Cyan("enter token:")
    var tok string
    fmt.Scanln(&tok)
    resp, err := discordReq("GET", "/guilds/"+gid+"?with_counts=true", tok, nil)
    if err != nil {
        ui.Red(err.Error())
        return
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    jsonStr := string(body)
    // add channels
    chResp, _ := discordReq("GET", "/guilds/"+gid+"/channels", tok, nil)
    chBody, _ := io.ReadAll(chResp.Body)
    chResp.Body.Close()
    full := "{\"guild\":" + jsonStr + ",\"channels\":" + string(chBody) + "}"
    out := "guild_backup_" + gid + ".json"
    os.WriteFile(out, []byte(full), 0644)
    ui.Green("saved to " + out)
}

func MessageLogger() {
    ui.Cyan("enter channel ID:")
    var cid string
    fmt.Scanln(&cid)
    ui.Cyan("enter token:")
    var tok string
    fmt.Scanln(&tok)
    resp, err := discordReq("GET", "/channels/"+cid+"/messages?limit=100", tok, nil)
    if err != nil {
        ui.Red(err.Error())
        return
    }
    defer resp.Body.Close()
    var msgs []map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&msgs)
    for _, m := range msgs {
        author := ""
        if a, ok := m["author"].(map[string]interface{}); ok {
            author = fmt.Sprintf("%v", a["username"])
        }
        fmt.Printf("[%s] %s: %v\n", author, m["timestamp"], m["content"])
    }
}

func VoiceSpam() {
    ui.Cyan("enter guild ID:")
    var gid string
    fmt.Scanln(&gid)
    ui.Cyan("enter channel ID:")
    var cid string
    fmt.Scanln(&cid)
    ui.Cyan("enter token:")
    var tok string
    fmt.Scanln(&tok)
    // join and leave voice repeatedly
    for i := 0; i < 10; i++ {
        discordReq("PATCH", "/guilds/"+gid+"/members/@me", tok, map[string]interface{}{"channel_id": cid})
        time.Sleep(500 * time.Millisecond)
        discordReq("PATCH", "/guilds/"+gid+"/members/@me", tok, map[string]interface{}{"channel_id": nil})
    }
}

func Selfbot() {
    ui.Cyan("launching python selfbot...")
    cmd := exec.Command("python", "internal/discord/selfbot.py")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin
    cmd.Run()
}