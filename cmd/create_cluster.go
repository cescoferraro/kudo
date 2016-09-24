package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CreateClusterCmd struct {
	Yes    bool
	Target string
}

var createClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "A brief description of your command",
	Long:  `Creates a full TLS Kubernetes cluster at Digital Ocean.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("create cluster")

	},
}

func init() {
	createCmd.AddCommand(createClusterCmd)

}
