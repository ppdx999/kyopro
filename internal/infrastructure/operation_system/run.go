package operation_system

import (
	"os/exec"

	"github.com/ppdx999/kyopro/internal/domain/model"
)

func (o *OperationSystem) Run(cmd []string, p *model.Pipeline) error {
	c := exec.Command(cmd[0], cmd[1:]...)
	c.Stdin = p.Inflow
	c.Stdout = p.Outflow
	c.Stderr = p.ErrFlow
	return c.Run()
}
