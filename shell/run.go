package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"tea.kareha.org/cup/termi"
)

func Run(sh *Shell, args []string) error {
	if args[0] == "" {
		return fmt.Errorf("command not found")
	}

	if args[0] == "q" {
		fmt.Print("quit\r\n")
		sh.Alive = false
		return nil
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
		return nil
	} else if args[0] == "cd" {
		if len(args) == 1 {
			return os.Chdir(os.Getenv("HOME"))
		} else {
			dir := strings.Join(args[1:], " ")
			return os.Chdir(dir)
		}
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	termi.FinishKey()
	termi.Cooked()

	err := cmd.Run()

	termi.Raw()
	termi.InitKey()

	return err
}
