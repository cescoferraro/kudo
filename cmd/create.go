package cmd

import "github.com/spf13/cobra"

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long:  `You can create all the resources`,
}

func init() {
	RootCmd.AddCommand(createCmd)

}
