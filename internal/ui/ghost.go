package ui

import (
	"fmt"
	"math/rand"
	"time"
)

const ghostArt = `
                          ▄▄▄▄▄▄▄▄▄▄▄
                      ▄███████████████████▄
                    ▄███████████████████████▄
                   ███████████████████████████
                  █████████████████████████████
                 ███████████████████████████████
                 ████████  ████████  ██████████
                 ████████  ████████  ██████████
                 ████████  ████████  ██████████
                 ███████████████████████████████
                 ███████████████████████████████
                 ███████████████████████████████
                 ███████████████████████████████
                 ███████████████████████████████
                 ███████████████████████████████
                  ████████                ████████
                   ████████              ████████
                    ████████            ████████
                     ████████          ████████
                     ████  ████      ████  ████
                     ███    ████    ████    ███
                     ██      ██      ██      ██
`

// Startup shows the glitched ghost, title, and line, then waits 3s.
func Startup() {
	Clear()
	// flicker phase — 6 frames of increasing corruption
	for i := 0; i < 6; i++ {
		Clear()
		fmt.Println(glitch(ghostArt, i))
		time.Sleep(time.Duration(50+i*15) * time.Millisecond)
	}
	Clear()
	fmt.Println(NeonGreen(ghostArt))
	fmt.Println()
	fmt.Println(Cyan("                                 g  h  o  s  t  l  i  n  e"))
	time.Sleep(400 * time.Millisecond)
	fmt.Println()
	fmt.Println(NeonGreen("        ────────────────────────────────────────────────────────────────"))
	time.Sleep(3 * time.Second)
}

// glitch corrupts the art with random glyphs and rotates colors.
func glitch(art string, intensity int) string {
	out := []rune(art)
	glyphs := []rune{'░', '▒', '▓', '█', '/', '\\', '|', '_', '·', '+', '*', '▄', '▀'}
	n := 15 + intensity*25
	for i := 0; i < n; i++ {
		idx := rand.Intn(len(out))
		if out[idx] != '\n' && out[idx] != ' ' {
			out[idx] = glyphs[rand.Intn(len(glyphs))]
		}
	}
	switch intensity % 3 {
	case 0:
		return Green(string(out))
	case 1:
		return Cyan(string(out))
	default:
		return NeonGreen(string(out))
	}
}