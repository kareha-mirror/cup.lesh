package shell

import (
	"fmt"
	"os/exec"
	"strings"

	"tea.kareha.org/cup/termi"
)

func (sh *Shell) Run() {
	args := strings.Split(sh.line.String(), " ")
	if len(args) < 1 {
		return
	}

	cmd := exec.Command(args[0], args[1:]...)
	setup(cmd)

	termi.StopInput()
	termi.Cooked()
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\r\n", err)
	}

	terminate()

	termi.Raw()
	termi.StartInput()
}
