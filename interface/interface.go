package _interface

import (
	"fmt"
	"math"
)

// Speaker interface
type Speaker interface {
	Speak() string
	Walk() string
}

// Dog struct
type Dog struct {
	Name string
}

// Person struct
type Person struct {
	Name string
}

// Dog's implemention of the Speaker intereface
func (d Dog) Speak() string {
	return "Woof"
}

func (p Person) Speak() string {
	return "Hello!"
}

func (d Dog) Walk() string {
	return "Walk!"
}

func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func Interface() {
	dog := Dog{Name: "Jack"}
	//person := Person{Name: "Alice"}

	makeSound(dog)
	//makeSound(person)

}

// ------------------------------------------ //
// Ex 2

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	length, width float64
}

// Area method for Ractangle
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Triangle struct
type Triangle struct {
	base, height float64
}

// Area method for Ractangle
func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

// Circle struct
type Circle struct {
	radius float64
}

// Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func Calculater(s Shape) {
	fmt.Println(s.Area())
}

func Interface2() {
	rectangle := Rectangle{length: 2, width: 2.5}
	triangle := Triangle{base: 4, height: 3.5}
	circle := Circle{radius: 4}

	Calculater(rectangle)
	Calculater(triangle)
	Calculater(circle)
}
