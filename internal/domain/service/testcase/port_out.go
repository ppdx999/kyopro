package testcase

type WdGetter interface {
	GetWd() (string, error)
}

type FileExister interface {
	ExistFile(path string) (bool, error)
}

type PublicFileReader interface {
	ReadPublicFile(path string) ([]byte, error)
}

type ChildFileNamesGetter interface {
	ChildFileNames(path string) ([]string, error)
}

type PublicDirMaker interface {
	MakePublicDir(path string) error
}

type PublicFileWriter interface {
	WritePublicFile(path string, date []byte) error
}
