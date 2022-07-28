package status

type ErrorResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

// NewProxy is a constructor for ErrorResponse.
func NewErrorResponse(code int, message string, details ...string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
		Details: details,
	}
}
