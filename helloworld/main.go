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
	var names [3]string
	names[0]="Alice"
	names[2]="Charly"
	names2 := [4]string {"toto", "tata"}
	fmt.Printf("name = %v (len = %v) \n", names, len(names))
	fmt.Printf("name = %v (len = %v) \n", names2, len(names2))

	fmt.Println(strings.ToUpper("Hello World, welcome to ags training"))
	fmt.Println(data.Name, data.Password)
	if data.Age > 10 {
		fmt.Println("sup a 10")
	} else if data.Age < 10 {
		fmt.Println("inf a 10")
	} else {
		fmt.Println("age = 10")
	}
	utils.Hello("Tonio")
	fmt.Println(sumNamedReturn(1,2,3))
}
