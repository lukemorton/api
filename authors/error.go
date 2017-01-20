package api

type ErrorBody struct {
	ErrorMessage string `json:"error,omitempty"`
}

func Error(message string) ErrorBody {
	return ErrorBody{message}
}
