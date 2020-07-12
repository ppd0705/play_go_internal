package main

import (
	"fmt"
	"unsafe"
)

type myStruct struct {
	myFloat float64
	myBool  bool
	myBool2  bool
	myBool3  bool
	myBool4  bool
	myInt   int32
	//b []byte
	//c [2]byte
	//d []int64
	e map[string]int64
}

func main() {
	//_ = add(3, 5)
	a := myStruct{}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(map[string]int64{}))
}

func add(a, b int) int {

	return a + b
}
