package cmd

import (
    "ghostline/internal/simulators"
    "ghostline/internal/ui"
)

func init() {
    ui.RegisterCategory(ui.Category{
        Name: "Simulators",
        Items: []ui.Item{
            {"Fake Token Gen", "generate fake discord tokens", simulators.FakeTokenGen},
            {"Fake Mail Gen", "generate fake email addresses", simulators.FakeMailGen},
            {"Fake Identity", "generate a full fake identity", simulators.FakeIdentity},
            {"Fake Credit Card", "generate fake CC numbers", simulators.FakeCreditCard},
            {"Social Botter", "simulate social media growth", simulators.SocialBotter},
            {"Fake PayPal OTP", "simulate paypal OTP page", simulators.FakePaypalOTP},
            {"Fake Fortnite", "simulate fortnite account check", simulators.FakeFortnite},
            {"Fake Exodu", "simulate exodus wallet page", simulators.FakeExodus},
            {"Hacker Terminal", "fake hacker terminal for pranks", simulators.HackerTerminal},
            {"Fake Bruteforcer", "simulate a bruteforce attack", simulators.FakeBruteforcer},
            {"QR Code Gen", "generate standard and fake QR", simulators.QRCodeGen},
            {"Fake Discord", "simulate discord login page", simulators.FakeDiscord},
            {"Fake Steam", "simulate steam login page", simulators.FakeSteam},
            {"Fake Instagram", "simulate instagram login page", simulators.FakeInstagram},
        },
    })
}