package services

import (
	"encoding/json"
	"fmt"
	"golang-grafana-api/internal/models"
	"io"
	"net/http"
)

func GetDashboardInfo(apiKey, apiUrl, dashboardUID string) (*models.Dashboard, error) {
	client := http.Client{}

	url := fmt.Sprintf("%s/api/dashboards/uid/%s", apiUrl, dashboardUID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s", err)
	}

	var dashboard models.Dashboard
	err = json.Unmarshal(body, &dashboard)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %s", err)
	}

	return &dashboard, nil
}
