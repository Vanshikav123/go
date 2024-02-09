package main

import "fmt"

//defined type which allows multiple real type identifiers
//collection of multiple data types

type user struct {
	name    string
	email   string
	phone   int
	address interface{}
}

func main() {
	var u user
	//fmt.Println(u)
	//1st way
	u = user{"vanshika", "vanshika@gmail.com", 12, "some address"}
	fmt.Println(u)
	//2nd way
	//u=user{name:"vanshika",email:"vanshikav@gmail.com",phone:12,address: "some address"}
	//u:=&user{name:"vanshika",email:"vanshikav@gmail.com"}
	//fmt.Println((*u).name)
	//fmt.Println(u.email)

}
