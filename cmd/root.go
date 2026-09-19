package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"ghostline/internal/ui"
)

var reader = bufio.NewReader(os.Stdin)

func Run() {
	ui.Clear()
	ui.PrintBanner()
	ui.RenderMainMenu()

	for {
		choice := prompt("ghostline")
		switch choice {
		case "a", "A":
			ui.CategoryLeft()
			ui.Clear()
			ui.PrintBanner()
			ui.RenderMainMenu()
		case "d", "D":
			ui.CategoryRight()
			ui.Clear()
			ui.PrintBanner()
			ui.RenderMainMenu()
		case "p", "P":
			ui.PagePrev()
			ui.Clear()
			ui.PrintBanner()
			ui.RenderMainMenu()
		case "n", "N":
			ui.PageNext()
			ui.Clear()
			ui.PrintBanner()
			ui.RenderMainMenu()
		case "99":
			fmt.Println()
			ui.Cyan("closing session. stay ghost.")
			os.Exit(0)
		default:
			handleCategoryChoice(choice)
		}
	}
}

func prompt(prefix string) string {
	fmt.Printf("\n%s@%s:~# ", ui.Green(prefix), ui.Cyan(ui.CurrentCategoryName()))
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func handleCategoryChoice(choice string) {
	cat := ui.CurrentCategory()
	if cat == nil {
		ui.Red("invalid category")
		return
	}
	if idx, err := parseInt(choice); err == nil {
		if idx >= 0 && idx < len(cat.Items) {
			item := cat.Items[idx]
			// brief loading transition
			ui.Transition("loading "+item.Name+"...", 300*time.Millisecond)
			ui.Clear()
			ui.PrintBanner()
			ui.Cyan("─── " + item.Name + " ───")
			ui.White(item.Desc)
			fmt.Println()
			item.Handler()
			pause()
			ui.Clear()
			ui.PrintBanner()
			ui.RenderMainMenu()
			return
		}
	}
	ui.Red("unknown command: " + choice)
	pause()
}

func pause() {
	fmt.Print("\n")
	ui.DimCyan("press enter to continue...")
	reader.ReadString('\n')
}

func parseInt(s string) (int, error) {
	n := 0
	_, err := fmt.Sscanf(s, "%d", &n)
	return n, err
}