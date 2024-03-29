package function

import "fmt"

func SayHello(name string, age int) {
	fmt.Println("Hi, I'm ", name)
	fmt.Printf("I'm %v yesrs old.\n", age)
}

func Add(num1, num2 int) int {
	return num1 + num2
}

type Student struct {
	FirstName string
	LastName  string
}

func PrintName() {
	student := Student{
		FirstName: "New",
		LastName:  "Chakrit",
	}
	fullName := student.FullName()
	fmt.Println("Full Name of student: ", fullName)
}

// -- Method --
// Method with reciever or type Student

func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}
