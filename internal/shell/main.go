package shell

import (
	"fmt"

	"tea.kareha.org/cup/termi"
)

func (sh *Shell) Main() {
	fmt.Print(sh.prompt)

	for sh.alive {
		seq := termi.ReadSeq()
		switch seq.Kind {
		case termi.SeqRune:
			if sh.line.Len() < 1 && seq.Rune == '\x04' { // Ctrl-D
				fmt.Print("\r\n")
				fmt.Print("quit\r\n")
				sh.alive = false
				continue
			}

			if seq.Rune == termi.RuneEnter || seq.Rune == '\n' {
				fmt.Print("\r\n")

				sh.Run()
				sh.line.Reset()

				if sh.alive {
					fmt.Print(sh.prompt)
				}
				continue
			}

			if seq.Rune == termi.RuneBackspace ||
				seq.Rune == termi.RuneDelete {
				sh.line.RemoveTail()
				fmt.Print("\r")
				fmt.Print(sh.line.String())
				fmt.Print(termi.ClearTail)
				continue
			}

			sh.line.WriteRune(seq.Rune)
			fmt.Printf("%c", seq.Rune)
		}
	}
}
