package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang-grafana-api/internal/services"
)

var	listDashboards = &cobra.Command{
	Use:   "list",
	Short: "Lists deployed dashboards",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := "eyJrIjoiMTN6YlVFdUJ3UkZEbjdQYkRGUERIYlcyQTBOaXVyMjQiLCJuIjoiVGVzdCIsImlkIjoxfQ=="
		apiUrl := "https://monitoring.airgap.zpecloud.local"
		response, err := services.ListDashboards(apiKey, apiUrl)
		if err != nil {
			fmt.Println("Error getting dashboards:", err)
			return
		}
		for _, dashboard := range response {
			fmt.Println("Dashboard Title:", dashboard.Title)
			fmt.Println("Dashboard ID:", dashboard.ID)
			fmt.Println("Dashboard UID:", dashboard.UID)
			fmt.Println("Dashboard URL:", apiUrl+dashboard.URL+"\n")
		}
	},
}

func init() {
	monitoringCmd.AddCommand(listDashboards)
}