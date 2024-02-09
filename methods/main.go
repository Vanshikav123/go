package main

import "fmt"

type data int

type user struct {
	name  string
	email string
}

func (v1 data) div(v2 data) data {
	return v1 / v2
}

func main() {
	//methods in golang'
	//func (receiver_name_type) method_name(parameter_list_type)(return_type))
    //method name can be same but the receiver function will be different
	value1 := data(23)
	value2 := data(10)

	res := value1.div(value2)
	fmt.Println("final result:", res)

    res1:=user{
		name:"vanshika",
		email:"vanshika@gmail.com",
	}

	fmt.Println("user name:",res1.name)
	fmt.Println("user email:",res1.email)

	p:=&res1
	p.correctEmail("vanshika1@gmail.com")
	
	fmt.Println("user name:",res1.name)
	fmt.Println("user email:",res1.email)
    
	func(a *user) correctEmail(email string) {
        (*a).email=email
	}



}
