package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	x := 3
	for x < 6 {
		fmt.Println("hello there")
		x++
	}

	for x := 0; x < 2; x++ {
		println("hello")
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ans, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	switch ans {
	case 1:
		println("one")
	case 2:
		println("two")
	default:
		println("not one or two")
	}
}
