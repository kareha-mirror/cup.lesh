package shell

import (
	"fmt"

	"tea.kareha.org/cup/termi"
)

func (sh *Shell) Main() {
	fmt.Print(Prompt)
	for sh.alive {
		seq := termi.ReadSeq()
		switch seq.Kind {
		case termi.SeqRune:
			if sh.line.Len() < 1 && seq.Rune == '\x04' { // Ctrl-D
				fmt.Print("\r\n")
				fmt.Print("exit\r\n")
				sh.alive = false
				continue
			}

			if seq.Rune == termi.RuneEnter {
				fmt.Print("\r\n")

				if sh.line.String() == "exit" {
					sh.alive = false
					continue
				}

				sh.Run()

				sh.line.Reset()
				fmt.Print(Prompt)
				continue
			}

			if seq.Rune == termi.RuneBackspace || seq.Rune == termi.RuneDelete {
				sh.line.RemoveTail()
				fmt.Print("\r")
				fmt.Print(Prompt)
				fmt.Print(sh.line.String())
				fmt.Print(termi.ClearTail)
				continue
			}

			sh.line.WriteRune(seq.Rune)
			fmt.Printf("%c", seq.Rune)
		}
	}
}
