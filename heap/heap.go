package heap

func BuildMinHeap(arr *[]int) MinHeap {
	var minHeap MinHeap
	minHeap = *arr
	length := len(minHeap)

	for i := length/2 - 1; i >= 0; i-- {
		minHeap.heapify(i)
	}

	return minHeap
}

func BuildMaxHeap(arr *[]int) MaxHeap {
	var maxHeap MaxHeap
	maxHeap = *arr
	length := len(maxHeap)

	for i := length/2 - 1; i >= 0; i-- {
		maxHeap.heapify(i)
	}

	return maxHeap
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftChild(index int) int {
	return index*2 + 1
}

func rightChild(index int) int {
	return index*2 + 2
}
