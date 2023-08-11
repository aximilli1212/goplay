package queue

type Queue struct {
	items    []int
	capacity int
}

func New(capacity int) Queue {

}

func (q *Queue) Append(item int) bool {
	if len(q.items) == q.capacity {
		return false
	}

	q.items = append(q.items, item)
	return true
}
