package models

type ResponseQuery struct {
	Id             int    `json:"id"`
	Country        string `json:"country"`
	CreditCardType string `json:"credit_card_type"`
	CreditCard     int    `json:"credit_card"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
}

type RequestData struct {
	Id int `json:"id"`
}

type ResponseServices struct {
	Data interface{} `json:"data"`
}

type ResponseHandler struct {
	Message string      `json:"message"`
	Items   interface{} `json:"items"`
}
