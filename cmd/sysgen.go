package cmd

import (
    "ghostline/internal/sysgen"
    "ghostline/internal/ui"
)

func init() {
    ui.RegisterCategory(ui.Category{
        Name: "SYS/GEN",
        Items: []ui.Item{
            {"Base64 Codec", "encode or decode strings", sysgen.Base64Codec},
            {"System Info", "inspect local hardware specs", sysgen.SystemInfo},
            {"IP Pinger", "ping target IPs for latency", sysgen.IPPinger},
            {"Obfuscator", "obfuscate python to prevent reverse", sysgen.Obfuscator},
            {"Metadata Scan", "analyze and remove EXIF data", sysgen.MetadataScan},
            {"App Info", "display version, license, developers", sysgen.AppInfo},
            {"App Config", "manage themes, updates, startup", sysgen.AppConfig},
            {"Disable AV", "add defender exclusions", sysgen.DisableAV},
            {"Windows Debloater", "launch utility to debloat and optimize", sysgen.Debloater},
            {"Proxy Scraper", "scrape HTTP, SOCKS4, SOCKS5", sysgen.ProxyScraper},
            {"Proxy Checker", "test scraped proxies for validity", sysgen.ProxyChecker},
            {"Registry Editor", "safe registry tweak manager", sysgen.RegistryEditor},
            {"Service Manager", "list and control windows services", sysgen.ServiceManager},
            {"Startup Manager", "manage startup programs", sysgen.StartupManager},
            {"File Shredder", "securely delete files", sysgen.FileShredder},
            {"Password Gen", "cryptographically secure passwords", sysgen.PasswordGen},
            {"Hash Cracker", "crack md5/sha1/sha256 hashes", sysgen.HashCracker},
            {"Port Forward", "set up local port forwarding", sysgen.PortForward},
            {"Packet Sniffer", "capture local network packets", sysgen.PacketSniffer},
            {"WiFi Scanner", "scan nearby wifi networks", sysgen.WiFiScanner},
        },
    })
}