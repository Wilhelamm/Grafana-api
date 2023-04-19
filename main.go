package main

import (
	"fmt"
	"golang-grafana-api/internal/services"
)

func main() {
	apiKey := "eyJrIjoiZlVRNVRlVUFmbFFrNlRmak12UTQ5QjVBWWFJS1NmQkkiLCJuIjoiYXBpdC1lc3QiLCJpZCI6MX0="
	apiUrl := "http://localhost:3000"
	//jsonFile := "rabbitmq.json"

	dashboards, err := services.ListDashboards(apiKey, apiUrl)
	if err != nil {
		fmt.Println("Error getting dashboards:", err)
		return
	}

	// Use the dashboards slice as needed
	fmt.Println("Dashboards:")
	for _, dashboard := range dashboards {
		fmt.Printf("ID: %d,ID: %s, Title: %s\n", dashboard.ID, dashboard.UID, dashboard.Title)
	}
}
