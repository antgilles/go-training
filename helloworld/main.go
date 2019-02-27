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
