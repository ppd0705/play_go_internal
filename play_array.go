package main

// shell: GOSSAFUNC=outOfRange go build play_array.go
func outOfRange() int {
	arr := [3]int{1, 2, 3}
	i := 4
	elem := arr[i]
	return elem
}
