package response

var (
	ErrInvalidPathParameterResponse = ErrorResponse{Error: "Invalid path parameter"}
	ErrBlankPathParameterResponse   = ErrorResponse{Error: "No path parameter"}
	ErrInvalidContentTypeResponse   = ErrorResponse{Error: "Invalid request content-type"}
	ErrInvalidMediaTypeResponse     = ErrorResponse{Error: "Invalid mediatype"}
	ErrInvalidBodyResponse          = ErrorResponse{Error: "Invalid request body"}
	ErrInvalidBodyNoSerialResponse  = ErrorResponse{Error: "No serial number provided"}
	ErrInternalServerErrorResponse  = ErrorResponse{Error: "Internal server error"}
)

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}
