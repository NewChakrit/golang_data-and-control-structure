package iteration

import "fmt"

func Iteration() {
	// -- Loop --
	//loop อยู่ทั้งหมด 3 ประเภทใหญ่ๆ
	//	1.For Loop
	//	2.Do While Loop
	//	3.While Loop
	//

	// -- For Loop
	//for i := 1; i <= 10; i++ {
	//	fmt.Println("number: ", i)
	//}

	// Do while
	// ประกาศสิ่งที่เป็น for และสิ้นสุดเมื่อใช้คำสั่ง break
	i := 1
	for {
		fmt.Println("number I: ", i)
		i++
		if i == 10 {
			break
		}
	}

	// While Loop
	j := 1
	for j < 10 {
		fmt.Println("number J: ", j)
		j++
	}
}
