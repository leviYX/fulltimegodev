package utils

import (
	"fmt"
	"math/rand"
)

func GetRandNum(seed int) int {
	fmt.Println("seed:", seed)
	return rand.Intn(seed)
}
