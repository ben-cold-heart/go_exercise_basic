package main

import (
	"fmt"
	"os"
)

// ขียนโปรแกรมเพื่อ:
// 1. สร้างไฟล์ชื่อ `data.txt`
// 2. เขียนข้อความ **"This is a test file."** ลงในไฟล์
// 3. อ่านไฟล์และพิมพ์ข้อความในไฟล์ออกมาทางหน้าจอ

func main() {

	// Create and Write File
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	file.WriteString("This is a test file.")
	defer file.Close()

	fileRead, err := os.ReadFile("data.txt")

	// Read File
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(fileRead))

}
