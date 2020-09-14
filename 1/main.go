package main

import "fmt"

var GlobalVar string = "global var"

func main() {
	var i, j, k int
	var foo, bar, zar = true, 2.2, "zar"
	const Const = 1

	var LocalVar1 string = "localVar 1"
	var LocalVar2 = "localVar 2"
	LocalVar3 := "localVar 3"

	fmt.Println("Hello World")
	fmt.Println(i, j, k)
	fmt.Println(foo)
	fmt.Println(bar)
	fmt.Println(zar)
	fmt.Println(GlobalVar)
	fmt.Println(LocalVar1)
	fmt.Println(LocalVar2)
	fmt.Println(LocalVar3)
	fmt.Println(Const)
}
