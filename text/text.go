package text

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/DazWilkin/go-buymeacoffee/types"
)

// Purchases is a method that converts a slice of Packages into tabbed output
func Purchases(pp []types.Purchase) string {
	var b bytes.Buffer
	w := tabwriter.NewWriter(&b, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "ID")
	for _, p := range pp {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n",
			p.ID,
			p.PayerName,
			p.PayerEmail,
			p.Currency,
			p.Amount)
	}
	if err := w.Flush(); err != nil {
		// TODO(dazwilkin) Avoid CWE-703 unhandled error
		return ""
	}

	return b.String()
}

// Subscriptions is a method that converts a slice of Supporters into tabbed output
func Subscriptions(ss []types.Subscription) string {
	var b bytes.Buffer
	w := tabwriter.NewWriter(&b, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tEmail")
	for _, s := range ss {
		fmt.Fprintf(w, "%d\t%s\t%s\n",
			s.ID,
			s.PayerName,
			s.PayerEmail,
		)
	}
	if err := w.Flush(); err != nil {
		// TODO(dazwilkin) Avoid CWE-703 unhandled error
		return ""
	}

	return b.String()
}

// Supporters is a method that converts a slice of Supporters into tabbed output
func Supporters(ss []types.Supporter) string {
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
