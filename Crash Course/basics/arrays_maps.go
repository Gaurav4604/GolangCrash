package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 100
	fmt.Println(arr)

	// explicitly defining elements of array
	arr_2 := []int{1, 2, 3, 4, 5}
	fmt.Println(arr_2)
	fmt.Println(len(arr_2))

	for i, elem := range arr_2 {
		fmt.Println(i, elem)
	}

	// maps
	var mp map[string]int = map[string]int{
		"apple":  1,
		"orange": 2,
		"banana": 3,
	}

	fmt.Println(mp)

	mp_2 := make(map[string]int)

	mp_2["inserted_value"] = 10

	delete(mp_2, "inserted_value")

	_, ok := mp_2["inserted_value"]

	if ok {
		fmt.Println(mp_2["inserted_value"])
	} else {
		fmt.Println("value not there")
	}
}
