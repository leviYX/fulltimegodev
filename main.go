package main

import (
	"fmt"
	"myprojectmod/utils"
)

func main() {

	user := utils.User{
		UserName: "John",
		Age:      utils.GetRandNum(100),
	}
	fmt.Printf("user数据为%+v", user)
}
