package shell

import (
	"tea.kareha.org/cup/termi"
	"tea.kareha.org/cup/termi/rbuf"
)

type Shell struct {
	hooks  Hooks
	Alive  bool
	line   rbuf.RuneBuf
	Prompt string
}

func Init(cfgDir string, paths []string, hooks Hooks) (*Shell, error) {
	sh := &Shell{
		hooks:  hooks,
		Alive:  true,
		line:   rbuf.RuneBuf{},
		Prompt: "> ",
	}

	termi.Raw()
	termi.InitKey()

	return sh, nil
}

func (sh *Shell) Finish() error {
	termi.FinishKey()
	termi.Cooked()
	return nil
}

func OnQuit(sh *Shell) error {
	return nil
}

var (
	compList = []string{
		"q",
		"ev",
		"cd",
	}

	dummyFileList = []string{
		"foobar.txt",
		"hello1.txt",
		"hello2.txt",
		"test.txt",
	}
)

func CompList(sh *Shell, args []string) []string {
	if len(args) > 1 {
		return append([]string(nil), dummyFileList...)
	}
	return append([]string(nil), compList...)
}
