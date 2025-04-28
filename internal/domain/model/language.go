package model

type Language interface {
	Name() string
	MainFile() string
	Build(p *Pipeline) error
	Run(p *Pipeline) error
	Clean(p *Pipeline) error
}
