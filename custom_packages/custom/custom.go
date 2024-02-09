package custom

import "fmt"

func PrintValue(s string) {
	printVal(s)
}

func printVal(data string) {
	fmt.Println("in reserve function")
	fmt.Println(data)
}
