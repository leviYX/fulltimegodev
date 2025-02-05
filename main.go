package main

import "fmt"

var name = "John Doe"
var age int

const (
	version = 1
	key     = 10
)

type Person struct {
	name string
	age  int8

	score struct {
		math    int8
		english int8
	}
}

func main() {
	// structDemo()
	// mapDemo1()
	//mapDemo3()
	// fmt.Printf("damage:%d\n", getDamage("ord"))

}

func getDamage(weaponType string) int {
	switch weaponType {
	case "sword":
		return 10
	case "gun":
		return 20
	default:
		panic("unknow weapon type")
	}
}

func mapDemo1() {
	var cars map[string]string = map[string]string{
		"benz": "mercedes",
		"bmw":  "bmw",
		"audi": "audi",
	}
	fmt.Println(cars)
	// 删除map中的某个key
	delete(cars, "benz")
}

func mapDemo2() {
	var users map[string]int = map[string]int{
		"levi":  18,
		"hanji": 20,
		"ailun": 14,
	}

	leviAge, hasFlag := users["levi"]
	if hasFlag {
		fmt.Printf("levi age is %d \n", leviAge)
	} else {
		fmt.Println("levi age is not exist")
	}
}

func mapDemo3() {
	var users map[string]int = make(map[string]int)
	users["levi"] = 18
	users["hanji"] = 20
	users["ailun"] = 14
	for k, v := range users {
		fmt.Printf("key:%s,value:%d\n", k, v)
	}

	fmt.Println("----------")
	for k := range users {
		fmt.Printf("key:%s \n", k)
	}

	fmt.Println("----------")

	for _, v := range users {
		fmt.Printf("value:%d \n", v)
	}
}

func structDemo() {
	var p1 = Person{
		name: "John Doe",
		age:  18,
		score: struct {
			math    int8
			english int8
		}{
			math:    90,
			english: 80,
		},
	}

	fmt.Printf("%+v\n", p1)
}
