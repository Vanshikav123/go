package main

import (
	"custom_packages/custom"
)

// functions or variables created with small letters can"t access outside the package
func main() {
	custom.PrintValue("vanshika")

}
