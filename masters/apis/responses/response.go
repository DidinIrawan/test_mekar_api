package responses

type ResponsesData struct {
	Status int         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ResponsesStatus struct {
	Status int    `json:"status"`
	Message    string `json:"message"`
}
