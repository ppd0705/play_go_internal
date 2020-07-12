package main

import (
	"fmt"
)

//func main() {
	//test2()
	//test4()
	//test5()
//}

func test1() {
	var array [10]int

	var slice = array[5:6]

	fmt.Println("length of slice: ", len(slice))
	fmt.Println("capacity of slice: ", cap(slice))
	fmt.Println(&slice[0] == &array[5])
}

func addElement(slice []int, e int) []int {
	return append(slice, e)
}

func test2() {
	var slice []int
	slice = append(slice, 1, 2, 3, 4)
	newSlice := addElement(slice, 5)
	fmt.Println(&slice[0] == &newSlice[0])
}

func test3() {
	orderLen := 5
	order := make([]uint16, 2*orderLen)

	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]

	fmt.Printf("pollorder len %d cap %d\n", len(pollorder), cap(pollorder))
	fmt.Printf("lockorder len %d cap %d\n", len(lockorder), cap(lockorder))
}

// shell:  GOSSAFUNC=newSlice go build play_slice.go
func newSlice() []int {
	arr := [3]int{1, 2, 3}
	slice := arr[0:1]
	return slice
}

func newSlice2() []int {
	slice := []int{1, 2, 3}
	return slice
}

func test4() {
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	newSlice := addElement(slice, 5)
	fmt.Println(&slice[0] == &newSlice[0])
	println(cap(slice), cap(newSlice))
}

func test5() {
	var slice []int
	slice = append(slice, 1,2,3,4, 5)
	println(cap(slice))
}

func test6() {
	var slice []int
	slice = append(slice, 1)
	println(cap(slice))

	slice = append(slice, 1)
	println(cap(slice))

	slice = append(slice, 1)
	println(cap(slice))

	slice = append(slice, 1)
	println(cap(slice))

	slice = append(slice, 1)
	println(cap(slice))
}