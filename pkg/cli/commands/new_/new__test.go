package new__test

import (
	"bytes"
	"testing"

	"github.com/sbarrios93/pypher/pkg/cli/commands/new_"
	"github.com/spf13/cobra"
)

func initialize_command() (*bytes.Buffer, *cobra.Command) {
	cmd := new_.NewCommand()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	return buf, cmd
}

func noUnknownArgs(err error, t *testing.T) {
	if err == nil {
		t.Fatal("Expected an error")
	}
	got := err.Error()
	expected := `invalid argument "anyarg" for "new"`
	if got != expected {
		t.Errorf("Expected: %q, got: %q", expected, got)
	}
}
func TestNewCommand(t *testing.T) {

	t.Run("new takes no unknown args", func(t *testing.T) {
		arg := "anyarg"
		_, cmd := initialize_command()

		cmd.SetArgs([]string{arg})
		_, err := cmd.ExecuteC()

		noUnknownArgs(err, t)

	})
}
