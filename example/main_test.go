package example

import "testing"

func TestTestable(t *testing.T) {
	if err := testable(); err != nil {
		t.Fatalf("Error returned: %s", err)
	}
}