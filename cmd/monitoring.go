package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var monitoringCmd = &cobra.Command{
	Use: "monitoring",
	Short: "The monitoring command contains subcommands such as create, list, get and delete",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("monitoring-cli")
	},
}

func init() {
	monitoringCmd.PersistentFlags().String("dash-name", "", "A flag to select the dashboard ")
	monitoringCmd.PersistentFlags().String("uid", "", "A flag for the command")
	rootCmd.AddCommand(monitoringCmd)
}