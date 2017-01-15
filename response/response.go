package response

type Response struct {
	OK bool `json:"ok"`
}

type ErrorResponse struct {
	Response
  ErrorMessage string `json:"error,omitempty"`
}

func OK() Response {
  return Response{true}
}

func Error(message string) ErrorResponse {
  return ErrorResponse{Response{false}, message}
}
