package heap

type MinHeap []int

func (h *MinHeap) swap(firstIndex, secondIndex int) {
	(*h)[firstIndex], (*h)[secondIndex] = (*h)[secondIndex], (*h)[firstIndex]
}

func (h *MinHeap) Insert(element int) {
	*h = append(*h, element)
	insertedElementIndex := len(*h) - 1
	h.heapifyUp(insertedElementIndex)
}

func (h *MinHeap) Remove() int {
	removedElement := (*h)[0]
	lastElementIndex := len(*h) - 1

	if len(*h) == 0 {
		return -1
	}

	(*h)[0] = (*h)[lastElementIndex]

	*h = (*h)[:lastElementIndex]

	h.heapifyDown(0)

	return removedElement
}

func (h *MinHeap) heapifyUp(index int) {
	for (*h)[parent(index)] > (*h)[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) heapifyDown(index int) {
	lastIndexToCheck := len(*h) - 1
	left, right := leftChild(index), rightChild(index)
	childToCompare := 0

	// find the smallest of the left anf right children
	for left <= lastIndexToCheck {
		if left == lastIndexToCheck {
			childToCompare = left
		} else if (*h)[left] < (*h)[right] {
			childToCompare = left
		} else {
			childToCompare = right
		}

		// swap the parent with the smallest child
		if (*h)[index] > (*h)[childToCompare] {
			h.swap(index, childToCompare)

			// update the indexes that are being compared
			index = childToCompare

			left, right = leftChild(index), rightChild(index)
		} else {
			return
		}
	}
}
