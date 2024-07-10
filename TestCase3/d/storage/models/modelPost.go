package models

type Data struct {
	Id             int    `json:"id"`
	Country        string `json:"country"`
	CreditCardType string `json:"credit_card_type"`
	CreditCard     int    `json:"credit_card"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
}

type RequestPost struct {
	Id             int    `json:"id"`
	Country        string `json:"country"`
	CreditCardType string `json:"credit_card_type"`
	CreditCard     int    `json:"credit_card"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
}

type ResponsePost struct {
	Id             int    `json:"id"`
	Country        string `json:"country"`
	CreditCardType string `json:"credit_card_type"`
	CreditCard     int    `json:"credit_card"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
}

type ResponseData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
