package cli

import (
	"bytes"
	"fmt"

	"github.com/ppdx999/kyopro/internal/domain/service/user"
)

type Dispatcher struct {
	cmds      map[string]Cmd
	msgSender user.MsgSender
}

func NewDispatcher(msgSender user.MsgSender) *Dispatcher {
	return &Dispatcher{
		cmds:      make(map[string]Cmd),
		msgSender: msgSender,
	}
}

func (d *Dispatcher) Name() string {
	return "kyopro"
}

func (d *Dispatcher) Description() string {
	return "競技プログラミングをちょっと便利にするコマンド群"
}

func (d *Dispatcher) Usage() string {
	var usage bytes.Buffer
	usage.WriteString(
		fmt.Sprintf("%s  -   %s\n", d.Name(), d.Description()),
	)
	usage.WriteString("\nUsage:\n")
	for name, cmd := range d.cmds {
		line := fmt.Sprintf("  %-15s %s\n", name, cmd.Description())
		usage.WriteString(line)
	}
	return usage.String()
}

func (d *Dispatcher) Register(cmd Cmd) {
	d.cmds[cmd.Name()] = cmd
}

func (d *Dispatcher) Run(args []string) ExitCode {
	if len(args) == 0 {
		d.msgSender.SendMsg(d.Usage())
		return ExitErr
	}

	if args[0] == "-h" || args[0] == "--help" {
		d.msgSender.SendMsg(d.Usage())
		return ExitErr
	}

	sub, ok := d.cmds[args[0]]
	if !ok {
		d.msgSender.SendMsg("unknown subcommand: " + args[0])
		d.msgSender.SendMsg(d.Usage())
		return ExitErr
	}

	return sub.Run(args[1:])
}
