package session

type UserInput interface {
	UserInput() (string, error)
}

type Home interface {
	Home() (string, error)
}

type MakePublicDir interface {
	MakePublicDir(path string) error
}

type ReadSecretFile interface {
	ReadSecretFile(path string) ([]byte, error)
}

type WriteSecretFile interface {
	WriteSecretFile(path string, data []byte) error
}
