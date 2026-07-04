package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"tea.kareha.org/cup/lesh/shell"
	"tea.kareha.org/cup/termi"
)

const appName = "lesh"

const failure = 1

func realMain() (totalErr error) {
	// parse options
	cfgDir := flag.String("d", "", "config directory")
	unlock := flag.Bool("unlock", false, "unlock")
	flag.Parse()

	if *cfgDir == "" {
		dir, err := os.UserConfigDir()
		if err != nil {
			return err
		}
		// default config directory
		*cfgDir = filepath.Join(dir, appName)
	}

	if *unlock {
		err := termi.Unlock(*cfgDir)
		if err != nil {
			return err
		}
		return nil
	}

	paths := flag.Args()

	// init shell
	sh, err := shell.Init(*cfgDir, paths, shell.DefaultHooks())
	if err != nil {
		return err
	}
	defer func() {
		if err := sh.Finish(); err != nil {
			totalErr = errors.Join(totalErr, err)
		}
	}()

	// enter main loop
	if err := sh.Main(); err != nil {
		return err
	}

	return totalErr
}

func main() {
	if err := realMain(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(failure)
	}
}
