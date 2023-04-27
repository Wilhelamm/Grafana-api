package cmd

import (
	"fmt"

	"golang-grafana-api/internal/services"
	"github.com/spf13/cobra"
)

var createDashboard = &cobra.Command{
	Use: "create",
	Short: "Create dashboards from available json",
	Long: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey, _ := cmd.Flags().GetString("api-token")
		apiUrl, _ := cmd.Flags().GetString("grafana-url")

		resp, err := services.CreateDashboard(apiKey, apiUrl, "/home/terres/go/src/golang-grafana-api/assets/dashboards/astarte-app-engine.json")
		if err != nil {
			return fmt.Errorf("Error getting dashboards: %s", err)
		}
		fmt.Println(resp)
		return nil 
	},
}

func init() {
	monitoringCmd.AddCommand(createDashboard)
}