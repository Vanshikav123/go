package main

import (
	"fmt"
	"time"
)

//every process which run councurrently in golang called goroutines
//light weighted thread
//creation cost of goroutines is every small as compaired to thread
//every program had at least single goroutine called main function
//when main is terminated then all goroutines will terminated means all routines works under main

//syntax
/*func dosomething(){
logic}*/

//using go keyword as the prefix of your function call you can create goroutine

func main() {
	go doSomeMagic()
	//doSomeMagic()
	/*fmt.Println("hello world")
	go func() {
	fmt.Println("go function")	}()
	time.Sleep(1*timeSecond)*/
}

func doSomeMagic() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second) //this delays the loop which helps in goroutines to execute

		fmt.Println("goroutines")
	}
}
