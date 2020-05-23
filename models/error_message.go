package models

// ErrorMessage represents a json struct with error
type ErrorMessage struct {
	Status int    `json:"status"`
	Err    string `json:"error"`
}
