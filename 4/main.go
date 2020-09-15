package main

import "fmt"

type testStruct struct {
	foo string
	bar int
}

func (s testStruct) getInfo() string {
	return fmt.Sprintf("foo %s bar %d\n", s.foo, s.bar)
}

func main() {
	map1 := map[string]string{
		"name": "john",
		"age":  "12",
		"job":  "programmer",
	}

	type customer struct {
		name string
		sex  string
		age  int
		corp string
	}

	customer1 := customer{
		name: "Alex",
		sex:  "M",
		age:  21,
		corp: "omg ltd.",
	}

	ts1 := testStruct{
		foo: "asd",
		bar: 123,
	}

	fmt.Printf("%s", ts1.getInfo())
	fmt.Printf("1 customer name: %s sex: %s, age: %d, corp: %s\n", customer1.name, customer1.sex, customer1.age, customer1.corp)
	fmt.Printf("%+v\n", customer1)

	fmt.Println("vim-go")
	for key, value := range map1 {
		fmt.Printf("key: %s\tvalue: %s\n", key, value)
	}
}
