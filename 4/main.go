package main

import "fmt"

type testStruct struct {
	foo string
	bar int
}

func newTestStruct(foo string, bar int) testStruct {
	return testStruct{
		foo: foo,
		bar: bar,
	}
}

func (t testStruct) getInfo() string {
	return fmt.Sprintf("foo %s bar %d\n", t.foo, t.bar)
}

func setBar(t *testStruct, bar int) {
	t.bar = bar
}

type age int

func (a age) isAdult() bool {
	return a >= 18
}

func main() {

	myAge := age(17)
	fmt.Println("is adult ? ", myAge.isAdult())

	map1 := map[string]string{
		"name": "Max",
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

	ts2 := newTestStruct("qwerty", 12345)

	fmt.Printf("%s", ts1.getInfo())
	fmt.Printf("%s", ts2.getInfo())
	setBar(&ts2, 2222222222222)
	fmt.Printf("%s", ts2.getInfo())
	fmt.Printf("1 customer name: %s sex: %s, age: %d, corp: %s\n", customer1.name, customer1.sex, customer1.age, customer1.corp)
	fmt.Printf("%+v\n", customer1)

	fmt.Println("vim-go")
	for key, value := range map1 {
		fmt.Printf("key: %s\tvalue: %s\n", key, value)
	}
}
