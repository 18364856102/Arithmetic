package model

// ArithmeticRequest define request struct
type ArithmeticRequest struct {
	RequestType string `json:"requestType"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}

// ArithmeticResponse define response struct
type ArithmeticResponse struct {
	Result int    `json:"result"`
	Error  string `json:"error"`
}
