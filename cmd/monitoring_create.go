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

		response, err := services.CreateDashboard(apiKey, apiUrl, "assets/dashboards/astarte-dashboard.json")
		if err != nil {
			return fmt.Errorf("Error getting dashboards: %s", err)
		}
		fmt.Println("Dashboard created with success", response)
		fmt.Println("Dashboard ID:", response.ID)
		fmt.Println("Dashboard UID:", response.UID)
		fmt.Println("Dashboard URL:", apiUrl+response.URL)
		return nil 
	},
}

func init() {
	monitoringCmd.AddCommand(createDashboard)
}