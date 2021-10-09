package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// basic math between diff data types
	var num1 int = 10
	var num2 float64 = 2

	// explicit type conversion is required
	ans := float64(num1) / num2

	fmt.Println(ans)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if input > 10 {
		fmt.Println("greater than 10")
	} else if input == 10 {
		fmt.Println("equal to 10")
	} else {
		fmt.Println("less than 10")
	}
}
