package user

import (
	"os"

	"github.com/ppdx999/kyopro/internal/domain/model"
)

type UserPipeline struct{}

func (u *UserPipeline) Pipeline() *model.Pipeline {
	return &model.Pipeline{
		Inflow:  os.Stdin,
		Outflow: os.Stdout,
		ErrFlow: os.Stderr,
	}
}
