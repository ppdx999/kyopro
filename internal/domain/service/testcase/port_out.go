package testcase

type WdGetter interface {
	GetWd() (string, error)
}

type FileExister interface {
	ExistFile(path string) (bool, error)
}

type PublicDirMaker interface {
	MakePublicDir(path string) error
}

type PublicFileWriter interface {
	WritePublicFile(path string, date []byte) error
}
