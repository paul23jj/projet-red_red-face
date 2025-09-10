package class

import "fmt"

type personnage struct {
	nom          string
	classe       string
	niveau       int
	hp           int
	max_hp       int
	vitesse      int
	force        int
	intelligence int
	resistance   int
	chance       int
	inventaire   []inventaire
}

type inventaire struct {
	name     string
	quantity int
}

func InitPlayer() {
	var p personnage

	fmt.Print("Ton blaze: ")
	fmt.Scan(&p.nom)

	fmt.Println("Ton origine: ")
	fmt.Println("1. Gitan")
	fmt.Println("2. Russe")
	fmt.Println("3. Tchetchene")
	fmt.Println("4. Malien")
	fmt.Println("5. Bresilien")
	fmt.Print("Choix: ")

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		p.classe = "Gitan"
		p.hp = 70
		p.max_hp = 70
		p.vitesse = 10
		p.force = 5
		p.intelligence = 3
		p.resistance = 5
		p.chance = 7
		p.inventaire = []inventaire{
			{name: "herisson", quantity: 1},
		}
	case 2:
		p.classe = "Russe"
		p.hp = 100
		p.max_hp = 100
		p.vitesse = 3
		p.force = 10
		p.intelligence = 3
		p.resistance = 7
		p.chance = 3
		p.inventaire = []inventaire{
			{name: "vodka", quantity: 1},
		}
	case 3:
		p.classe = "tchetchene"
		p.hp = 80
		p.max_hp = 80
		p.vitesse = 5
		p.force = 7
		p.intelligence = 5
		p.resistance = 10
		p.chance = 5
		p.inventaire = []inventaire{
			{name: "manuel de soumission", quantity: 1},
		}

	case 4:
		p.classe = "Malien"
		p.hp = 30
		p.max_hp = 30
		p.vitesse = 7
		p.force = 3
		p.intelligence = 10
		p.resistance = 3
		p.chance = 3
		p.inventaire = []inventaire{
			{name: "bissap", quantity: 1},
		}
	case 5:
		p.classe = "Bresilien"
		p.hp = 50
		p.max_hp = 50
		p.vitesse = 5
		p.force = 5
		p.intelligence = 5
		p.resistance = 5
		p.chance = 10
		p.inventaire = []inventaire{
			{name: "shamballa", quantity: 1},
		}
	}
	fmt.Printf("Te voila enfin %s le %s\n", p.nom, p.classe)
}
