package main

import (
	"fmt"
)

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
