package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// input := scanner.Text()

	// fmt.Printf("You typed %s \n", input)
	// fmt.Printf("this is in quotes %q \n", input)

	fmt.Println("Type your birth year")

	// if type conversion doesn't work, error is stored inside _
	birth_year, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println(birth_year)
}
