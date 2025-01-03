package main

import (
	"fmt"
	"math/rand"
	"time"
)

// สร้างโปรแกรมที่:
// 1. สุ่มตัวเลข 1-100
// 2. ให้ผู้ใช้ทายตัวเลข จนกว่าจะถูกต้อง
// 3. แสดงข้อความว่า **"ถูกต้องแล้ว!"** เมื่อผู้ใช้ทายถูก และจำนวนครั้งที่ทาย
func randomNumber(min int, max int) (r int) {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(max-min+1) + min
	return randomNum
}

func main() {
	r := randomNumber(1, 100)
	for {
		var guestNum int
		fmt.Println("Guest Number :")
		fmt.Scan(&guestNum)

		if guestNum == r {
			fmt.Println("ถูกต้องแล้ว! คุณทายครั้งที่", guestNum)
			break
		} else {
			fmt.Println("ไม่ถูกต้อง!")
		}
	}

}
