package shell

import (
	"tea.kareha.org/cup/termi"
)

var Prompt = "> "

type Shell struct {
	alive bool
	line  termi.RuneBuf
}

func Init(args []string) *Shell {
	sh := &Shell{
		alive: true,
		line:  termi.RuneBuf{},
	}

	termi.Raw()
	termi.StartInput()

	return sh
}

func (sh *Shell) Finish() {
	termi.StopInput()
	termi.Cooked()
}
