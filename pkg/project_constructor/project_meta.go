package projectconstructor

type ProjectMeta struct {
	Name          string
	Version       string
	Description   string
	Readme        string
	Author        []map[string]string
	PythonVersion string
}

func NewProjectMeta() *ProjectMeta {
	return &ProjectMeta{
		Name:        "",
		Version:     "",
		Description: "",
		Readme:      "",
	}
}
