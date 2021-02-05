package examples

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() float64 {
	return float64(r.width * r.height)
}

func (r *rect) perimeter() float64 {
	return float64(2 * (r.height + r.width))
}

func main() {
	r := rect{width: 10, height: 10}
	fmt.Println(r.area(), r.perimeter())

	rp := &r

	fmt.Println(rp.perimeter(), rp.area())
}
