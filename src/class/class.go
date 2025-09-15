package Classe

import "fmt"

type Personnage struct {
	Nom          string
	Classe       string
	Niveau       int
	Hp           int
	Max_hp       int
	Vitesse      int
	Force        int
	Intelligence int
	Resistance   int
	Chance       int
	Inventaire   []Inventaire
}

type Inventaire struct {
	Name     string
	Quantity int
}

func InitPlayer() {
	var p Personnage

	fmt.Print("Ton blaze: ")
	fmt.Scan(&p.Nom)

	fmt.Println("Ton origine: ")
	fmt.Println("1. Nomade")
	fmt.Println("2. Russe")
	fmt.Println("3. Tchetchene")
	fmt.Println("4. Malien")
	fmt.Println("5. Bresilien")
	fmt.Print("Choix: ")

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		p.Classe = "Gitan"
		p.Hp = 70
		p.Max_hp = 70
		p.Vitesse = 10
		p.Force = 5
		p.Intelligence = 3
		p.Resistance = 5
		p.Chance = 7
		p.Inventaire = []Inventaire{
			{Name: "herisson", Quantity: 1},
		}
	case 2:
		p.Classe = "Russe"
		p.Hp = 100
		p.Max_hp = 100
		p.Vitesse = 3
		p.Force = 10
		p.Intelligence = 3
		p.Resistance = 7
		p.Chance = 3
		p.Inventaire = []Inventaire{
			{Name: "vodka", Quantity: 1},
		}
	case 3:
		p.Classe = "tchetchene"
		p.Hp = 80
		p.Max_hp = 80
		p.Vitesse = 5
		p.Force = 7
		p.Intelligence = 5
		p.Resistance = 10
		p.Chance = 5
		p.Inventaire = []Inventaire{
			{Name: "manuel de soumission", Quantity: 1},
		}

	case 4:
		p.Classe = "Malien"
		p.Hp = 30
		p.Max_hp = 30
		p.Vitesse = 7
		p.Force = 3
		p.Intelligence = 10
		p.Resistance = 3
		p.Chance = 3
		p.Inventaire = []Inventaire{
			{Name: "bissap", Quantity: 1},
		}
	case 5:
		p.Classe = "Bresilien"
		p.Hp = 50
		p.Max_hp = 50
		p.Vitesse = 5
		p.Force = 5
		p.Intelligence = 5
		p.Resistance = 5
		p.Chance = 10
		p.Inventaire = []Inventaire{
			{Name: "shamballa", Quantity: 1},
		}
	}
	fmt.Printf("Te voila enfin %s le %s\n", p.Nom, p.Classe)
}
