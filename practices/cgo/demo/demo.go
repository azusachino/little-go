package demo

// // #include <stdio.h>
// // #include <stdlib.h>
// /*
// void print(char *s){
// printf("print used by C: %s\n", s);
// };
// void SayHello(const char* s);
// */
// import "C"
// import "unsafe"

// // 代码通过import "C"语句启用CGO特性，紧邻这行语句前面注释是一种特殊语法，里面包含的是正常的C语言代码。
// func main() {
// 	s := "Hello"
// 	cs := C.CString(s)
// 	defer C.free(unsafe.Pointer(cs))
// 	C.print(cs)
// 	C.SayHello(C.CString("Hello World\n"))
// }
