package shell

import (
	"fmt"
	"strings"

	"tea.kareha.org/cup/termi"
)

func (sh *Shell) Main() error {
	fmt.Print(sh.Prompt)

	for sh.Alive {
		key := <-termi.Keys()
		switch key.Kind {
		case termi.KeyRune:
			if sh.line.RuneCount() < 1 && key.Rune == '\x04' { // Ctrl-D
				fmt.Print("\r\n")
				fmt.Print("quit\r\n")
				err := sh.hooks.OnQuit(sh)
				if err != nil {
					fmt.Printf("%v\r\n", err)
					fmt.Print(sh.Prompt)
					continue
				}
				sh.Alive = false
				continue
			}

			if key.Rune == termi.RuneEnter || key.Rune == termi.RuneNewline {
				fmt.Print("\r\n")

				args := strings.Split(sh.line.String(), " ")
				err := sh.hooks.Run(sh, args)
				if err != nil {
					fmt.Printf("%v\r\n", err)
				}
				sh.line.Reset()

				if sh.Alive {
					fmt.Print(sh.Prompt)
				}
				continue
			}

			if key.Rune == termi.RuneBackspace ||
				key.Rune == termi.RuneDelete {
				sh.line.RemoveTail()
				fmt.Print("\r")
				fmt.Print(sh.Prompt)
				fmt.Print(sh.line.String())
				fmt.Print(termi.ClearTail)
				continue
			}

			sh.line.WriteRune(key.Rune)
			fmt.Printf("%c", key.Rune)
		}
	}
	return nil
}
