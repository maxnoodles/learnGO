package queue

type Queue []interface{}

// Pushes the element into the queue
//     e.g. q.push(123)
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}
