package types

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

// Supporters is a type that represents the JSON type returned by the API's supporters method
type Supporters struct {
	CurrentPage uint        `json:"current_page"`
	Data        []Supporter `json:"data"`
	FirstPage   string      `json:"first_page"`
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

// Text is a method that converts a slice of Supporters into tabbed output
func (s *Supporters) Text() string {
	var b bytes.Buffer
	w := tabwriter.NewWriter(&b, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tEmail")
	for _, supporter := range s.Data {
		fmt.Fprintf(w, "%d\t%s\t%s\n",
			supporter.ID,
			supporter.Name,
			supporter.Email,
		)
	}
	if err := w.Flush(); err != nil {
		// TODO(dazwilkin) Avoid CWE-703 unhandled error
		return ""
	}

	return b.String()
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
