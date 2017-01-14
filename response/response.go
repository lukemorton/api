package response

type Response struct {
	Success bool `json:"success" valid:"required"`
}

func SuccessResponse() Response {
  return Response{true}
}
