package heap

import (
	"testing"
)

// Helper function to create a MinHeap with some initial values
func createHeap(values []int) *MinHeap {
	h := &MinHeap{}
	for _, value := range values {
		h.Push(value)
	}
	return h
}

// TestMinHeap tests the MinHeap implementation.
func TestMinHeap(t *testing.T) {
	// Create a new heap and add some elements.
	h := createHeap([]int{5, 3, 8, 1, 2})

	// Check the minimum element after initial inserts.
	if min := (*h)[0]; min != 1 {
		t.Errorf("expected minimum to be 1, got %d", min)
	}

	// Push a new element and verify the minimum element.
	h.Push(0)
	if min := (*h)[0]; min != 0 {
		t.Errorf("expected minimum to be 0, got %d", min)
	}

	// Pop elements and check their order.
	expected := []int{0, 1, 2, 3, 5, 8}
	for _, e := range expected {
		if top := h.Pop(); top != e {
			t.Errorf("expected %d, got %d", e, top)
		}
	}

	// Test popping from an empty heap.
	if popped := h.Pop(); popped != -1 {
		t.Errorf("expected -1 when popping from an empty heap, got %d", popped)
	}
}

// TestHeapify ensures the heapify operations maintain the heap property.
func TestHeapify(t *testing.T) {
	// Create a new slice with initial unsorted values.
	arr := []int{4, 10, 3, 5, 1}
	h := BuildMinHeap(&arr)

	// The heap should now be a min-heap.
	expected := []int{1, 4, 3, 5, 10}
	for i, e := range expected {
		if got := h[i]; got != e {
			t.Errorf("expected %d at index %d, got %d", e, i, got)
		}
	}
}

// TestEdgeCases tests edge cases for MinHeap operations.
func TestEdgeCases(t *testing.T) {
	h := &MinHeap{}

	// Push and pop a single element
	h.Push(10)
	if min := (*h)[0]; min != 10 {
		t.Errorf("expected minimum to be 10, got %d", min)
	}

	if popped := h.Pop(); popped != 10 {
		t.Errorf("expected 10, got %d", popped)
	}

	// The heap should now be empty
	if popped := h.Pop(); popped != -1 {
		t.Errorf("expected -1 when popping from an empty heap, got %d", popped)
	}
}
