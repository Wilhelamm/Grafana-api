package cmd

import (
	"fmt"

	"golang-grafana-api/internal/services"
	"github.com/spf13/cobra"
)

var getDashboard = &cobra.Command{
	Use:   "get",
	Short: "Lists deployed dashboards",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey, _ := cmd.Flags().GetString("api-token")
		apiUrl, _ := cmd.Flags().GetString("grafana-url")
	    flagValue, _ := cmd.Flags().GetString("uid")
		response, err := services.GetDashboardInfo(apiKey, apiUrl, flagValue)
		if err != nil {
			return fmt.Errorf("Error getting dashboards: %s", err)
		}
		fmt.Println("Dashboard ID:", response.Dashboard.ID)
		fmt.Println("Dashboard UID:", response.Dashboard.UID)
		fmt.Println("Dashboard Version:", response.Dashboard.Version)
		fmt.Println("Dashboard Title:", response.Dashboard.Title)
		fmt.Println("Dashboard Description:", response.Dashboard.Description)
		fmt.Println("Dashboard URL:", apiUrl+response.Dashboard.URL)
		return nil 
	},
}

func init() {
	monitoringCmd.AddCommand(getDashboard)
}
