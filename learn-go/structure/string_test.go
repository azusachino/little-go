package structure

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	try_modify_string_slice()
	try_modify_string_pointer()
}

func TestNano(t *testing.T) {
	var s = "中国欢迎你"
	nonascii(&s)
}

func try_modify_string_slice() {

	s := "hello"

	arr := []rune(s)
	arr[0] = 't'

	fmt.Println(string(arr))
	fmt.Println(s)
}

func try_modify_string_pointer() {
	var s = "hello"
	fmt.Println("prev string is ", s)

	ModifyString(&s)
	fmt.Println("after string is ", s)
}

func ModifyString(s *string) {

	fmt.Printf("pointer of string: %p\n", s)
	// get string raw pointer of size 8
	p := (*uintptr)(unsafe.Pointer(s))

	fmt.Printf("unsafe raw pointer of string: %p\n", p)

	// get base byte array pointer
	var arr *[5]byte = (*[5]byte)(unsafe.Pointer(p))

	fmt.Printf("raw pointer of base array: %p\n", arr)

	var len *int = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Sizeof((*uintptr)(nil))))

	for i := 0; i < (*len); i++ {
		fmt.Printf("%p ==> %c\n", &((*arr)[i]), (*arr)[i])
		p1 := &((*arr)[i])
		v := (*p1)
		(*p1) = v + 1
	}
}

func nonascii(s *string){
	rs := []rune(*s)
	sl := []byte(*s)

	for i,v := range rs {
		var utf8Bytes []byte
		for j:=i*3; j < (i+1)*3;j++ {
			utf8Bytes = append(utf8Bytes, sl[j])
		}
		fmt.Printf("%s -> %X -> %X\n", string(v), v, utf8Bytes)
	}
}
