package main

import "fmt"

//syntax
//func function_name(Parameter-list Type)(Return_type) {}
//types
//1. pass by value
//2. pass by reference

func main() {
	a := a()
	b := b()

	sum(a, b)
	passbyref(&a, &b)
	variadicFunc(a, b, 11)
}

//anonymous function
/*func(){
	fmt.Println("anonymous function")
}()

val:=func()int {
	return 10
}
fmt.Println(val())

func(v ...string) {
	fmt.Println(v)
}("test","best","abc")
*/
func a() int {
	return 10
}

func b() int {
	return 11
}

func sum(a, b int) {
	fmt.Println("pass by value:", a+b)
}

func passbyref(a, b *int) {
	fmt.Println("pass by reference:", *a+*b)
}

// variadic function
func variadicFunc(v ...int) {
	fmt.Println("elements are", v)
}

//init function  executed line by line they are called before main

func init() {
	fmt.Println("init function")

}
