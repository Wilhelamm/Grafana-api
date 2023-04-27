package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var AppVersion string

var rootCmd = &cobra.Command{
	Use: "monitoring-cli",
	Short: "A simple CLI to manage certain aspects of the Grafana API",
	Long: "A CLI to create, list, get and delete Grafana Dashboards via API",
	Version: AppVersion,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("api-token", "eyJrIjoibUpUaEIyM1ZGZzFBZ1dSVGNNVWYySGxGelpIbUI2YnEiLCJuIjoidGVzdC1icnVubyIsImlkIjoxfQ==", "Grafana API token")
	rootCmd.PersistentFlags().String("grafana-url", "https://monitoring.airgap.zpecloud.local", "Grafana URL")
}