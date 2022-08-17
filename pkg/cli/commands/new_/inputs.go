package new_

import (
	"fmt"
	"log"
	"regexp"

	"github.com/charmbracelet/lipgloss"
	"github.com/erikgeiser/promptkit/textinput"
	projectconstructor "github.com/sbarrios93/pypher/pkg/project_constructor"
)

func promptPackageName(p *projectconstructor.ProjectMeta, name string) {

	titleStyle := lipgloss.NewStyle().
		MarginLeft(1).
		MarginRight(5).
		Padding(0, 1).
		Italic(true).
		Foreground(lipgloss.Color("#FFF7DB")).Align(lipgloss.Bottom)

	fmt.Print("\n")
	fmt.Println(titleStyle.Render("Title"))
	const customTemplate = `
	{{ Bold ( Underline (Color "5" .Prompt)) }}: {{ .Input -}}
	{{- if .ValidationError -}}
		{{- Foreground "1" (Bold " ✘") -}}
	{{- else -}}
		{{- Foreground "2" (Bold " ✔") -}}
	{{- end -}}
	{{- if .ValidationError -}}
		{{- (print " Error: " (Foreground "1" .ValidationError.Error)) -}}
	{{- end -}}
	`
	isKebab := regexp.MustCompile(`^[a-z0-9-]*$`).MatchString

	const customResultTemplate = `
	{{ Bold (print "🖥️  Connecting to " (Foreground "32" .FinalValue) ) -}}`

	input := textinput.New("Package Name")
	input.InitialValue = name

	input.Validate = func(input string) error {
		if len(input) == 0 {
			return fmt.Errorf("package name cannot be empty")
		} else if !isKebab(input) {
			return fmt.Errorf("package name must be kebab-case")
		} else if input[len(input)-1] == '-' {
			return fmt.Errorf("package name cannot end with a hyphen")
		} else if input[0] == '-' {
			return fmt.Errorf("package name cannot start with a hyphen")
		}
		return nil
	}
	input.Template = customTemplate
	// input.InputCursorStyle = lipgloss.NewStyle().Blink(false)
	input.ResultTemplate = customResultTemplate
	input.CharLimit = 1000

	packageName, err := input.RunPrompt()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	p.Name = packageName
}
