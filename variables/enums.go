package main

import "fmt"

type WeaponType int

const (
	Sword WeaponType = iota
	Knife
	Axe
)

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Sword:
		return 10
	case Knife:
		return 5
	default:
		panic("weaponType is not valid")
	}
}

func getDamage2(weaponType string) int {
	switch weaponType {
	case "sword":
		return 10
	case "bow":
		return 5
	default:
		panic("weaponType is not valid")
	}
}

func main() {
	fmt.Println("weapon damage is ", getDamage(Sword))
}
