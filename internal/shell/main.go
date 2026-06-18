package shell

import (
	"fmt"

	"tea.kareha.org/cup/termi"
)

func (sh *Shell) Main() {
	fmt.Print(sh.prompt)

	for sh.alive {
		key := <-termi.Keys()
		switch key.Kind {
		case termi.KeyRune:
			if sh.line.Len() < 1 && key.Rune == '\x04' { // Ctrl-D
				fmt.Print("\r\n")
				fmt.Print("quit\r\n")
				sh.alive = false
				continue
			}

			if key.Rune == termi.RuneEnter || key.Rune == termi.RuneNewline {
				fmt.Print("\r\n")

				sh.Run()
				sh.line.Reset()

				if sh.alive {
					fmt.Print(sh.prompt)
				}
				continue
			}

			if key.Rune == termi.RuneBackspace ||
				key.Rune == termi.RuneDelete {
				sh.line.RemoveTail()
				fmt.Print("\r")
				fmt.Print(sh.line.String())
				fmt.Print(termi.ClearTail)
				continue
			}

			sh.line.WriteRune(key.Rune)
			fmt.Printf("%c", key.Rune)
		}
	}
}
