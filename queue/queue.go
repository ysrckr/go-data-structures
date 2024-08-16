package queue

type Queue []string

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(str string) {
	*q = append(*q, str)
}

func (q *Queue) Dequeue() (string, bool) {
	if q.IsEmpty() {
		return "", false
	}

	index := 0
	firstElement := (*q)[index]
	*q = (*q)[index+1:]

	return firstElement, true
}
