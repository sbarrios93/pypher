// PEP 621 specifications: Declaring Project Metadata can be found here
// https://packaging.python.org/en/latest/specifications/declaring-project-metadata/

package pyproject

type PyProject struct {
	BuildSystem BuildSystemInfo `toml:"build-system"`
	ProjectMeta ProjectMetaInfo `toml:"project"`
	Tool        Tools           `toml:"tool"`
}

type BuildSystemInfo struct {
	Backend  string   `toml:"build-backend"`
	Requires []string `toml:"requires"`
}

type ProjectMetaInfo struct {
	// TODO include Entry Points
	Name                 string                   `toml:"name"`
	Version              string                   `toml:"version"`
	Description          string                   `toml:"description"`
	Readme               string                   `toml:"readme"`
	RequiresPython       string                   `toml:"requires-python"`
	License              LicenseInfo              `toml:"license, inline"`
	Keywords             []string                 `toml:"keywords"`
	Authors              []AuthorInfo             `toml:"authors,inline"`
	Classifiers          []string                 `toml:"classifiers"`
	Dependencies         []string                 `toml:"dependencies"`
	OptionalDependencies OptionalDependenciesInfo `toml:"optional-dependencies"`
	Dynamic              []string                 `toml:"dynamic"`
	Urls                 UrlsInfo                 `toml:"urls"`
}

type LicenseInfo map[string]string
type UrlsInfo map[string]string
type OptionalDependenciesInfo map[string]string

type AuthorInfo struct {
	Name  string `toml:"name, inline"`
	Email string `toml:"email, inline"`
}

type Tools struct {
	Pypher map[string]interface{} `toml:"pypher"`
}

func NewPyProject() *PyProject {
	return &PyProject{}
}
