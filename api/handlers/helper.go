package handlers

type ValidationResponse struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}
