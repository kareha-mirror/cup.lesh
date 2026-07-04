package shell

type Hooks struct {
	Run    func(*Shell, []string) error
	OnQuit func(*Shell) error
}

func DefaultHooks() Hooks {
	return Hooks{
		Run:    Run,
		OnQuit: OnQuit,
	}
}
