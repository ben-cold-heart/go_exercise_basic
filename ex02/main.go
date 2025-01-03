package main

import (
	"fmt"
)

// สร้างโปรแกรมที่มีตัวแปร name และ age
// พิมพ์ข้อความ "My name is [name] and I am [age] years old." ออกมา
func Greeting(name string, age int) {
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)
}

func main() {
	// Not Using Function
	var name string = "John"
	var age int = 30
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Using Function
	Greeting("John", 30)
}
