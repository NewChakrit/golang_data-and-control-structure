package control_structure

import "fmt"

func ControlStructure() {
	// -- SWITCH --
	//การใช้เงื่อนไข switch เป็นอีกหนึ่งวิธีที่ทำได้ลักษณะคล้ายกันกับ if else if โดยใช้เพียงแค่ switch และระบุ statement จาก switch เพียงตัวเดียวในการจัดการ
	//
	//โดยปกติไอเดียการใช้ switch จะมีอยู่ 2 แบบคือ
	//
	//ดักจับจาก variable
	//ดักจับจาก condition (เหมือน if else if)
	//โดยวิธีการใช้งานคือ

	//switch variable {
	//case value1:
	//	// code ทำงานเมื่อ variable == value1
	//case value2:
	//	// code ทำงานเมื่อ variable == value2
	//default:
	//	// code ทำงานเมื่อ ไม่มี variable match กับเคสไหนเลย
	//}

	// -- Pre process if else --
	num1 := 5
	num2 := 10

	// วิธีปกติ
	//sumNum := num1 + num2

	//if sumNum >= 10 {
	//	fmt.Println("sumNum more than 10")
	//}

	// รูปแบบย่อ **ในกรณีนี้ sumNum จะอยู่ใน scope if เท่านั้น
	if sumNum := num1 + num2; sumNum >= 10 {
		fmt.Println("sumNum more than 10")
	}
}
