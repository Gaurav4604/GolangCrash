package main

import "fmt"

func main() {
	// explicitly defined type
	var name string = "Gaurav Singh"

	fmt.Println(name[0:])

	// implicitly defined type
	x := 10
	x += 5
	fmt.Println(x)

	// printing type of variable
	fmt.Printf("%T\n", x)

	// storing the print value onto a variable
	stored_print := fmt.Sprintf("Hi! my name is %s", name)

	fmt.Println(stored_print)
}
