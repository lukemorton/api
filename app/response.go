package app

type Response struct {
	Status int
	Body interface{}
}

type ErrorBody struct {
  ErrorMessage string `json:"error,omitempty"`
}

func OK(body interface{}) Response {
  return Response{200, body}
}

func Error(status int, message string) Response {
  return Response{status, ErrorBody{message}}
}
