package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang-grafana-api/internal/services"
	"os"
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Creates dashboard from Json",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			apiKey := "eyJrIjoiMTN6YlVFdUJ3UkZEbjdQYkRGUERIYlcyQTBOaXVyMjQiLCJuIjoiVGVzdCIsImlkIjoxfQ=="
			apiUrl := "https://monitoring.airgap.zpecloud.local"
			response, err := services.CreateDashboard(apiKey, apiUrl, "assets/dashboards/astarte-app-engine.json")
			if err != nil {
				fmt.Println("Error getting dashboards:", err)
				return
			}
			fmt.Println("Dashboard created with success")
			fmt.Println("Dashboard ID:", response.ID)
			fmt.Println("Dashboard UID:", response.UID)
			fmt.Println("Dashboard URL:", apiUrl+response.URL)
		},
	}

	listCmd := &cobra.Command{
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

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Lists deployed dashboards",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			apiKey := "eyJrIjoiMTN6YlVFdUJ3UkZEbjdQYkRGUERIYlcyQTBOaXVyMjQiLCJuIjoiVGVzdCIsImlkIjoxfQ=="
			apiUrl := "https://monitoring.airgap.zpecloud.local"
			flagValue, _ := cmd.Flags().GetString("uid")
			response, err := services.GetDashboardInfo(apiKey, apiUrl, flagValue)
			if err != nil {
				fmt.Println("Error getting dashboards:", err)
				return
			}
			fmt.Println("Dashboard ID:", response.Dashboard.ID)
			fmt.Println("Dashboard UID:", response.Dashboard.UID)
			fmt.Println("Dashboard Version:", response.Dashboard.Version)
			fmt.Println("Dashboard Title:", response.Dashboard.Title)
			fmt.Println("Dashboard Description:", response.Dashboard.Description)
			fmt.Println("Dashboard URL:", apiUrl+response.Dashboard.URL)
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Lists deployed dashboards",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			apiKey := "eyJrIjoiMTN6YlVFdUJ3UkZEbjdQYkRGUERIYlcyQTBOaXVyMjQiLCJuIjoiVGVzdCIsImlkIjoxfQ=="
			apiUrl := "https://monitoring.airgap.zpecloud.local"
			flagValue, _ := cmd.Flags().GetString("uid")
			err := services.DeleteDashboardByUID(apiKey, apiUrl, flagValue)
			if err != nil {
				fmt.Println("Error getting dashboards:", err)
				return
			}
		},
	}

	rootCmd.PersistentFlags().String("uid", "", "A flag for the command")
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
