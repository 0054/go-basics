package main

import "fmt"

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
}
