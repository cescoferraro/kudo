package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(initCmd)
}

// initialize Command
var initCmd = &cobra.Command{
	Use:     "init [name]",
	Aliases: []string{"initialize", "initialise", "create"},
	Short:   "Initialize a Cobra Application",
	Long: `Initialize (cobra init) will create a new application, with a license
and the appropriate structure for a Cobra-based CLI application.
  * If a name is provided, it will be created in the current directory;
  * If no name is provided, the current directory will be assumed;
  * If a relative path is provided, it will be created inside $GOPATH
    (e.g. github.com/spf13/hugo);
  * If an absolute path is provided, it will be created;
  * If the directory already exists but is empty, it will be used.
Init will not use an existing directory with contents.`,

	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			inputPath = ""

		case 1:
			inputPath = args[0]

		default:
			er("init doesn't support more than 1 parameter")
		}
		guessProjectPath()
		initializePath(projectPath)
	},
}

func initializePath(path string) {
	b, err := exists(path)
	if err != nil {
		er(err)
	}

	if !b { // If path doesn't yet exist, create it
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			er(err)
		}
	} else { // If path exists and is not empty don't use it
		empty, err := exists(path)
		if err != nil {
			er(err)
		}
		if !empty {
			er("Cobra will not create a new project in a non empty directory")
		}
	}
	// We have a directory and it's empty.. Time to initialize it.

}
