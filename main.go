package main

import (
	"fmt"
	"github.com/NewChakrit/golang_data-and-control-structure/function"
)

func main() {

	//variable.Variable()
	//control_structure.ControlStructure()
	//iteration.Iteration()
	//array_and_slice.ArraySlice()
	//_map.Map()
	//_struct.Struct()
	function.SayHello("New", 28)
	function.SayHello("Wen", 28)

	result := function.Add(5, 4)
	fmt.Println("Result function Add = ", result)

	function.PrintName()
}
