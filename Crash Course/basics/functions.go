package main

import "fmt"

func test() {
	fmt.Println("Test")
}

func test_params(x, y int) int {
	fmt.Println("Test", x, y)
	return x + y
}

func labeled_return(name, greeting string) (greet_string string) {
	greet_string = fmt.Sprintf("%s, %s", greeting, name)
	return // returns greet_string as the return value without specifying it explicitly in the return
}

func main() {
	test()
	ans := test_params(10, 12)
	fmt.Println(ans)

	greeting := labeled_return("Gaurav", "hello there")
	fmt.Println(greeting)

	// assigning a function to variable
	x := test_params
	x(10, 12)

	// creating function inside function (anonymous function)
	anonymous := func(x int) {
		fmt.Println(x)
	}
	anonymous(1000)

	// iffi
	iffi := func(name string) string {
		return "Hello there " + name
	}("Gaurav")

	fmt.Printf("iffi value %q\n", iffi)

}
