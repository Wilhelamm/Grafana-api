package cmd

import (
	"fmt"
	"golang-grafana-api/internal/services"
	"github.com/spf13/cobra"
)


var deleteDashboard = &cobra.Command{
	Use:   "delete",
	Short: "Lists deployed dashboards",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey, _ := cmd.Flags().GetString("api-token")
		apiUrl, _ := cmd.Flags().GetString("grafana-url")
		// argValue := args[0]
        flagValue, _ := cmd.Flags().GetString("uid")
		err := services.DeleteDashboardByUID(apiKey, apiUrl, flagValue)
		if err != nil {
			return fmt.Errorf("Error getting dashboards: %s", err)
		}
		fmt.Println("flag", flagValue)
		fmt.Println(err)
		return nil 
	},
}

func init() {
	monitoringCmd.AddCommand(deleteDashboard)
}
