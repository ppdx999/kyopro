package cli

type Cmd interface {
	Name() string
	Description() string
	Run(args []string) ExitCode
}
