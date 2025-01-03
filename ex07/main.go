package main

import "fmt"

// สร้าง Struct ชื่อ `Person` ที่มีฟิลด์ `Name` และ `Age`:
// - เพิ่ม Method ชื่อ `Introduce` เพื่อพิมพ์ข้อความ **"Hi, I am [Name] and I am [Age] years old."**
// - ทดลองเรียกใช้ Method นี้ใน `main`

type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() {
	fmt.Printf("Hi, I am %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	p := Person{"John", 30}
	p.Introduce()
}
