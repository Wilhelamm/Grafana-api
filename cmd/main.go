package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Dashboard struct {
	Meta   interface{} `json:"meta"`
	Model  interface{} `json:"dashboard"`
	Search string      `json:"search"`
	URI    string      `json:"uri"`
}

type DashboardResponse struct {
	ID      int    `json:"id"`
	UID     string `json:"uid"`
	URL     string `json:"url"`
	Status  string `json:"status"`
	Version int    `json:"version"`
}

func CreateDashboard(apiKey, apiUrl, jsonFile string) (*DashboardResponse, error) {
	// Read the JSON file containing the dashboard template
	template, err := ioutil.ReadFile(jsonFile)
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
	var dashboardResp DashboardResponse
	err = json.NewDecoder(resp.Body).Decode(&dashboardResp)
	if err != nil {
		return nil, err
	}

	return &dashboardResp, nil
}

func main() {
	apiKey := "eyJrIjoiMU1EbHFqQVJGRzJNak9OcVNDajBPemEzT0NDWmU2aHQiLCJuIjoiYXBpLXRlc3QiLCJpZCI6MX0="
	apiUrl := "http://localhost:3000"
	jsonFile := "astarte-dash.json"

	dashboardResp, err := CreateDashboard(apiKey, apiUrl, jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("New dashboard created: ID=%d, URL=%s\n", dashboardResp.ID, dashboardResp.URL)
}

