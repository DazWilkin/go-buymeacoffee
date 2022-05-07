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
	Data []Subscription `json:"data"`
}
type Subscription struct {
	ID uint `json:"subscription_id"`
}
