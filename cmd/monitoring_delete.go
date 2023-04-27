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

// func init() {
// 	createCmd := &cobra.Command{
// 		Use:   "create",
// 		Short: "Creates dashboard from Json",
// 		Long:  ``,
// 		Run: func(cmd *cobra.Command, args []string) {
// 			apiKey := "eyJrIjoibUpUaEIyM1ZGZzFBZ1dSVGNNVWYySGxGelpIbUI2YnEiLCJuIjoidGVzdC1icnVubyIsImlkIjoxfQ=="
// 			apiUrl := "http://localhost:3000"
// 			resp, err := services.CreateDashboard(apiKey, apiUrl, "/home/terres/go/src/golang-grafana-api/assets/dashboards/astarte-app-engine.json")

// 			if err != nil {
// 				fmt.Println("Error getting dashboards:", err)
// 				return
// 			}
// 			fmt.Println(resp)
// 		},
// 	}

// 	listCmd := &cobra.Command{
// 		Use:   "list",
// 		Short: "Lists deployed dashboards",
// 		Long:  ``,
// 		Run: func(cmd *cobra.Command, args []string) {
// 			apiKey := "eyJrIjoibUpUaEIyM1ZGZzFBZ1dSVGNNVWYySGxGelpIbUI2YnEiLCJuIjoidGVzdC1icnVubyIsImlkIjoxfQ=="
// 			apiUrl := "http://localhost:3000"
// 			resp, err := services.ListDashboards(apiKey, apiUrl)

// 			if err != nil {
// 				fmt.Println("Error getting dashboards:", err)
// 				return
// 			}
// 			fmt.Println(resp)
// 		},
// 	}

// 	getCmd := &cobra.Command{
// 		Use:   "get",
// 		Short: "Lists deployed dashboards",
// 		Long:  ``,
// 		Run: func(cmd *cobra.Command, args []string) {
// 			apiKey := "eyJrIjoibUpUaEIyM1ZGZzFBZ1dSVGNNVWYySGxGelpIbUI2YnEiLCJuIjoidGVzdC1icnVubyIsImlkIjoxfQ=="
// 			apiUrl := "http://localhost:3000"
// 			// argValue := args[0]
//             flagValue, _ := cmd.Flags().GetString("uid")
// 			dashboard, err := services.GetDashboardInfo(apiKey, apiUrl, flagValue)
// 			if err != nil {
// 				fmt.Println("Error getting dashboards:", err)
// 				return
// 			}
// 			fmt.Println("flag", flagValue)
// 			fmt.Println("Dashboard ID:", dashboard)
// 		},
// 	}

// 	deleteCmd := &cobra.Command{
// 		Use:   "delete",
// 		Short: "Lists deployed dashboards",
// 		Long:  ``,
// 		Run: func(cmd *cobra.Command, args []string) {
// 			apiKey := "eyJrIjoibUpUaEIyM1ZGZzFBZ1dSVGNNVWYySGxGelpIbUI2YnEiLCJuIjoidGVzdC1icnVubyIsImlkIjoxfQ=="
// 			apiUrl := "http://localhost:3000"
// 			// argValue := args[0]
//             flagValue, _ := cmd.Flags().GetString("uid")
// 			err := services.DeleteDashboardByUID(apiKey, apiUrl, flagValue)
// 			if err != nil {
// 				fmt.Println("Error getting dashboards:", err)
// 				return
// 			}
// 			fmt.Println("flag", flagValue)
// 		},
// 	}

// 	rootCmd.PersistentFlags().String("uid", "", "A flag for the command")
// 	rootCmd.AddCommand(createCmd)
// 	rootCmd.AddCommand(getCmd)
// 	rootCmd.AddCommand(listCmd)
// 	rootCmd.AddCommand(deleteCmd)

// }

// // func grafana() {
// // 	if err := rootCmd.Execute(); err != nil {
// // 		fmt.Println(err)
// // 		os.Exit(1)
// // 	}
// // }
