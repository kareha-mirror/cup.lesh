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
	termi.StartKey()

	return sh
}

func (sh *Shell) Finish() {
	termi.StopKey()
	termi.Cooked()
}
