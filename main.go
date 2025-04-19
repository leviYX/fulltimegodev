package main

import (
	"fmt"
	"mygoproject/utils"
)

func main() {
	var num = utils.GetRandNum(10)
	fmt.Println(num)
}
