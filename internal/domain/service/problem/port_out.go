package problem

type WdGetter interface {
	GetWd() (string, error)
}

type PublicDirMaker interface {
	MakePublicDir(path string) error
}
