package main

import (
	"fmt"
)

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
	u.display()
}

func (u User) display() {
	fmt.Printf("%s, %d/%d hp, %s, level %d\n", u.name, u.hp, u.max_hp, u.classe, u.level)
}

func (u *User) make0ld() {
	u.level += 3
}

func main() {
	user := User{"enzo", 4, 'F', "Demon", []inventaire{{"piment", 3}, {"essence", 1}}, 100, 100}
	user.display()
	user.make0ld()
	user.display()
}
