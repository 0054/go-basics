package main

import (
	"errors"
	"fmt"
)

type employee struct {
	id     int
	name   string
	age    int
	salary int
}

type storage interface {
	insert(e employee) error
	get(id int) (employee, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]employee
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]employee),
	}
}

func (s *memoryStorage) insert(e employee) error {
	s.data[e.id] = e
	return nil
}

func (s *memoryStorage) get(id int) (employee, error) {
	e, exists := s.data[id]
	if !exists {
		return employee{}, errors.New("employee with such id  doesn't exist")
	}
	return e, nil
}

func (s *memoryStorage) delete(id int) error {
	delete(s.data, id)
	return nil
}

func printType(value interface{}) {
	if _, ok := value.(int); ok {
		fmt.Println("тип аргумента int")
	} else if _, ok := value.(string); ok {
		fmt.Println("тип аргумента string")
	} else {
		fmt.Println("тип аргумента отличный от int и string")
	}
}
func printType2(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("тип аргумента int")
	case string:
		fmt.Println("тип аргумента string")
	default:
		fmt.Println("тип аргумента отличный от int и string")
	}
}

func main() {
	var s storage

	printType(3)
	printType("adsada")
	printType([]string{"foo", "bar"})

	printType2(3)
	printType2("adsada")
	printType2([]string{"foo", "bar"})

	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)

	s = newMemoryStorage()

	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)

}
