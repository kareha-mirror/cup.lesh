package shell

import (
	"fmt"
	"strings"

	"tea.kareha.org/cup/termi"
	"tea.kareha.org/cup/termi/rutil"
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

				line := strings.TrimSpace(sh.line.String())
				if line == "" {
					sh.line.Reset()
					fmt.Print(sh.Prompt)
					fmt.Print(sh.line.String())
					fmt.Print(termi.ClearTail)
					continue
				}
				args := strings.Split(line, " ")
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

			if key.Rune == 0x09 { // Tab
				args := strings.Split(sh.line.String(), " ")
				list := sh.hooks.CompList(sh, args)
				if len(list) < 1 {
					continue
				}
				list = rutil.FilterByPrefix(list, args[len(args)-1])
				if len(list) < 1 {
					continue
				}
				if len(list) == 1 {
					args[len(args)-1] = list[0]
					args = append(args, "")
					sh.line.Reset()
					sh.line.WriteString(strings.Join(args, " "))
					fmt.Print("\r")
					fmt.Print(sh.Prompt)
					fmt.Print(sh.line.String())
					fmt.Print(termi.ClearTail)
					continue
				}
				args[len(args)-1] = rutil.CommonPrefix(list)
				sh.line.Reset()
				sh.line.WriteString(strings.Join(args, " "))

				fmt.Print("\r")
				fmt.Print(strings.Join(list, "  "))
				fmt.Print(termi.ClearTail)
				fmt.Print("\r\n")
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
