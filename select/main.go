//syntax
/*select {
case SendOrReceivecase: //statement
.......
default: //statement
}*/

package main

/*In Go language, the select statement is just like switch statement,
 but in the select statement, case statement refers to communication, i.e. sent or receive operation on the channel.
 Syntax:


select{

case SendOrReceive1: // Statement
case SendOrReceive2: // Statement
case SendOrReceive3: // Statement
.......
default: // Statement
Important points:


Select statement waits until the communication(send or receive operation) is prepared for some cases to begin. */

import (
	"fmt"
	"time"
)

// function 1
func portal1(channel1 chan string) {

	time.Sleep(3 * time.Second)
	channel1 <- "Welcome to channel 1"
}

// function 2
func portal2(channel2 chan string) {

	time.Sleep(9 * time.Second)
	channel2 <- "Welcome to channel 2"
}

// main function
func main() {

	// Creating channels
	R1 := make(chan string)
	R2 := make(chan string)

	// calling function 1 and
	// function 2 in goroutine
	go portal1(R1)
	go portal2(R2)

	select {

	// case 1 for portal 1
	case op1 := <-R1:
		fmt.Println(op1)

	// case 2 for portal 2
	case op2 := <-R2:
		fmt.Println(op2)
	}

}

/*
Explanation: In the above program, portal 1 sleep for 3 seconds and portal 2 sleep for 9 seconds after their
sleep time over they will ready to proceed. Now, select statement waits till their sleep time, when the portal 1 wakes up, it selects case 1 and prints “Welcome to channel 1”.
If the portal 2 wakes up before portal 1 then the output is “welcome to channel 2”. */
