package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"tea.kareha.org/cup/termi"
)

func (sh *Shell) Run() {
	args := strings.Split(sh.line.String(), " ")
	if args[0] == "" {
		return
	}

	if args[0] == "q" {
		fmt.Print("quit\r\n")
		sh.alive = false
		return
	} else if args[0] == "ev" {
		if len(args) == 1 {
			for _, s := range os.Environ() {
				fmt.Printf("%s\r\n", s)
			}
		} else if len(args) == 2 {
			fmt.Printf("%s %s\r\n", args[1], os.Getenv(args[1]))
		} else {
			os.Setenv(args[1], strings.Join(args[2:], " "))
		}
		return
	} else if args[0] == "cd" {
		if len(args) == 1 {
			err := os.Chdir(os.Getenv("HOME"))
			if err != nil {
				fmt.Printf("%v\r\n", err)
			}
			return
		} else {
			dir := strings.Join(args[1:], " ")
			err := os.Chdir(dir)
			if err != nil {
				fmt.Printf("%v\r\n", err)
			}
			return
		}
	}

	cmd := exec.Command(args[0], args[1:]...)
	stdio, err := termi.DupStdio()
	if err != nil {
		fmt.Printf("%v\r\n", err)
		return
	}
	stdio.AttachTo(cmd)

	termi.StopKey()
	termi.Cooked()

	err = cmd.Run()
	if err != nil {
		fmt.Printf("%v\r\n", err)
	}
	stdio.Close()

	termi.Raw()
	termi.StartKey()
}
