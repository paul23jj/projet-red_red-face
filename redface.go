package main

import "fmt"

type inventaire struct {
	name     string
	quantity int
}

type User struct {
	name   string
	level  int
	rang   rune
	classe string
	inv    []inventaire
	hp     int
	max_hp int
}

func Perso() {
	u := User{
		name:   "Redface",
		level:  4,
		rang:   'F',
		classe: "Demon",
		inv: []inventaire{
			{name: "piment", quantity: 3},
			{name: "essence", quantity: 1},
		},
		max_hp: 100,
		hp:     100,
	}
	fmt.Println(u)
}
