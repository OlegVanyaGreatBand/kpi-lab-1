package example_tests

import "testing"

func TestTestable(t *testing.T) {
	res := testable()
	if res != 2 {
		t.Fatalf("Error, expected 4 but got %d", res)
	}
}
