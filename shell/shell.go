package shell

import (
	"tea.kareha.org/cup/termi"
)

type Shell struct {
	hooks  Hooks
	Alive  bool
	line   termi.RuneBuf
	Prompt string
}

func Init(cfgDir string, paths []string, hooks Hooks) (*Shell, error) {
	sh := &Shell{
		hooks:  hooks,
		Alive:  true,
		line:   termi.RuneBuf{},
		Prompt: "> ",
	}

	termi.Raw()
	termi.StartKey()

	return sh, nil
}

func (sh *Shell) Finish() error {
	termi.StopKey()
	termi.Cooked()
	return nil
}

func OnQuit(sh *Shell) error {
	return nil
}
