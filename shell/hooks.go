package shell

type Hooks struct {
	Run      func(*Shell, []string) error
	OnQuit   func(*Shell) error
	CompList func(*Shell, []string) []string
}

func DefaultHooks() Hooks {
	return Hooks{
		Run:      Run,
		OnQuit:   OnQuit,
		CompList: CompList,
	}
}
