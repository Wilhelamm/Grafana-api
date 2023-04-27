package cmd

import (
	"fmt"

	"golang-grafana-api/internal/services"
	"github.com/spf13/cobra"
)

var	listDashboards = &cobra.Command{
	Use:   "list",
	Short: "Lists deployed dashboards",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, _ := cmd.Flags().GetString("api-token")
		apiUrl, _ := cmd.Flags().GetString("grafana-url")
		resp, err := services.ListDashboards(apiKey, apiUrl)

		if err != nil {
			fmt.Println("Error getting dashboards:", err)
			return
		}

		fmt.Println(resp)
	},
}

func init() {
	monitoringCmd.AddCommand(listDashboards)
}