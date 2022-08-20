package new_

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/pelletier/go-toml/v2"
	"github.com/sbarrios93/pypher/pkg/pyproject"
	stringvalidator "github.com/sbarrios93/pypher/pkg/utils/string_validator"
	"github.com/sbarrios93/pypher/pkg/utils/sysinfo"
)

const initialVersion = "0.1.0"

const inputTemplate = `
 {{ Bold .Prompt }}: {{ .Input -}}
`

const inputTemplateWithValidation = `
 {{ Bold .Prompt }}: {{ .Input -}}
{{- if .ValidationError -}}
	{{- Foreground "1" (Bold " ✘") -}}
{{- else -}}
	{{- Foreground "2" (Bold " ✔") -}}
{{- end -}}
{{- if .ValidationError -}}
	{{- (print " Error: " (Foreground "1" .ValidationError.Error)) -}}
{{- end -}}
`

const resultTemplate = `
 {{Bold .Prompt }}: {{ (Foreground "32" .FinalValue) -}}{{- Foreground "2" (Bold " ✔") -}}`

func initPrompt() {
	titleStyle := lipgloss.NewStyle().
		MarginLeft(1).
		MarginRight(5).
		Padding(0, 1).
		Italic(true).
		Foreground(lipgloss.Color("#FFF7DB")).Align(lipgloss.Bottom).Background(lipgloss.AdaptiveColor{Light: "30", Dark: "30"})

	fmt.Print("\n")
	fmt.Printf("%s\n", titleStyle.Render("New Package"))
}

func promptPackageName(p *pyproject.PyProject, name string) {
	// https://packaging.python.org/en/latest/specifications/declaring-project-metadata/#name
	// FIXME: should normnalize according to PEP 503: https://peps.python.org/pep-0503/
	input := textinput.New("Package Name")
	input.InitialValue = name

	input.Validate = func(input string) error {
		if len(input) == 0 {
			return fmt.Errorf("package name cannot be empty")
		} else if !stringvalidator.IsKebab(input) {
			return fmt.Errorf("package name must be kebab-case")
		} else if input[len(input)-1] == '-' {
			return fmt.Errorf("package name cannot end with a hyphen")
		} else if input[0] == '-' {
			return fmt.Errorf("package name cannot start with a hyphen")
		}
		return nil
	}
	input.Template = inputTemplateWithValidation
	// input.InputCursorStyle = lipgloss.NewStyle().Blink(false)
	input.ResultTemplate = resultTemplate
	input.CharLimit = 128

	packageName, err := input.RunPrompt()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	p.ProjectMeta.Name = packageName
}

func promptVersion(p *pyproject.PyProject) {
	// https://packaging.python.org/en/latest/specifications/declaring-project-metadata/#version
	// FIXME: should comply with PEP 440 : https://peps.python.org/pep-0440/
	input := textinput.New("Version")
	input.InitialValue = initialVersion

	input.Validate = func(input string) error {
		if len(input) == 0 {
			return fmt.Errorf("version cannot be empty")
		} else if !stringvalidator.IsSemVer(input) {
			return fmt.Errorf("version must comply with semantic versioning specification")
		}
		return nil
	}

	input.Template = inputTemplateWithValidation
	input.ResultTemplate = resultTemplate
	input.CharLimit = 16

	semVer, err := input.RunPrompt()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	p.ProjectMeta.Version = semVer
}

func promptDescription(p *pyproject.PyProject) {
	input := textinput.New("Description")
	input.Validate = nil
	input.Template = inputTemplate
	input.ResultTemplate = resultTemplate
	input.CharLimit = 4096 // FIXME: https://packaging.python.org/en/latest/specifications/core-metadata/#description establishes that the description length shouldn't be limited

	description, err := input.RunPrompt()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	p.ProjectMeta.Description = description
}

func promptAuthor(p *pyproject.PyProject) {
	// https://packaging.python.org/en/latest/specifications/declaring-project-metadata/#authors-maintainers
	// FIXME: should assign a value according to specifications

	inputAuthor := textinput.New("Author")
	inputAuthor.InitialValue = sysinfo.GetGitConfig().UserName
	inputAuthor.Validate = nil
	inputAuthor.Template = inputTemplate
	inputAuthor.ResultTemplate = resultTemplate
	inputAuthor.CharLimit = 4096

	authorName, err := inputAuthor.RunPrompt()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	authorName = strings.TrimSpace(authorName)

	inputEmail := textinput.New("Email")
	inputEmail.InitialValue = sysinfo.GetGitConfig().UserEmail
	inputEmail.Validate = func(input string) error {
		if len(input) == 0 {
			return nil
		} else if !stringvalidator.IsEmail(input) {
			return fmt.Errorf("not valid email address")
		}
		return nil
	}

	inputEmail.Template = inputTemplateWithValidation
	inputEmail.ResultTemplate = resultTemplate
	inputEmail.CharLimit = 4096
	authorEmail, err := inputEmail.RunPrompt()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	p.ProjectMeta.Authors = []pyproject.AuthorInfo{
		{
			Name:  authorName,
			Email: authorEmail,
		},
	}
}

func promptPythonVersion(p *pyproject.PyProject) {
	input := textinput.New("Python Version")
	input.Validate = nil

	pythonSplitVersion := strings.Split(sysinfo.PythonVersion(), ".")
	input.InitialValue = "^" + pythonSplitVersion[0] + "." + pythonSplitVersion[1]

	input.Template = inputTemplate
	input.ResultTemplate = resultTemplate
	input.CharLimit = 32

	pythonVersion, err := input.RunPrompt()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	p.ProjectMeta.RequiresPython = pythonVersion
}

func promptReadme(p *pyproject.PyProject) {
	// https://packaging.python.org/en/latest/specifications/declaring-project-metadata/#readme
	// FIXME: promptReadme can be either a string or a table.
	// The readme field may also take a table. The file key has a string value representing a path relative to pyproject.toml to a file containing the full description. The text key has a string value which is the full description. These keys are mutually-exclusive, thus tools MUST raise an error if the metadata specifies both keys.

	input := textinput.New("Readme file name")
	input.InitialValue = "README.md"
	input.Validate = nil

	input.Template = inputTemplate
	input.ResultTemplate = resultTemplate
	input.CharLimit = 128

	readmeFilename, err := input.RunPrompt()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	p.ProjectMeta.Readme = readmeFilename
}

func RunPrompt(PyProject *pyproject.PyProject, name string) {
	initPrompt()
	promptPackageName(PyProject, name)
	promptVersion(PyProject)
	promptDescription(PyProject)
	promptAuthor(PyProject)
	promptPythonVersion(PyProject)
	promptReadme(PyProject)

	projectWrite, errMarshal := toml.Marshal(PyProject)

	if errMarshal != nil {
		panic(errMarshal)
	}

	errWrite := os.WriteFile("./py/toml/encoded_test.toml", projectWrite, os.ModePerm)
	if errWrite != nil {
		panic(errWrite)
	}
}
