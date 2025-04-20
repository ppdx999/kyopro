package model

type Workspace struct {
	Path string
}

func NewWorkspace(path string) *Workspace {
	if path == "" {
		panic("workspace path is empty")
	}
	return &Workspace{
		Path: path,
	}
}
