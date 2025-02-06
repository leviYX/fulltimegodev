package main

import "fmt"

func main() {

	user := User{
		userName: "John",
		age:      getNumber(),
	}
	fmt.Println(user)
}
