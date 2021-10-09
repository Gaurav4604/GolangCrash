package main

import "fmt"

// pass by reference
func changeValue(str *string) {
	*str = "changed" // dereference memory loc value and change its value
}

func main() {
	x := 10
	fmt.Println(x)

	y := &x // y points to x's memory loc

	x -= 1

	fmt.Println(y, x)

	// dereference value pointed by y
	// and assign new value to it
	*y = 9

	fmt.Println(x)

	arr := []func(*string){changeValue}

	str := "string"
	fmt.Println(str)
	// pointer to str's memory loc
	arr[0](&str)
	fmt.Println(str)
}
