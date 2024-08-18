package heap

import (
	"testing"
)

// Helper function to create a MaxHeap with some initial values
func createHeapForMaxHeap(values []int) *MaxHeap {
	h := &MaxHeap{}
	for _, value := range values {
		h.Push(value)
	}
	return h
}

// TestMaxHeap tests the MaxHeap implementation.
func TestMaxHeap(t *testing.T) {
	// Create a new heap and add some elements.
	h := createHeapForMaxHeap([]int{5, 3, 8, 1, 2})

	// Check the maximum element after initial inserts.
	if max := (*h)[0]; max != 8 {
		t.Errorf("expected maximum to be 8, got %d", max)
	}

	// Push a new element and verify the maximum element.
	h.Push(12)
	if max := (*h)[0]; max != 12 {
		t.Errorf("expected maximum to be 12, got %d", max)
	}

	// Pop elements and check their order.
	expected := []int{12, 8, 5, 3, 2, 1}
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
func TestHeapifyForMaxHeap(t *testing.T) {
	// Create a new slice with initial unsorted values.
	arr := []int{6, 8, 12, 7, 1}
	h := BuildMaxHeap(&arr)

	// The heap should now be a min-heap.
	expected := []int{12, 8, 6, 7, 1}
	for i, e := range expected {
		if got := h[i]; got != e {
			t.Errorf("expected %d at index %d, got %d", e, i, got)
		}
	}
}

// TestEdgeCases tests edge cases for MaxHeap operations.
func TestEdgeCasesForMaxHeap(t *testing.T) {
	h := &MaxHeap{}

	// Push and pop a single element
	h.Push(10)
	if max := (*h)[0]; max != 10 {
		t.Errorf("expected maximum to be 10, got %d", max)
	}

	if popped := h.Pop(); popped != 10 {
		t.Errorf("expected 10, got %d", popped)
	}

	// The heap should now be empty
	if popped := h.Pop(); popped != -1 {
		t.Errorf("expected -1 when popping from an empty heap, got %d", popped)
	}
}
