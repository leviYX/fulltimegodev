package main

import "fmt"

var name string = "golang,你好"

type Player struct {
	name        string
	age         int
	health      int
	attackPower float64
}

type Weapon string

func main1() {
	var player = Player{
		name:        "jack",
		age:         40,
		health:      100,
		attackPower: 125.01,
	}
	fmt.Printf("the player is %+v\n", player)
	fmt.Printf("player health is %d\n", getHealth(player))
	fmt.Printf("player name is %s\n", player.getName())
}

func getWeapon(weapon Weapon) string {
	return string(weapon)
}

func getHealth(player Player) int {
	return player.age
}

func (player Player) getName() string {
	return player.name
}

func loopStr() {
	var nameRune = []rune(name)
	for _, r := range nameRune {
		fmt.Printf("%c", r)
	}
	fmt.Println()
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c", name[i])
	}
}
