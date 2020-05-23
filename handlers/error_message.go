package handlers

type errorMessage struct {
	Status int    `json:"status"`
	Err    string `json:"error"`
}
