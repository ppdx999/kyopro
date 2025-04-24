package problem

type GetWd interface {
	GetWd() (string, error)
}

type PublicDirMaker interface {
	MakePublicDir(path string) error
}
