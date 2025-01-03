package main

import (
	"fmt"
)

// สร้างโปรแกรมที่เก็บข้อมูลชื่อและอายุใน Map โดย:
// - เพิ่มค่าชื่อ-อายุอย่างน้อย 3 รายการ
// - แสดงอายุของชื่อที่ระบุ เช่น ถ้าใส่ชื่อ **"John"** ให้แสดงอายุของ John

func showAge(dic map[string]int, name string) {
	fmt.Println(name, "is", dic[name], "year olds.")
}

func main() {
	// Create Map
	// empty
	dic := make(map[string]int)

	// with KV
	dic2_ex06 := map[string]int{
		"John": 20,
		"Jane": 25,
		"Bob":  30,
	}

	// Add KV to Map
	dic["John"] = 30

	// Edit KV
	dic["John"] = 999

	// delete KV
	delete(dic, "John")

	fmt.Println(dic)
	fmt.Println(dic2_ex06)

	showAge(dic2_ex06, "John")
}
