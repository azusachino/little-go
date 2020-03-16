package queue

// interface == any(TS)
type Queue []interface{}

func (q *Queue) Push(num interface{}) {
	*q = append(*q, num)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
