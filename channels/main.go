/*In Go language, a channel is a medium through which a goroutine communicates with another goroutine and this communication is lock-free.
Or in other words, a channel is a technique which allows to let one goroutine to send data to another goroutine. By default channel is bidirectional,
 means the goroutines can send or receive data through the same channel as shown in the below image:


In Go language, a channel is created using chan keyword and it can only transfer data of the same type,
 different types of data are not allowed to transport from the same channel.
*/

/*var Channel_name chan Type
You can also create a channel using make() function using a shorthand declaration.

Syntax:

channel_name:= make(chan Type)*/

package main

import "fmt"

func main() {
	chanl1 := make(chan string)
	chanl2 := make(chan string)

	go sending(chanl1)

	valueFromChannel := <-chanl1

	fmt.Println("valueFromChannel:", valueFromChannel)

	go recieving(chanl2)

	chanl2 <- valueFromChannel
	//craeting a Biderectional channel
	chanl3 := make(chan string)
	go convert(chanl1)
	fmt.Println(<-chanl3)
}
func convert(s chan<- string) { //send only
	s <- "some value"
}
func recieving(s chan string) {
	fmt.Println(<-s)
}

func sending(s chan string) {
	s <- "vanshika"
}
