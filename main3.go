package main

import "fmt"

type Position struct {
	x, y int
}

func (p Position) Move(val int) {
	fmt.Println("位置移动距离 :", val)
}

type Player struct {
	Position
}

func main3() {
	p := Player{}
	p.Move(1)
}
