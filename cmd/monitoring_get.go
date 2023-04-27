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
		dashboard, err := services.GetDashboardInfo(apiKey, apiUrl, flagValue)
		if err != nil {
			return fmt.Errorf("Error getting dashboards: %s", err)
		}
		fmt.Println(dashboard)
		return nil 
	},
}

func init() {
	monitoringCmd.AddCommand(getDashboard)
}
