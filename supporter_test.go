package client

import (
	"encoding/json"
	"testing"
)

func TestExampleSupporters(t *testing.T) {
	supporters := &Supporters{}
	if err := json.Unmarshal(exampleSupporters, supporters); err != nil {
		t.Fatal(err)
	}
}
func TestExampleSupporter(t *testing.T) {
	supporter := &Supporter{}
	if err := json.Unmarshal(exampleSupporter, supporter); err != nil {
		t.Fatal(err)
	}
}
