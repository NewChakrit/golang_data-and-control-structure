package main

import "fmt"

func main() {
	//การประกาศ Variable มี 4 วิธี
	// 1. ใช้ var ประกาศแบบไม่กำหนดค่าเริ่มต้น
	//var fullName string
	//var age int

	// 2. ใช้ var ประกาศแบบกำหนดค่าเริ่มต้น
	var fullName = "New"
	var age = 27

	fmt.Println("Hello ", fullName)
	fmt.Println("I'm ", age)
	fmt.Printf("Hi, My name is %s! and I'm %d years old\n", fullName, age)

	// 3. ประกาศแบบไม่ประกาศ type แต่มีค่าเริ่มต้น
	var myFriend = "Nick"
	fmt.Printf("My friend is %s, Do you know him?\n", myFriend)

	// 4. shorthand
	yourBro := "Mike"
	fmt.Printf("Is your brother is %s ?\n", yourBro)

	// การเปลี่ยนแปลงค่าให้ใช้ =
	yourBro = "Mos"
	fmt.Println("No, My brother is ", yourBro)

	// ไม่สามารถประกาศชื่อตัวแปลซ้ำได้ใน function เดียวกัน เช่นเดี่ยวกับการเปลี่ยน type ของ variable
}
