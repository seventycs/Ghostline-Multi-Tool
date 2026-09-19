package ui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unicode/utf8"

	"golang.org/x/term"
)

type Item struct {
	Name    string
	Desc    string
	Handler func()
}

type Category struct {
	Name  string
	Items []Item
}

var (
	categories   []Category
	currentCat   int
	currentPage  int
	itemsPerPage = 14
)

func RegisterCategory(c Category) {
	categories = append(categories, c)
}

func CurrentCategory() *Category {
	if currentCat < 0 || currentCat >= len(categories) {
		return nil
	}
	return &categories[currentCat]
}

func CurrentCategoryName() string {
	c := CurrentCategory()
	if c == nil {
		return "unknown"
	}
	return c.Name
}

func CategoryLeft() {
	currentCat--
	if currentCat < 0 {
		currentCat = len(categories) - 1
	}
	currentPage = 0
	Clear()
	PrintBanner()
	RenderMainMenu()
}

func CategoryRight() {
	currentCat++
	if currentCat >= len(categories) {
		currentCat = 0
	}
	currentPage = 0
	Clear()
	PrintBanner()
	RenderMainMenu()
}

func PagePrev() {
	if currentPage > 0 {
		currentPage--
	}
}

func PageNext() {
	c := CurrentCategory()
	if c == nil {
		return
	}
	maxPage := (len(c.Items) - 1) / itemsPerPage
	if currentPage < maxPage {
		currentPage++
	}
}

const bannerArt = ` ██████╗ ██╗  ██╗ ██████╗ ███████╗████████╗██╗     ██╗███╗   ██╗███████╗
██╔════╝ ██║  ██║██╔═══██╗██╔════╝╚══██╔══╝██║     ██║████╗  ██║██╔════╝
██║  ███╗███████║██║   ██║███████╗   ██║   ██║     ██║██╔██╗ ██║█████╗
██║   ██║██╔══██║██║   ██║╚════██║   ██║   ██║     ██║██║╚██╗██║██╔══╝
╚██████╔╝██║  ██║╚██████╔╝███████║   ██║   ███████╗██║██║ ╚████║███████╗
 ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚══════╝   ╚═╝   ╚══════╝╚═╝╚═╝  ╚═══╝╚══════╝`

func termWidth() int {
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || w <= 0 {
		return 120
	}
	return w
}

func centerBlock(s string, w int) string {
	var out strings.Builder
	for _, line := range strings.Split(s, "\n") {
		trimmed := strings.TrimRight(line, " \t")
		vis := utf8.RuneCountInString(trimmed)
		if vis >= w {
			out.WriteString(line)
			out.WriteString("\n")
			continue
		}
		pad := (w - vis) / 2
		out.WriteString(strings.Repeat(" ", pad))
		out.WriteString(line)
		out.WriteString("\n")
	}
	return out.String()
}

func centerLine(s string, w int) string {
	vis := utf8.RuneCountInString(s)
	if vis >= w {
		return s
	}
	return strings.Repeat(" ", (w-vis)/2) + s
}

func PrintBanner() {
	w := termWidth()
	fmt.Print(NeonGreen(centerBlock(bannerArt, w)))
	fmt.Println()
	bar := "──────────────────────────────────────────────────────────────────────────────"
	fmt.Println(DimCyan(centerLine(bar, w)))
	fmt.Println(Cyan(centerLine("advanced multi-tool  ·  v1.0  ·  made by seventycs", w)))
	fmt.Println(DimCyan(centerLine(bar, w)))
	fmt.Println()
}

func paginationDots(current, total int) string {
	if total <= 1 {
		return ""
	}
	var b strings.Builder
	for i := 0; i < total; i++ {
		if i == current {
			b.WriteString(NeonGreen("●"))
		} else {
			b.WriteString(DimCyan("○"))
		}
		if i < total-1 {
			b.WriteString(" ")
		}
	}
	return b.String()
}

func RenderMainMenu() {
	c := CurrentCategory()
	if c == nil {
		return
	}
	w := termWidth()

	// category tabs
	var tabs []string
	for i, cat := range categories {
		if i == currentCat {
			tabs = append(tabs, "\033[1;92m▶ "+strings.ToUpper(cat.Name)+" ◀\033[0m")
		} else {
			tabs = append(tabs, DimCyan(strings.ToUpper(cat.Name)))
		}
	}
	fmt.Println(centerLine(strings.Join(tabs, DimCyan("  ·  ")), w))
	fmt.Println()

	// section title
	fmt.Println(centerLine(NeonGreen("─── ")+Green(strings.ToUpper(c.Name))+NeonGreen(" ───"), w))
	fmt.Println()

	start := currentPage * itemsPerPage
	end := start + itemsPerPage
	if end > len(c.Items) {
		end = len(c.Items)
	}

	for i := start; i < end; i++ {
		item := c.Items[i]
		num := fmt.Sprintf("[%02d]", i)
		line := fmt.Sprintf("%s  %-32s %s %s",
			Green(num),
			item.Name,
			DimCyan("─"),
			White(item.Desc),
		)
		fmt.Println(centerLine(line, w))
	}
	fmt.Println()

	maxPage := (len(c.Items) - 1) / itemsPerPage
	bar := "──────────────────────────────────────────────────────────────────────────────"
	fmt.Println(DimCyan(centerLine(bar, w)))
	dots := paginationDots(currentPage, maxPage+1)
	if dots != "" {
		fmt.Println(centerLine(dots, w))
	}
	footer := fmt.Sprintf("%s %s   %s %s   %s %s   %s %s",
		Green("[P/N]"), DimCyan("page"),
		Green("[A/D]"), DimCyan("category"),
		Green("[#]"),   DimCyan("run"),
		Green("[99]"),  DimCyan("exit"),
	)
	fmt.Println(centerLine(footer, w))
	if dots != "" {
		fmt.Println(centerLine(fmt.Sprintf("%s page %d/%d", DimCyan("·"), currentPage+1, maxPage+1), w))
	}
	fmt.Println(DimCyan(centerLine(bar, w)))
}

func Clear() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[2J\033[H")
	}
}

func Green(s string) string     { return "\033[92m" + s + "\033[0m" }
func NeonGreen(s string) string { return "\033[38;5;46m" + s + "\033[0m" }
func Cyan(s string) string      { return "\033[96m" + s + "\033[0m" }
func DimCyan(s string) string   { return "\033[38;5;30m" + s + "\033[0m" }
func White(s string) string     { return "\033[97m" + s + "\033[0m" }
func Red(s string) string       { return "\033[91m" + s + "\033[0m" }
func Yellow(s string) string    { return "\033[93m" + s + "\033[0m" }