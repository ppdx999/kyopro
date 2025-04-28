package language

import (
	"bytes"

	"github.com/ppdx999/kyopro/internal/domain/model"
)

type testcaseRunner struct{}

func (r *testcaseRunner) Run(l model.Language, t *model.TestCase) ([]byte, []byte, error) {
	inflow := bytes.NewReader(t.Input)
	outflow := bytes.NewBuffer(nil)
	errflow := bytes.NewBuffer(nil)
	p := model.NewPipeline(inflow, outflow, errflow)

	if err := l.Run(p); err != nil {
		return nil, nil, err
	}
	return outflow.Bytes(), errflow.Bytes(), nil
}
