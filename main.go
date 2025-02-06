package main

import (
	"fmt"
	"myprojectmod/utils"
)

type Entity struct {
	version    string
	name       string
	id         int
	posx, posy int
}

type specialEntity struct {
	entity       Entity
	specialField string
}

type demoEntity struct {
	Entity
	specialField string
}

func main1() {
	e1 := Entity{
		version: "1.0.0",
		name:    "John",
		id:      1,
		posx:    100,
		posy:    200,
	}

	e2 := demoEntity{
		Entity:       e1,
		specialField: "special",
	}

	e3 := specialEntity{
		entity: Entity{
			version: "1.0.0",
			name:    "John",
			id:      1,
			posx:    100,
			posy:    200,
		},
		specialField: "special",
	}

	fmt.Printf("e1 data is %+v\ne2 data is %+v\ne3 data is %+v\n", e1, e2, e3)
	fmt.Println(e3.entity.name, e2.name)
}

func main0() {

	user := utils.User{
		UserName: "John",
		Age:      utils.GetRandNum(100),
	}
	fmt.Printf("user数据为%+v", user)
}
