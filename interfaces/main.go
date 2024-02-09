package main

import "fmt"

//collection of methods signatures
/*
Go language interfaces are different from other languages.
In Go language, the interface is a custom type that is used to specify a set of one or more method signatures and the interface is abstract,
so you are not allowed to create an instance of the interface. But you are allowed to create a variable of an interface type and this variable
can be assigned with a concrete type value that has the methods the interface requires.
Or in other words, the interface is a collection of methods as well as it is a custom type.*/

type values struct {
	first  int
	second int
}

type mathsTest interface {
	//all methods
	add(a, b int) int
	//multiply(a,b int) int

}

func (v values) add(a, b int) int {
	return a + b + v.first + v.second
}

//func (v values) multiply(a,b int)  int {

func main() {
	var te mathsTest = values{1, 2}
	fmt.Println(te.add(1, 2))
	var test interface{}
	test = "some string"

	val, ok := test.(string)
	if ok {
		fmt.Println(val)
	}
}
