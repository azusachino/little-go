package examples

import "fmt"

func init() {

	s := make([]string, 3)
	fmt.Println("empty", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set", s)
	fmt.Println("get", s[1])

	fmt.Println("len", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append", s)

	c := make([]string, len(s))

	copy(c, s)
	fmt.Println("copy", c)

	l := s[2:5]
	fmt.Println("slice", l)

	l = s[:5]

	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	twoDi := make([][]int, 3)

	for i := 0; i < 3; i++ {

		innerLen := i + 1
		twoDi[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDi[i][j] = i + j
		}
	}
	fmt.Println("2d", twoDi)
}
