package main

import "fmt"

type Animal interface {
	Speak() string
	Eat(string)
}

type Dog struct {
	name string
	food string
}

func (dog Dog) Speak() string {
	return dog.name + " says Woof!"
}

func (dog Dog) Eat(food string) {
	fmt.Printf("%s eat %s!\n", dog.name, food)
}

func main() {
	dog := Dog{
		name: "Rex",
	}
	var animal Animal = dog
	animal.Eat("Woof")
	fmt.Println(animal.Speak())
}
