package heap

type MinHeap []int

func (h *MinHeap) Push(element int) {
	*h = append(*h, element)
	insertedElementIndex := len(*h) - 1
	h.heapifyUp(insertedElementIndex)
}

func (h *MinHeap) Pop() int {
	if len(*h) == 0 {
		return -1
	}

	removedElement := (*h)[0]
	lastElementIndex := len(*h) - 1

	(*h)[0] = (*h)[lastElementIndex]

	*h = (*h)[:lastElementIndex]

	h.heapify(0)

	return removedElement
}

func (h *MinHeap) heapifyUp(index int) {
	for (*h)[parent(index)] > (*h)[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) heapify(index int) {
	smallest, left, right := index, leftChild(index), rightChild(index)

	if left < len(*h) && (*h)[left] < (*h)[smallest] {
		smallest = left
	}

	if right < len(*h) && (*h)[right] < (*h)[smallest] {
		smallest = right
	}

	if smallest != index {
		h.swap(smallest, index)
		h.heapify(smallest)
	}
}

func (h *MinHeap) swap(firstIndex, secondIndex int) {
	(*h)[firstIndex], (*h)[secondIndex] = (*h)[secondIndex], (*h)[firstIndex]
}
