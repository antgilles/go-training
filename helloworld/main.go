package main

import (
	"ags.training.go/helloworld/utils"
	"ags.training.go/helloworld/data"
	"fmt"
	"strings"
)
func sumNamedReturn(x, y, z int) (sum int, name string){
	sum = x + y + z
	name = "toto"
	return
}
func main() {

	// Slices
	nums := []int{1,2,3,4}
	nums[0] = 5
	nums[1] = -6
	fmt.Printf("%v (len = %v, cap = %v) \n", nums, len(nums), cap(nums))
	nums =append(nums, 7)
	fmt.Printf("%v (len = %v, cap = %v) \n", nums, len(nums), cap(nums))
	nums =append(nums, 5, 6)
	fmt.Printf("%v (len = %v, cap = %v) \n", nums, len(nums), cap(nums))

	// Arrays
	var names [3]string
	names[0]="Alice"
	names[2]="Charly"
	names2 := [4]string {"toto", "tata"}
	fmt.Printf("name = %v (len = %v) \n", names, len(names))
	fmt.Printf("name = %v (len = %v) \n", names2, len(names2))

	// Print
	fmt.Println(strings.ToUpper("Hello World, welcome to ags training"))
	fmt.Println(data.Name, data.Password)
	
	// If then else
	if data.Age > 10 {
		fmt.Println("sup a 10")
	} else if data.Age < 10 {
		fmt.Println("inf a 10")
	} else {
		fmt.Println("age = 10")
	}

	// Package and public/private vars/func
	utils.Hello("Tonio")
	fmt.Println(sumNamedReturn(1,2,3))

	// loop for
	for i := 0; i < 5; i++ {
		println("loop :", i)
	}
}
