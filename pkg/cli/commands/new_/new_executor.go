package new_

import (
	"log"

	"github.com/iancoleman/strcase"
	"github.com/sbarrios93/pypher/pkg/pyproject"
	"github.com/sbarrios93/pypher/pkg/utils/paths"
)

func RunNew(projectPath *paths.ProjectPath, name string, unattended bool, init_ bool) {
	// Initialize new project
	PyProject := pyproject.NewPyProject()

	// Cannot run unattended mode if we dont provide a name
	if unattended && name == "" {
		log.Fatalf("cannot run unattended mode without a name")
	}

	if name == "" {
		name = strcase.ToKebab(projectPath.Name)
	} else {
		name = strcase.ToKebab(name)
	}

	if !unattended {
		RunPrompt(PyProject, name)
	}

}
