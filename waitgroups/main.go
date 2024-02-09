// how can we control go routines using wait groups
// methods which controlls go routines
// synchronized channel
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// func worker(id int,wg *sync.WaitGroup)
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
	//wg.Done()
}

func main() {
	//chan1:=make(chan bool)
	//also wg.Add(5) works before the for loop
	for i := 1; i <= 5; i++ {
		wg.Add(1) //number of goroutines
		i := i
		go func() {
			defer wg.Done()
			worker(i)
			//worker(i,&wg)

		}()
	}
	wg.Wait()
}
