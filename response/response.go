package response

type Response struct {
	Status int
	Body interface{}
}


type ResponseBody struct {
	OK bool `json:"ok"`
}

type ErrorBody struct {
	ResponseBody
  ErrorMessage string `json:"error,omitempty"`
}

func DefaultOK() Response {
  return Response{200, ResponseBody{true}}
}

func OK(body interface{}) Response {
  return Response{200, body}
}

func Error(status int, message string) Response {
  return Response{status, ErrorBody{ResponseBody{false}, message}}
}
