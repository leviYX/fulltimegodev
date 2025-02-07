package main

import "fmt"

type Color int

const (
	Red Color = iota
	Blue
)

type DataInterface interface {
	String() string
}

/*
*
实现了String接口，所以传入Color类型就会调用该方法。
go的接口实现，只要你满足了接口的所有方法，就实现了该接口。并且你的func(类型)这就表示这个类型实现了接口。
*/
func (c Color) String() string {
	switch c {
	case Red:
		return "red"
	case Blue:
		return "blue"
	default:
		return "unknown color"
	}
}

func main2() {
	/**
	Println 方法会调用参数的String 方法，然后我的func (c Color) String() string 实现了该方法，所以会打印出对应的颜色。
	等于是String的接口实现，这里Color实现了String接口。所以传入Color类型，就会调用该方法。
	*/
	fmt.Println(Red)
}
