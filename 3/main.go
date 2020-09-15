package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var list1 = [3]string{
		"1qaz",
		"2wsx",
		"3edc",
	}

	for index, item := range list1 {
		fmt.Printf("index: %d  item: %s\n", index, item)
	}

	var list2 = [3]int{}
	list2[0] = 1
	list2[1] = 2
	list2[2] = 3

	for _, item := range list2 {
		fmt.Printf("item: %d\n", item)
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	var slice1 = []int{}
	slice1 = append(slice1, 1, 3, 4, 5)
	fmt.Println(slice1)

	a(20)
}

func a(n int) error {
	if n <= 0 {
		return errors.New("n must be positive")
	}
	for i := 1; i <= n; i++ {
		fmt.Println(math.Pow(2, float64(i)))
	}
	return nil
}

// func b(arr []int) []int {
// 	l := len(arr)
// 	res := []int{}
// 	for index, item := range arr {
// 		append(res, item)
// 	}
// 	return res
// }
