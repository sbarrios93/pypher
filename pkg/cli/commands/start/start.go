package start

import (
	"fmt"

	"github.com/sbarrios93/pypher/pkg/common/sysinfo"
	"github.com/spf13/cobra"
)

func StartCommand() *cobra.Command {

	startCommand := &cobra.Command{
		Use: "start",
		Args: func(command *cobra.Command, args []string) error {
			if len(args) > 1 {
				return fmt.Errorf("%v command only accepts at most 1 argument, got %v", command.Name(), len(args))
			}
			return nil
		},
		Short: "Initialize a Python Project under the current directory or directory path specified",
		RunE: func(command *cobra.Command, args []string) error {
			// TODO: Refactor command to new <-> start. New should create new directory, start should only create a pyproject.toml file on the current  directory
			var path string
			if len(args) == 0 {
				path = sysinfo.Getwd()
			} else {
				path = args[0]
			}

			Run(path)

			return nil
		},
	}

	return startCommand
}
