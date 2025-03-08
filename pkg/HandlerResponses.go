package pkg

type dataResponse struct {
	Data interface{} `json:"data"`
}

type meessageResponse struct {
	Message string `json:"message"`
}

func NewMessageResponse(msg string) *meessageResponse {
	return &meessageResponse{
		Message: msg,
	}
}

func NewDataResponse[T interface{}](data T) *dataResponse {
	return &dataResponse{
		Data: data,
	}
}
