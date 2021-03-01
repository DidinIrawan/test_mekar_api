package utils

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"result"`
}

func GenerateResponse(status int, message string, data interface{}) Response {
	var responseProduct Response
	responseProduct.Status = status
	responseProduct.Message = message
	responseProduct.Data = data
	return responseProduct
}
