package types

// Status is a type that represents the possible values for the API's Subscriptions status
type Status int

const (
	Active   Status = 0
	Inactive Status = 1
	All      Status = 2
)

// String is a function that converts a Status into a string
func (s Status) String() string {
	switch s {
	case Active:
		return "active"
	case Inactive:
		return "inactive"
	default:
		return "all"
	}
}

// Subscriptions is a type that represents the JSON type returned by the API's subscriptions method
type Subscriptions struct {
	Page
	Data []Subscription `json:"data"`
}

// Subscription is a type that represents the JSON type returned by the API's subscription method
type Subscription struct {
	ID                 uint   `json:"subscription_id"`
	Cancelled          string `json:"subscription_cancelled_on"`
	Created            string `json:"subscription_created_on"`
	Updated            string `json:"subscription_updated_on"`
	CurrentPeriodStart string `json:"subscription_current_period_start"`
	CurrentPeriodEnd   string `json:"subscription_current_period_end"`
	CoffeePrice        string `json:"subscription_coffee_price"`
	CoffeeNum          uint   `json:"subscription_coffee_num"`
	TransactionID      string `json:"subscription_transaction_id"`
	PayerEmail         string `json:"subscription_payer_email"`
	PayerName          string `json:"subscription_payer_name"`
}
