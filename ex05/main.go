package main

import (
	"fmt"
)

func main() {
	// slice เริ่มต้น
	sc1 := []int{1, 2, 3, 4, 5}
	fmt.Println("start", sc1)

	sc2 := make([]int, 0, 5)
	fmt.Println("start", sc2)

	// Add value into slice
	sc1 = append(sc1, 6, 7, 8, 9)
	fmt.Println("Add", sc1)

	sc2 = append(sc2, 1, 2, 3, 4, 5)
	fmt.Println("Add", sc2)

	// Delete Value from slice
	sc1 = append(sc1[:2], sc1[3:]...)
	sc2 = append(sc2[:2], sc2[3:]...)
	fmt.Println("Delete", sc1)
	fmt.Println("Delete", sc2)
}
