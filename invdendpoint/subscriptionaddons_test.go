package invdendpoint

import (
	"encoding/json"
	"testing"
)

func TestUnMarshalSubscriptionAddonObject(t *testing.T) {
	s := `{
    "id": 3,
    "catalog_item": "delivery",
    "plan" : "test-plan",
    "quantity": 11,
    "created_at": 1420391704
}`

	so := new(SubscriptionAddon)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

}
