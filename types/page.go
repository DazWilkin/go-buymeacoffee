package types

// Page is a type that represents the paging metadata
// Page is embeded in each of the "plural" types
// Purchases, Subscriptions, and Supporters
type Page struct {
	Current uint   `json:"current_page"`
	First   string `json:"first_page"`
	From    uint   `json:"from"`
	Last    uint   `json:"last_page"`
	LastURL string `json:"last_page_url"`
	NextURL string `json:"next_page_url"`
	Path    string `json:"path"`
	PerPage uint   `json:"per_page"`
	PrevURL string `json:"prev_page_url"`
	To      uint   `json:"to"`
	Total   uint   `json:"total"`
}
