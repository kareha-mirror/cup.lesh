package shell

import (
	"tea.kareha.org/cup/termi"
)

type Shell struct {
	alive  bool
	line   termi.RuneBuf
	prompt string
}

func Init(args []string) *Shell {
	sh := &Shell{
		alive:  true,
		line:   termi.RuneBuf{},
		prompt: "> ",
	}

	termi.Raw()
	termi.StartInput()

	return sh
}

func (sh *Shell) Finish() {
	termi.StopInput()
	termi.Cooked()
}
