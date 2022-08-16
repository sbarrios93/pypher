package commands

import (
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	// New command can start a new project given a directory name (or in the current directory if its empty). If starting a new project, the command will create the required project structure (and the pyproject.toml file).
	// The `new` command can also initialize a project, where the current path is already an ongoing python project. In this case the command will only create a pyproject.toml file.
	// The following options are available
	// `init`
	// The following flags are available
	// `--dir`
	// `--name`
	// `--interactive`
	newCommand := &cobra.Command{
		Use:       "new [init]",
		Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.MaximumNArgs(1)),
		ValidArgs: []string{"init"},
		Short:     "Start a new Python Project on the given path or initialize a python project under the current directory. Pass the `init` argument to initialize a project instead of starting a new one",
		RunE: func(command *cobra.Command, args []string) error {

			// var err error
			dir, _ := command.Flags().GetString("directory")
			name, _ := command.Flags().GetString("name")
			unattended, _ := command.Flags().GetBool("unattended")

			// resolve directory path
			dir = filepath.Base(dir)
			// if err != nil {
			// log.Fatalf("could not resolve path specified. Got %s", dir)
			// }
			log.Printf("%v, %v, %v", dir, name, unattended)

			if len(args) == 1 && args[0] == "init" {
				RunInit(dir)
			} else if len(args) == 0 {
				RunNew(dir)
			} else {
				log.Fatalf("Number of args does not match valid arguments. Got %q", args)
			}
			return nil
		},
	}

	newCommand.PersistentFlags().StringP("directory", "d", ".", "directory where to run the command")
	newCommand.PersistentFlags().StringP("name", "n", "", "Name of project")
	newCommand.PersistentFlags().BoolP("unattended", "u", false, "Run command in unattended mode (non interactive)")

	return newCommand
}
