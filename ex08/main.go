package main

import (
	"fmt"
)

// เขียนโปรแกรมที่รับค่าตัวเลข 2 ตัวจากผู้ใช้ และให้ผู้ใช้เลือกว่าจะ:
// - บวก (+)
// - ลบ (-)
// - คูณ (*)
// - หาร (/)
// แล้วแสดงผลลัพธ์ตามตัวเลือกของผู้ใช้

type Number struct {
	num1 int
	num2 int
}

func (n *Number) add() int {
	result := n.num1 + n.num2
	fmt.Println("Add Result:", result)
	return result
}

func (n *Number) subtract() int {
	result := n.num1 - n.num2
	fmt.Println("Subtract Result:", result)
	return result
}

func (n *Number) multiplier() int {
	result := n.num1 * n.num2
	fmt.Println("Multiplier result:", result)
	return result
}

func (n *Number) divide() int {
	result := n.num1 / n.num2
	fmt.Println("Divide Result:", result)
	return result
}

func main() {

	// Input Number
	var num1, num2 int
	fmt.Printf("Enter First Numbers: ")
	fmt.Scan(&num1)
	fmt.Printf("Enter Second Numbers: ")
	fmt.Scan(&num2)

	n := Number{num1: num1, num2: num2}

	// Add
	n.add()

	// Subtract
	n.subtract()

	// Multiplier
	n.multiplier()

	// Divide
	n.divide()

}
