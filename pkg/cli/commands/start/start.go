package start

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type flags struct {
	dir string
}

func StartCommand() *cobra.Command {

	var startFlags flags

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

			var path string
			var err error

			if len(args) == 0 {
				path, err = os.Getwd()
				if err != nil {
					return errors.New("could not get current working directory")
				}
			} else {
				path = args[0]
			}

			fmt.Println(path)
			return nil
		},
	}

	startCommand.Flags().StringVarP(&startFlags.dir, "dir", "d", ".", "The directory where to start the project")

	return startCommand
}
