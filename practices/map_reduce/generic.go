package map_reduce

import (
	"fmt"
	"reflect"
	"strings"
)

// 泛型MapReduce ==> 编译通过，运行时报错

// 通过 reflect.ValueOf() 来获得 interface{} 的值，其中一个是数据 vdata，另一个是函数 vfn，
// 然后通过 vfn.Call() 方法来调用函数，通过 []reflect.Value{vdata.Index(i)}来获得数据。
func Map(data interface{}, fn interface{}) []interface{} {
	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	result := make([]interface{}, vdata.Len())
	for i := 0; i < vdata.Len(); i++ {
		result[i] = vfn.Call([]reflect.Value{vdata.Index(i)})[0].Interface()
	}
	return result
}

func main2() {
	square := func(x int) int {
		return x * x
	}
	nums := []int{1, 2, 3, 4}
	squared_arr := Map(nums, square)
	fmt.Println(squared_arr)
	//[1 4 9 16]
	upcase := func(s string) string {
		return strings.ToUpper(s)
	}
	strs := []string{"Hao", "Chen", "MegaEase"}
	upstrs := Map(strs, upcase)
	fmt.Println(upstrs)
	//[HAO CHEN MEGAEASE]
}
