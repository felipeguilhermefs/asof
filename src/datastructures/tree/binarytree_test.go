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

	t.Run("should traverse In Order", func(t *testing.T) {
		testData := []int{2, 1, 4, 3, 5}

		bt := BinaryTree{}

		for _, data := range testData {
			bt.Insert(data)
		}

		visited := []int{}

		bt.Traverse(InOrder, func(value int) {
			visited = append(visited, value)
		})

		expected := []int{1, 2, 3, 4, 5}
		for i := range expected {
			if visited[i] != expected[i] {
				t.Fatalf("Unexpected: index= %d expected= %v visited= %v", i, expected, visited)
			}
		}
	})

	t.Run("should traverse Pre Order", func(t *testing.T) {
		testData := []int{2, 1, 4, 3, 5}

		bt := BinaryTree{}

		for _, data := range testData {
			bt.Insert(data)
		}

		visited := []int{}

		bt.Traverse(PreOrder, func(value int) {
			visited = append(visited, value)
		})

		expected := []int{2, 1, 4, 3, 5}
		for i := range expected {
			if visited[i] != expected[i] {
				t.Fatalf("Unexpected: index= %d expected= %v visited= %v", i, expected, visited)
			}
		}
	})

	t.Run("should traverse Post Order", func(t *testing.T) {
		testData := []int{2, 1, 4, 3, 5}

		bt := BinaryTree{}

		for _, data := range testData {
			bt.Insert(data)
		}

		visited := []int{}

		bt.Traverse(PostOrder, func(value int) {
			visited = append(visited, value)
		})

		expected := []int{1, 3, 5, 4, 2}
		for i := range expected {
			if visited[i] != expected[i] {
				t.Fatalf("Unexpected: index= %d expected= %v visited= %v", i, expected, visited)
			}
		}
	})
}
