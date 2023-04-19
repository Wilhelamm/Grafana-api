package services

import (
	"encoding/json"
	"fmt"
	"golang-grafana-api/internal/models"
	"io/ioutil"
	"net/http"
)

func ListDashboards(apiKey, apiUrl string) ([]models.Dashboard, error) {
	client := http.Client{}

	url := fmt.Sprintf("%s/api/search?type=dash-db", apiUrl)

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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s", err)
	}

	var dashboards []models.Dashboard
	err = json.Unmarshal(body, &dashboards)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %s", err)
	}

	return dashboards, nil
}
