package models

type Dashboard struct {
	ID     int         `json:"id"`
	UID    string      `json:"uid"`
	Title  string      `json:"title"`
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
