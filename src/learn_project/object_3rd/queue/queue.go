package queue

type Queue []int

func (q *Queue) Push(num int) {
	*q = append(*q, num)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
