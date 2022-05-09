package types

// Supporters is a type that represents the JSON type returned by the API's supporters method
type Supporters struct {
	Page
	Data []Supporter `json:"data"`
}

// Supporter is a type that represents the JSON type returned by the API's supporter method
type Supporter struct {
	ID            uint   `json:"support_id"`
	Node          string `json:"support_note"`
	Coffees       uint   `json:"coffees"`
	TransactionID string `json:"transaction_id"`
	Name          string `json:"supporter_name"`
	Email         string `json:"supporter_email"`
}
