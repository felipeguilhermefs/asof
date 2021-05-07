package set

import "testing"

// TestSet general set test
func TestSet(t *testing.T) {
	set := Set{}

	if set.Contains(1) {
		t.Fatalf("Should not contain values")
	}

	if set.Len() != 0 {
		t.Fatalf("Should be empty")
	}

	set.Add(1)
	set.Add(1)
	set.Add(1)
	set.Add(1)
	if !set.Contains(1) {
		t.Fatalf("Should contain 1")
	}

	if set.Len() != 1 {
		t.Fatalf("Should contain only 1 item")
	}

	set.Add(2)
	set.Add(2)
	set.Add(3)
	set.Add(4)

	if set.Len() != 4 {
		t.Fatalf("Should contain 4 items")
	}
}
