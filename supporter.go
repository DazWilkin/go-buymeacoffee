package client

type Supporters struct {
	CurrentPage uint        `json:"current_page"`
	Data        []Supporter `json:"data"`
	From        uint        `json:"from"`
	LastPage    uint        `json:"last_page"`
	LastPageURL string      `json:"last_page_url"`
	NextPageURL string      `json:"next_page_url"`
	Path        string      `json:"path"`
	PerPage     uint        `json:"per_page"`
	PrevPageURL string      `json:"prev_page_url"`
	To          uint        `json:"to"`
	Total       uint        `json:"total"`
}
type Supporter struct {
	ID            uint   `json:"support_id"`
	Node          string `json:"support_note"`
	Coffees       uint   `json:"coffees"`
	TransactionID string `json:"transaction_id"`
	Name          string `json:"supporter_name"`
	Email         string `json:"supporter_email"`
}
