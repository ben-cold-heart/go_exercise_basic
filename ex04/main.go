package main

import (
	"fmt"
)

// สร้างฟังก์ชันชื่อ add ที่รับค่าตัวเลข 2 ตัว และคืนค่าผลรวม จากนั้นเรียกใช้ฟังก์ชันใน main
func sum(a, b int) int {
	return a + b
}

func main() {
	result := sum(5, 2)
	fmt.Println("Add Result :", result)
}
