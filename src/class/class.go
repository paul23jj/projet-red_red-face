package Class

import "fmt"

type Personnage struct {
	Nom                 string
	Classe              string
	Niveau              int
	HP                  int
	MaxHP               int
	Vitesse             int
	Force               int
	Intelligence        int
	Resistance          int
	Chance              int
	Saccoche            []Inventaire
	Pouvoirs            []string
	PouvoirCooldown     int
	Kishta              int
	SeringueTourRestant int
	PainTourRestant     int
}

type Inventaire struct {
	Name     string
	Quantity int
}

var P Personnage

func InitPlayer() Personnage {
	var choix int // <-- déclaration ici, visible partout

	fmt.Print("Ton blaze: ")
	fmt.Scan(&P.Nom)
	P.Kishta = 100 // Par exemple, 100 au départ
	P.Niveau = 1
	if P.Nom == "Kavtiv" {
		fmt.Println("=====Bienvenue Maître Kavtiv !=====")
		P.Classe = "Maître"
		P.HP = 1000
		P.MaxHP = 1000
		P.Force = 100
		P.Vitesse = 100
		P.Intelligence = 100
		P.Resistance = 100
		P.Chance = 100
		P.Kishta = 1000000
		P.Pouvoirs = []string{"Tacos 3 Viandes"}
		P.Saccoche = []Inventaire{{Name: "Sanglier", Quantity: 1}}
	} else {
		fmt.Println("Ton origine: ")
		fmt.Println("1. Nomade")
		fmt.Println("2. Russe")
		fmt.Println("3. Tchetchene")
		fmt.Println("4. Malien")
		fmt.Println("5. Bresilien")
		fmt.Print("Choix: ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			P.Classe = "Nomade"
			P.HP = 70
			P.MaxHP = 70
			P.Vitesse = 10
			P.Force = 5
			P.Intelligence = 3
			P.Resistance = 5
			P.Chance = 7
			P.Pouvoirs = []string{"lancer de cuivre"}
			P.Saccoche = []Inventaire{
				{Name: "herisson", Quantity: 1},
			}
		case 2:
			P.Classe = "Russe"
			P.HP = 100
			P.MaxHP = 100
			P.Vitesse = 3
			P.Force = 10
			P.Intelligence = 3
			P.Resistance = 7
			P.Chance = 3
			P.Pouvoirs = []string{"ak47"}
			P.Saccoche = []Inventaire{
				{Name: "vodka", Quantity: 1},
			}
		case 3:
			P.Classe = "Tchetchene"
			P.HP = 80
			P.MaxHP = 80
			P.Vitesse = 5
			P.Force = 7
			P.Intelligence = 5
			P.Resistance = 10
			P.Chance = 5
			P.Pouvoirs = []string{"corps à corps"}
			P.Saccoche = []Inventaire{
				{Name: "manuel de soumission", Quantity: 1},
			}
		case 4:
			P.Classe = "Malien"
			P.HP = 30
			P.MaxHP = 30
			P.Vitesse = 7
			P.Force = 3
			P.Intelligence = 10
			P.Resistance = 3
			P.Chance = 3
			P.Pouvoirs = []string{"magie noire"}
			P.Saccoche = []Inventaire{
				{Name: "bissap", Quantity: 1},
			}
		case 5:
			P.Classe = "Bresilien"
			P.HP = 50
			P.MaxHP = 50
			P.Vitesse = 5
			P.Force = 5
			P.Intelligence = 5
			P.Resistance = 5
			P.Chance = 10
			P.Pouvoirs = []string{"joga bonito"}
			P.Saccoche = []Inventaire{
				{Name: "shamballa", Quantity: 1},
			}
		}
	}
	fmt.Printf("Te voila enfin %s le %s\n", P.Nom, P.Classe)
	fmt.Println(P.Pouvoirs)
	fmt.Printf("veut tu regarder tes stats ?\n")
	fmt.Printf("1.Oui\n 2.Non\n")
	fmt.Scan(&choix)

	if choix == 1 {
		fmt.Printf("Hp: %d\n Force: %d\n Resistance: %d\n Intelligence: %d\n Vitesse: %d\n Chance: %d\n", P.HP, P.Force, P.Resistance, P.Intelligence, P.Vitesse, P.Chance)
		fmt.Println("Pouvoirs :", P.Pouvoirs)
	} else {
		fmt.Println("Pas grave tu peux toujours les voir dans le menu principal")
	}
	return P
}

type Monstre struct {
	Nom   string
	HP    int
	MaxHP int
}

func UtiliserPouvoir(P *Personnage, pouvoir string, cible *Monstre) {
	// Boost spécial pour Kavtiv
	var boost float64 = 1.0
	if P.Nom == "Kavtiv" {
		boost = 10.0 // Kavtiv est 10x plus fort pour la démo
	}

	switch pouvoir {
	case "lancer de cuivre":
		fmt.Println("Tu lances du cuivre !")
		cible.HP -= int(float64(P.Force) * 1.5 * boost)
		if cible.HP < 0 {
			cible.HP = 0
		}
	case "Flash":
		fmt.Println("Tu bois un flash !")
		cible.HP += int(float64(P.Force) * 2.0 * boost)
		if cible.HP < 0 {
			cible.HP = 0
		}
	case "corps à corps":
		fmt.Println("Attaque corps à corps !")
		cible.HP -= int(float64(P.Force) * 1.3 * boost)
		if cible.HP < 0 {
			cible.HP = 0
		}
	case "magie noire":
		fmt.Println("Tu utilises la magie noire !")
		P.HP += int(float64(P.Intelligence) * 2.0 * boost)
	case "joga bonito":
		fmt.Println("Tu esquives gracieusement !")
		P.Vitesse += int(float64(P.Vitesse) * 2.0 * boost)
	case "Ultime Kantin":
		fmt.Println("Kantin utilise son pouvoir ultime !")
		cible.HP -= int(float64(P.Force) * 20 * boost)
		if cible.HP < 0 {
			cible.HP = 0
		}
	default:
		fmt.Println("Pouvoir inconnu.")
	}
}
