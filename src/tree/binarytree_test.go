package tree

import "testing"

// TestBinaryTree
func TestBinaryTree(t *testing.T) {
	t.Run("should store and find data", func(t *testing.T) {
		testData := []int{2, 1, 4, 3, 5}

		bt := BinaryTree{}

		for _, data := range testData {
			bt.Insert(data)
		}

		for _, data := range testData {
			if !bt.Contains(data) {
				t.Fatalf("Inserted data not found: data %d", data)
			}
		}

		if bt.Contains(6) {
			t.Fatal("Found data that it shouldn't")
		}
	})
}
