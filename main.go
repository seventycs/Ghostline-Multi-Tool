package main

import (
	"ghostline/cmd"
	"ghostline/internal/ui"
	"ghostline/internal/updater"
)

func main() {
	go updater.Check()
	ui.Startup()
	ui.Clear()
	ui.PrintBanner()
	cmd.Run()
}