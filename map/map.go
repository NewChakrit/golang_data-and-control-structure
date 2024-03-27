package _map

import "fmt"

func Map() {
	// -- Map --
	//map เป็น unordered collection ที่แต่ละ key "ต้อง unique" และจะต้อง map กับ value เสมอ (เมื่อมีการประกาศ key)
	//key ต้อง unique. 1 key สามารถมีได้เพียง 1 value เท่าั้น
	//key ใน map จะ "ไม่ guaranteed" ว่าถูกเรียงอย่างถูกต้องหรือไม่ (ถ้าซีเรียสว่าข้อมูลต้องเรียงกัน ไม่แนะนำให้ใช้ map หรือควร sort ก่อน)
	//Map เป็น reference types เมื่อมีการส่งเข้า function จะเป็นการ pass by reference เช่นเดียวกันกับ Slice

	// make(map[<ประเภท key>]<ประเภท value>)
	myMap := make(map[string]int)
	// Add key-value pairs to the map
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8

	// Access and print a value for a key
	fmt.Println("Apples:", myMap["apple"]) // Apples: 5

	// Update the value for a key
	myMap["banana"] = 12
	fmt.Println("Banana:", myMap["banana"]) // Banana: 12

	// Delete a key-value pair
	delete(myMap, "orange")
	fmt.Println("Orange:", myMap["orange"]) // Orange: 0

	// Iterate over the map
	for key, value := range myMap {
		fmt.Printf("%s -> %d \n", key, value)
	}

	// Checking if a key exists
	myMap["pear"] = 14
	val, ok := myMap["pear"]
	if ok {
		fmt.Println("Pear's value:", val)
	} else {
		fmt.Println("Pear not found inmap")
	}
}
