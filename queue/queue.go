package queue

type Queue []string

func (q *Queue) IsEmpty() {
	return len(*q) == 0
}
