package example_docs

import "testing"

func TestDocs(t *testing.T) {
	res := docs()
	if res != 4 {
		t.Fatalf("Error, expected 1 but got %d", res)
	}
}
