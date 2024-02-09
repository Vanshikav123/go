package main

import "fmt"

//collection of unordered pair of keys and values
//widely used to do operations like lookup,update,delete
//it is a refrence type
//inexpensive to pass
//keys must be unique
//values may not be unique
//it is also known as hash map, dictionary
//by default maps are nil

// map[Key_Type]Value_Type{}

//map wih key value pair
//map[Key_Type]Value_Type{key1:value1,...,keyN:valueN}

func main() {
	var map1 map[string]int

	fmt.Println(map1 == nil)

}
