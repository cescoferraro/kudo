package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func AskForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

func AskQuestion(question string) string {
	fmt.Printf(question + " - ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()
}

func CreateShortStringFlag(cmd *cobra.Command, envname, description, name, short, thisdefault string) {
	viper.BindEnv(envname)
	cmd.Flags().StringP(name, short, thisdefault, description)
	viper.BindPFlag(envname, cmd.Flags().Lookup(name))
}

func CreateShortBoolFlag(cmd *cobra.Command, envname, description, name, short string, thisdefault bool) {
	viper.BindEnv(envname)
	cmd.Flags().BoolP(name, short, thisdefault, description)
	viper.BindPFlag(envname, cmd.Flags().Lookup(name))
}
