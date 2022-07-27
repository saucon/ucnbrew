/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/saucon/ucnbrew/cmd/boilerplate"
	"github.com/spf13/cobra"
)

// brewCmd represents the brew command
var brewCmd = &cobra.Command{
	Use:   "brew",
	Short: "Create a boilerplate using gin, gorm, postgresql, logrus etc",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("brewing a boilerplate for your lazy ass")

		if len(args) == 2 {
			if err := boilerplate.Brewer(args[0], args[1]); err != nil {
				fmt.Println("Error: ", err)
			}
		} else {
			fmt.Println("Fuck off!!!")
		}
	},
}

func init() {
	rootCmd.AddCommand(brewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// brewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// brewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
