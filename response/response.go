package response

type Response struct {
	OK bool `json:"ok"`
  ErrorMessage string `json:"error,omitempty"`
}

func OK() Response {
  return Response{OK: true}
}

func Error(message string) Response {
  return Response{OK: false, ErrorMessage: message}
}
