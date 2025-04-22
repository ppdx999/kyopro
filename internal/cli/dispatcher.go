package cli

type Dispatcher struct {
	cmds      map[string]Cmd
	msgSender MsgSender
}

func NewDispatcher(msgSender MsgSender) *Dispatcher {
	return &Dispatcher{
		cmds:      make(map[string]Cmd),
		msgSender: msgSender,
	}
}

func (d *Dispatcher) Register(name string, cmd Cmd) {
	d.cmds[name] = cmd
}

func (d *Dispatcher) Run(args []string) ExitCode {
	if len(args) == 0 {
		d.msgSender.SendMsg("no subcommand provided")
		return ExitErr
	}

	sub, ok := d.cmds[args[0]]
	if !ok {
		d.msgSender.SendMsg("unknown subcommand: " + args[0])
		return ExitErr
	}

	return sub.Run(args[1:])
}
