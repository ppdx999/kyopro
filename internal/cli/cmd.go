package cli

type Cmd interface {
	Run(args []string) ExitCode
}
