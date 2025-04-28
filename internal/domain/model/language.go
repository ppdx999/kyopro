package model

type Language interface {
	Name() string
	SourceCode() *SourceCode
	Build(p *Pipeline) error
	Run(p *Pipeline) error
	Clean(p *Pipeline) error
}
