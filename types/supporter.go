package types

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

// Supporters is a type that represents the JSON type returned by the API's supporters method
type Supporters struct {
	Page
	Data []Supporter `json:"data"`
}

// SupportersToText is a method that converts a slice of Supporters into tabbed output
func SupportersToText(ss []Supporter) string {
	var b bytes.Buffer
	w := tabwriter.NewWriter(&b, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tEmail")
	for _, s := range ss {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n",
			s.ID,
			s.Name,
			s.Email,
			s.TransactionID,
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
