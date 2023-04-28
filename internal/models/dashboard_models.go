package models

type Dashboard struct {
	ID          int    `json:"id"`
	UID         string `json:"uid"`
	Version     int    `json:"version"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type DashboardGet struct {
	Meta struct {
		URL string `json:"url"`
	} `json:"meta"`
	Dashboard struct {
		ID          int    `json:"id"`
		UID         string `json:"uid"`
		Version     int    `json:"version"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"dashboard"`
}
