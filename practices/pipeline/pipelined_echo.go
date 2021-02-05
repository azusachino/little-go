package pipeline

type EchoFunc func([]int) <-chan int
type PipeFunc func(<-chan int) <-chan int

func pipeline(nums []int, echo EchoFunc, pipeFns ...PipeFunc) <-chan int {
	ch := echo(nums)
	for i := range pipeFns {
		ch = pipeFns[i](ch)
	}
	return ch
}

//var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//for n := range pipeline(nums, gen, odd, sq, sum) {
//fmt.Println(n)
//}
