package main

import "fmt"

// Creating structure
type Author struct {
	name   string
	branch string
	year   int
}

// Creating nested structure
type HR struct {

	// structure as a field
	details Author
}

func main() {

	// Initializing the fields
	// of the structure
	result := HR{

		details: Author{"Sona", "ECE", 2013},
	}

	// Display the values
	fmt.Println("\nDetails of Author")
	fmt.Println(result)
}
