package client

import (
	"encoding/json"
	"testing"
)

// TestExampleSupporters tests whether the API documentation's example supporters unmarshals correctly to the Supporters type
func TestExampleSupporters(t *testing.T) {
	supporters := &Supporters{}
	if err := json.Unmarshal(exampleSupporters, supporters); err != nil {
		t.Fatal(err)
	}
}

// TestExampleSupporter tests whether the API documentation's example supporter unmarshals correctly to the Supporter type
func TestExampleSupporter(t *testing.T) {
	supporter := &Supporter{}
	if err := json.Unmarshal(exampleSupporter, supporter); err != nil {
		t.Fatal(err)
	}
}
