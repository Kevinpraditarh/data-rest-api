package models

type Item struct {
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Model       string   `json:"model"`
	Tech        []string `json:"tech"`
	Status      string   `json:"status"`
}
