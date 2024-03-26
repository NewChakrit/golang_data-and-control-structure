package array_and_slice

import "fmt"

func ArraySlice() {
	//Array และ Slice ใช้สำหรับเก็บข้อมูลที่เป็น sequence ของ Variable ใช้สำหรับการเก็บข้อมูลประเภทเดียวกันแต่เก็บทีละหลายตัวไปพร้อมๆกันได้
	// - โดยทั้ง Array และ Slice จะต้องประกาศ [] เพิ่มเพื่อเป็นการบอกว่าตัวแปรนี้จะเป้นประเภท Array
	// - โดย Array จะต้องระบุขนาดตอนประกาศว่าจะใช้ตัวแปรประเภทนี้เป็น Array ที่มีขนาด (จำนวนตัวแปรที่จะใช้งาน) จำนวนเท่าไหร่ เช่น ถ้าจะใช้งาน 3 ตัวก็ต้องประกาศเป็น var myArray[3] int เป็นต้น
	// - Slice จะกลับกัน slice ไม่จำเป็นต้องประกาศขนาดของ Array ไว้ก่อน สามารถใส่เพียง [] และค่อยมาจัดการเพิ่ม data ได้

	// -- Array --
	// - ลำดับแรกสุดของภาษา programming คือ 0
	// - ลำดับสุดท้ายจะเป็นขนาดของ array - 1
	var myArray [3]int // An array of 3 integers
	myArray[0] = 10    // Assign values
	myArray[1] = 20
	myArray[2] = 30

	myArray[0] = 100 // ressign
	//fmt.Println(myArray)    // Output: [100 20 30]
	//fmt.Println(myArray[1]) // 20

	for i := 0; i < len(myArray); i++ {
		fmt.Printf("Value of index %d : %v \n", i, myArray[i])
	}

	// -- Slice --
	// - มีค่าเริ่มต้นหรือไม่มีก็ได้
	// ยิ่งใช้เยอะ memory ก็ยิ่งยอะ

	mySlice := []int{10, 20, 30, 40, 50}

	fmt.Println(mySlice)      // Output: [10 20 30 40]
	fmt.Println(len(mySlice)) // Length of the slice: 5
	fmt.Println(cap(mySlice)) // Capacity of the slice : 5

	// Slicing a slice
	subSlice := mySlice[1:3]
	fmt.Println(subSlice)      //Output: [20 30]
	fmt.Println(len(subSlice)) // Length of the slice: 2
	fmt.Println(cap(subSlice)) //Capacity of the slice : 4

	// Append ต่อ slice
	mySlice = append(mySlice, 60, 70)
	fmt.Println(mySlice) // [10 20 30 40 50 60 70]

	// convert Array to Slice
	mySlice2 := myArray[:]
	mySlice2 = append(mySlice2, 40, 50)
	fmt.Println(mySlice2) // [100 20 30 40 50]

}
