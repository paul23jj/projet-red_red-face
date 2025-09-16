package Class

import "fmt"

type Personnage struct {
	Nom             string
	Classe          string
	Niveau          int
	HP              int
	MaxHP           int
	Vitesse         int
	Force           int
	Intelligence    int
	Resistance      int
	Chance          int
	Inventaire      []Inventaire
	Pouvoirs        []string // Ajout du pouvoir unique
	PouvoirCooldown int
}

type Inventaire struct {
	Name     string
	Quantity int
}

func InitPlayer() Personnage {
	var p Personnage

	fmt.Print("Ton blaze: ")
	fmt.Scan(&p.Nom)
	if p.Nom == "Kantin" {
		fmt.Println("=====Bienvenue Maître Kantin !=====")
		p.Classe = "Dieu vivant"
		p.HP = 1000
		p.MaxHP = 1000
		p.Force = 100
		p.Vitesse = 100
		p.Intelligence = 100
		p.Resistance = 100
		p.Chance = 100
		p.Pouvoirs = []string{"Ultime Kantin"}

		fmt.Printf("Te voila enfin %s le %s\n", p.Nom, p.Classe)

		var choix int
		fmt.Printf("Veux-tu regarder tes stats ?\n1.Oui\n2.Non\n")
		fmt.Scan(&choix)

		if choix == 1 {
			fmt.Printf("Hp: %d\nForce: %d\nResistance: %d\nIntelligence: %d\nVitesse: %d\nChance: %d\n",
				p.HP, p.Force, p.Resistance, p.Intelligence, p.Vitesse, p.Chance)
		} else {
			fmt.Println("Pas grave, tu peux toujours les voir dans le menu principal")
		}

		return p // ensuite on retourne Kantin
	}

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
		p.Classe = "Nomade"
		p.HP = 70
		p.MaxHP = 70
		p.Vitesse = 10
		p.Force = 5
		p.Intelligence = 3
		p.Resistance = 5
		p.Chance = 7
		p.Pouvoirs = []string{"lancer de cuivre"}
		p.Inventaire = []Inventaire{
			{Name: "herisson", Quantity: 1},
		}
	case 2:
		p.Classe = "Russe"
		p.HP = 100
		p.MaxHP = 100
		p.Vitesse = 3
		p.Force = 10
		p.Intelligence = 3
		p.Resistance = 7
		p.Chance = 3
		p.Pouvoirs = []string{"ak47"}
		p.Inventaire = []Inventaire{
			{Name: "vodka", Quantity: 1},
		}
	case 3:
		p.Classe = "Tchetchene"
		p.HP = 80
		p.MaxHP = 80
		p.Vitesse = 5
		p.Force = 7
		p.Intelligence = 5
		p.Resistance = 10
		p.Chance = 5
		p.Pouvoirs = []string{"corps à corps"}
		p.Inventaire = []Inventaire{
			{Name: "manuel de soumission", Quantity: 1},
		}
	case 4:
		p.Classe = "Malien"
		p.HP = 30
		p.MaxHP = 30
		p.Vitesse = 7
		p.Force = 3
		p.Intelligence = 10
		p.Resistance = 3
		p.Chance = 3
		p.Pouvoirs = []string{"magie noire"}
		p.Inventaire = []Inventaire{
			{Name: "bissap", Quantity: 1},
		}
	case 5:
		p.Classe = "Bresilien"
		p.HP = 50
		p.MaxHP = 50
		p.Vitesse = 5
		p.Force = 5
		p.Intelligence = 5
		p.Resistance = 5
		p.Chance = 10
		p.Pouvoirs = []string{"joga bonito"}
		p.Inventaire = []Inventaire{
			{Name: "shamballa", Quantity: 1},
		}
	}
	fmt.Printf("Te voila enfin %s le %s\n", p.Nom, p.Classe)
	fmt.Printf("veut tu regarder tes stats ?\n")
	fmt.Printf("1.Oui\n 2.Non\n")
	fmt.Scan(&choix)

	if choix == 1 {
		fmt.Printf("Hp: %d\n Force: %d\n Resistance: %d\n Intelligence: %d\n Vitesse: %d\n Chance: %d\n", p.HP, p.Force, p.Resistance, p.Intelligence, p.Vitesse, p.Chance)
	} else {
		fmt.Println("Pas grave tu peux toujours les voir dans le menu principal")
	}
	return p
}

// Definition du monstre
type Monstre struct {
	Nom   string
	HP    int
	MaxHP int
}

func UtiliserPouvoir(p *Personnage, pouvoir string, cible *Monstre) {
	// Boost spécial pour Kantin
	var boost float64 = 1.0
	if p.Nom == "Kantin" {
		boost = 10.0 // Kantin est 10x plus fort pour la démo
	}

	switch pouvoir {
	case "lancer de cuivre":
		fmt.Println("Tu lances du cuivre !")
		cible.HP -= int(float64(p.Force) * 1.5 * boost)
		if cible.HP < 0 {
			cible.HP = 0
		}
	case "Flash":
		fmt.Println("Tu bois un flash !")
		cible.HP += int(float64(p.Force) * 2.0 * boost)
		if cible.HP < 0 {
			cible.HP = 0
		}
	case "corps à corps":
		fmt.Println("Attaque corps à corps !")
		cible.HP -= int(float64(p.Force) * 1.3 * boost)
		if cible.HP < 0 {
			cible.HP = 0
		}
	case "magie noire":
		fmt.Println("Tu utilises la magie noire !")
		p.HP += int(float64(p.Intelligence) * 2.0 * boost)
	case "joga bonito":
		fmt.Println("Tu esquives gracieusement !")
		p.Vitesse += int(float64(p.Vitesse) * 2.0 * boost)
	case "Ultime Kantin":
		fmt.Println("Kantin utilise son pouvoir ultime !")
		cible.HP -= int(float64(p.Force) * 20 * boost)
		if cible.HP < 0 {
			cible.HP = 0
		}
	default:
		fmt.Println("Pouvoir inconnu.")
	}
}
