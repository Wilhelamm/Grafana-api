package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-grafana-api/internal/models"
	"net/http"
	"os"
)

func CreateDashboard(apiKey, apiUrl, jsonFile string) (*models.DashboardResponse, error) {
	// Read the JSON file containing the dashboard template
	template, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	// Make a POST request to the Grafana API to create the dashboard
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dashboards/db", apiUrl), bytes.NewBuffer(template))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the JSON response and return the new dashboard ID and URL
	var dashboardResp models.DashboardResponse
	err = json.NewDecoder(resp.Body).Decode(&dashboardResp)
	if err != nil {
		return nil, err
	}

	return &dashboardResp, nil
}
