package cli

import (
	"fmt"
	"os"

	"github.com/sbarrios93/pypher/pkg/cli/commands/start"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands

func newCliCommand() *cobra.Command {
	cliCommand := &cobra.Command{
		Use:   "pypher",
		Short: "Pypher is a package manager for Python written in Go",
		Long:  `Pypher manages your packages and dependencies for your Python Projects, based on a pyproject.toml file`,
	}

	cliCommand.AddCommand(start.StartCommand())

	return cliCommand
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Start() {

	if err := newCliCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
