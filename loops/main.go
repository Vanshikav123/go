package main

import "fmt"

func main() {
	//for loops in go

	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}

	//while loop mimic in golang
	j := 0
	for j < 4 {
		fmt.Println(j)
		j++
	}

	//range loop works on slice or array

	rvariable := []string{"a", "b", "c", "d", "e", "f", "g"}
	for i, j := range rvariable { //i,j returns index and value
		fmt.Println(i, j)
	}

	//similarly the above loop can be used on strings there only

	//loop on channels
	chnl := make(chan int)
	go func() {
		chnl <- 10
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		close(chnl)
	}()
	for i := range chnl {
		fmt.Println(i)
	}
	//we can run for loop on objects also

	mmap := map[int]string{
		0: "a",
		1: "b",
		2: "c",
	}

	for key, value := range mmap {
		fmt.Println(key, value)
	}
}
