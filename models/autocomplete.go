package models

type RequestGetAutocomplete struct {
	Text  string `json:"text"`
	Limit int64  `json:"limit,omitempty"`
}
