package testcase

type GetWd interface {
	GetWd() (string, error)
}

type PublicDirMaker interface {
	MakePublicDir(path string) error
}

type PublicFileWriter interface {
	WritePublicFile(path string, date []byte) error
}
