package model

type Language struct {
	Name     string
	MainFile string
	Builder  LanguageBuilder
	Runner   LanguageRunner
	Cleaner  LanguageCleaner
}

func NewLanguage() *Language {
	return &Language{}
}

type LanguageBuilder interface {
	Build(sourceFile string, p *Pipeline) error
}

type LanguageRunner interface {
	Run(entryFile string, p *Pipeline) error
}

type LanguageCleaner interface {
	Clean(entryFile string, p *Pipeline) error
}

func (l *Language) Build(p *Pipeline) error {
	return l.Builder.Build(l.MainFile, p)
}

func (l *Language) Run(p *Pipeline) error {
	return l.Runner.Run(l.MainFile, p)
}

func (l *Language) Clean(p *Pipeline) error {
	return l.Cleaner.Clean(l.MainFile, p)
}
