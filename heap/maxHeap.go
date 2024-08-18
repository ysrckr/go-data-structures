package heap

type MaxHeap []int

func (h *MaxHeap) Push(element int) {
	*h = append(*h, element)
	h.heapifyUp()
}

func (h *MaxHeap) Pop() int {
	if len(*h) == 0 {
		return -1
	}

	removed := (*h)[0]
	lastElement := len(*h) - 1

	(*h)[0] = (*h)[lastElement]

	*h = (*h)[:lastElement]

	h.heapify(0)

	return removed
}

func (h *MaxHeap) swap(first, second int) {
	(*h)[first], (*h)[second] = (*h)[second], (*h)[first]
}

func (h *MaxHeap) heapifyUp() {
	lastIndex := len(*h) - 1

	for (*h)[parent(lastIndex)] < (*h)[lastIndex] {
		h.swap(parent(lastIndex), lastIndex)
		lastIndex = parent(lastIndex)
	}
}

func (h *MaxHeap) heapify(index int) {
	length := len(*h)
	largest, left, right := index, leftChild(index), rightChild(index)

	if left < length && (*h)[left] > (*h)[largest] {
		largest = left
	}

	if right < length && (*h)[right] > (*h)[largest] {
		largest = right
	}

	if largest != index {
		h.swap(largest, index)
		h.heapify(largest)
	}
}
