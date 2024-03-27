package _struct

import "fmt"

// -- Struct --
type Student struct {
	Name   string
	Height int
	Weight int
	Grande string
}

func Struct() {
	// Create an instance of the Student struct
	var student1 Student
	student1.Name = "New"
	student1.Height = 171
	student1.Weight = 68
	student1.Grande = "A"

	fmt.Println(student1) //{New 171 68 A}

	// Struct with Array
	var student [3]Student
	student[0].Name = "Wen"
	student[0].Height = 150
	student[0].Weight = 50
	student[0].Grande = "AA"

	student[1].Name = "Enw"
	student[1].Height = 160
	student[1].Weight = 60
	student[1].Grande = "AA"

	student[2].Name = "Wne"
	student[2].Height = 170
	student[2].Weight = 70
	student[2].Grande = "AA"

	fmt.Println(student) //[{Wen 150 50 AA} {Enw 160 60 AA} {Wne 170 70 AA}]

	var countStuden int
	for i := 0; i < len(student); i++ {
		countStuden = i + 1
		fmt.Println("Student", (i + 1))
		fmt.Printf("Name of studen %d : %v\n", countStuden, student[i].Name)
		fmt.Printf("Height of studen %d : %v\n", countStuden, student[i].Height)
		fmt.Printf("Weight of studen %d : %v\n", countStuden, student[i].Weight)
		fmt.Printf("Grande of Grande %d : %v\n", countStuden, student[i].Grande)
	}

	students := make(map[string]Student)
	students["st01"] = Student{Name: "MikeLopster", Height: 180, Weight: 72, Grande: "B+"}
	students["st02"] = Student{Name: "Alice", Height: 152, Weight: 38, Grande: "B-"}
	students["st03"] = Student{Name: "Bob", Height: 164, Weight: 58, Grande: "c"}

	fmt.Println("Students Map", students)

	// Access and print a specific student by key
	fmt.Println("Student st01:", students["st01"])

	var person Person
	person.Name = "Phone"
	person.Age = 25
	person.Address = Address{
		Street:  "56/7",
		City:    "Bangkok",
		Zipcode: 10111,
	}

	fmt.Println(person)
}

// Struct สามารถมี struct อยู่ข้างในได้

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Street  string
	City    string
	Zipcode int
}
