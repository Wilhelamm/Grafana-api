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
		var jsonFile string

		apiKey, _ := cmd.Flags().GetString("api-token")
		apiUrl, _ := cmd.Flags().GetString("grafana-url")
		jsonName, _ := cmd.Flags().GetString("dash-name")

		switch jsonName {
		case "astarte-app-engine":
			jsonFile = "assets/dashboards/astarte-app-engine.json"
		case "astarte-data-updater-plant":
			jsonFile = "assets/dashboards/astarte-data-updater-plant.json"
		case "astarte-pods-memory-usage":
			jsonFile = "assets/dashboards/astarte-pods-memory-usage.json"
		case "astarte-rabbitmq":
			jsonFile = "assets/dashboards/astarte-rabbitmq.json"
		case "astarte-dashboard":
			jsonFile = "assets/dashboards/astarte-dashboard.json"
		case "astarte-pods-cpu-usage":
			jsonFile = "assets/dashboards/astarte-pods-cpu-usage.json"
		case "astarte-pods-network-io":
			jsonFile = "assets/dashboards/astarte-pods-network-io.json"
		case "astarte-vernemq":
			jsonFile = "assets/dashboards/astarte-vernemq.json"
		default:
			// handle invalid file name
		}

		response, err := services.CreateDashboard(apiKey, apiUrl, jsonFile)
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