package queue

type Queue []int

func (queue *Queue) Push(num int) {
	*queue = append(*queue, num)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
