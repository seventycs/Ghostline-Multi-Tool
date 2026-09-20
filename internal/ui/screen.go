package ui

import (
	"fmt"
	"strings"
)

// Screen renders a tool header box.
func Screen(title, desc string) {
	w := 70
	inner := w - 4

	title = strings.ToUpper(title)
	tLine := padOrTrim("◈ "+title, inner)
	dLine := padOrTrim(desc, inner)

	fmt.Println()
	fmt.Println(NeonGreen("  ╭" + strings.Repeat("─", w-2) + "╮"))
	fmt.Println(NeonGreen("  │ ") + Green(tLine) + NeonGreen(" │"))
	fmt.Println(NeonGreen("  │ ") + DimCyan(dLine) + NeonGreen(" │"))
	fmt.Println(NeonGreen("  ╰" + strings.Repeat("─", w-2) + "╯"))
	fmt.Println()
}

func padOrTrim(s string, n int) string {
	// rough byte length — fine for ascii + emoji-ish headers
	r := []rune(s)
	if len(r) > n {
		return string(r[:n-1]) + "…"
	}
	return s + strings.Repeat(" ", n-len(r))
}