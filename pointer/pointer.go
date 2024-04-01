package pointer

import "fmt"

type Employee struct {
	Name   string
	Salary int
}

func Pointer() {
	// Phase by reference จริงๆเป็นการ copy value ไปอีก address หนึ่ง แล้วทำการ verify address นั้น ผ่านตัว reference ที่มีการส่งเข้าไปแทน
	//x := 20
	//changeValue(&x)
	//fmt.Println(x) // 50

	// Declare an integer cariable
	x := 10

	// Declare a pointer to an integet and assign it
	var p *int = &x

	// Print the value of x and the value at the poiner
	fmt.Println("Value of x:", x) //
	fmt.Println("Value of p:", *p)

	// Modify the value at the pointer p
	*p = 20

	//x is modified since p points to x
	fmt.Println("New value of x: ", x)

	// -- Pointer with struct --
	emp := Employee{Name: "John", Salary: 50000}

	giveRaise(&emp, 5000)
	fmt.Println("After raise:", emp) //After raise: {John 55000}
}

func changeValue(ptr *int) {
	*ptr = 50
}

func giveRaise(e *Employee, raise int) {
	e.Salary += raise
}
