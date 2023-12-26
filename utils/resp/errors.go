package resp

import "net/http"

const (
	InvalidAuthToken = "InvalidAuthToken"
	ValidationError  = "ValidationError"
	ServerError      = "ServerError"
)

var ErrorMap = map[string]struct {
	Message    string
	HTTPStatus int
}{
	InvalidAuthToken: {"Invalid or missing authorization token", http.StatusUnauthorized},
	ValidationError:  {"Validation error occurred", http.StatusBadRequest},
	ServerError:      {"Internal server error", http.StatusInternalServerError},
	// Add more mappings as needed
}
