package main

import "fmt"

/*a closure is the function value that references variable from outside its body.The function may access and assign to the referenced variables;in this sense the function is "bound" to the variables */

/*anonymous function
func main()
{  func(s string){
	fmt.Println(s)
}("hello")
	}*/

func company() func() int {
	a := 0
	return func() int {
		a++
		return a
	}
}

func main() {
	val := company()

	fmt.Println(val())
	fmt.Println(val())
	fmt.Println(val())
	fmt.Println(val())

	v := company()
	fmt.Println(v())

}
