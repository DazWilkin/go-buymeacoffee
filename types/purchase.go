package types

type Purchases struct {
	Page
	Data []Purchase `json:"data"`
}
type Purchase struct {
	ID         uint   `json:"purchase_id"`
	Purchased  string `json:"purchased_on"`
	Updated    string `json:"purchase_updated_on"`
	IsRevoked  uint   `json:"purchase_is_revoked"`
	Amount     string `json:"purchase_amount"`
	Currency   string `json:"purchase_currency"`
	Question   string `json:"purchase_question"`
	PayerEmail string `json:"purchase_payer_email"`
	PayerName  string `json:"purchase_payer_name"`
	Extra      Reward `json:"extra"`
}
type Reward struct {
	ID                  uint   `json:"reward_id"`
	Title               string `json:"reward_title"`
	Description         string `json:"reward_description"`
	ConfirmationMessage string `json:"reward_confirmation_message"`
	Question            string `json:"reward_question"`
	Used                uint   `json:"reward_use"`
	Created             string `json:"reward_created_on"`
	Updated             string `json:"reward_updated_on"`
	Deleted             string `json:"reward_deleted_on"`
	IsActive            uint   `json:"reward_is_active"`
	Image               string `json:"reward_url"`
	Slots               uint   `json:"reward_slots"`
	CoffeePrice         string `json:"reward_coffee_price"`
	Order               uint   `json:"reward_order"`
}
