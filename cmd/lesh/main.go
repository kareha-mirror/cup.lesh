package main

import (
	"os"

	"tea.kareha.org/cup/lesh/internal/shell"
)

func main() {
	sh := shell.Init(os.Args)
	defer sh.Finish()
	sh.Main()
}
