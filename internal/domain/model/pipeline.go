package model

type Pipeline struct {
	Inflow  Inflow
	Outflow Outflow
	ErrFlow ErrFlow
}

func NewPipeline(inflow Inflow, outflow Outflow, errflow ErrFlow) *Pipeline {
	return &Pipeline{
		Inflow:  inflow,
		Outflow: outflow,
		ErrFlow: errflow,
	}
}

type Inflow interface {
	Read(p []byte) (n int, err error)
}

type Outflow interface {
	Write(p []byte) (n int, err error)
}

type ErrFlow interface {
	Write(p []byte) (n int, err error)
}
