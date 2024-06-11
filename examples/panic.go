package examples

import "os"

func Panic_() {

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
	// 相当于throw Exception
	panic("a problem")
}
