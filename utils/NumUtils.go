package utils

import (
	"math/rand"
)

func GetRandNum(seed int) int {
	return rand.Intn(seed)
}
