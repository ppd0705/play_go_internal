package main

import "fmt"

type Priority int

const (
	LogErr Priority = iota
	LogInfo
	LogDebug
)

const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	muteWaiterShift = iota
)

const (
	bit0, mask0 = 1 << iota, 1 << iota-1
	bit1, mask1
	_, _
	bit3, mask3
)

func main() {
	fmt.Printf("Priority %v %v %v\n", LogErr, LogInfo, LogDebug)

	fmt.Printf("mutex %d %d %d %d\n", mutexLocked, mutexWoken,mutexStarving, muteWaiterShift)

	fmt.Printf("bit mask\n %d %d\n %d %d\n %d %d\n", bit0, mask0, bit1, mask1, bit3, mask3)
}
