package client

type Status int

const (
	Active   Status = 0
	Inactive Status = 1
	All      Status = 2
)

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

type Subscriptions struct {
	CurrentPage uint           `json:"current_page"`
	Data        []Subscription `json:"data"`
	FirstPage   string         `json:"first_page"`
	From        uint           `json:"from"`
	LastPage    uint           `json:"last_page"`
	LastPageURL string         `json:"last_page_url"`
	NextPageURL string         `json:"next_page_url"`
	Path        string         `json:"path"`
	PerPage     uint           `json:"per_page"`
	PrevPageURL string         `json:"prev_page_url"`
	To          uint           `json:"to"`
	Total       uint           `json:"total"`
}
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
