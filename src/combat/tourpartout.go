<<<<<<< HEAD
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
	Pouvoirs        []string
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
		fmt.Println("Bienvenue Maître Kantin !")
		// Boost spécial
		p.HP += 1000
		p.MaxHP += 1000
		p.Force += 100
		p.Vitesse += 100
		p.Intelligence += 100
		p.Resistance += 100
		p.Chance += 100
		// Pouvoir spécial
		p.Pouvoirs = append(p.Pouvoirs, "Ultime Kantin")
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
	fmt.Println(p.Pouvoirs)
	fmt.Printf("veut tu regarder tes stats ?\n")
	fmt.Printf("1.Oui\n 2.Non\n")
	fmt.Scan(&choix)

	if choix == 1 {
		fmt.Printf("Hp: %d\n Force: %d\n Resistance: %d\n Intelligence: %d\n Vitesse: %d\n Chance: %d\n", p.HP, p.Force, p.Resistance, p.Intelligence, p.Vitesse, p.Chance)
		fmt.Println("Pouvoirs :", p.Pouvoirs)
	} else {
		fmt.Println("Pas grave tu peux toujours les voir dans le menu principal")
	}
	return p
}

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
=======
package combat

import (
	class "PROJETRED/src/class"
	Monstre "PROJETRED/src/monstre"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func TourPartoutCombat(Personnage *class.Personnage, Monstre *Monstre.Monstre) {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	for Personnage.HP > 0 && Monstre.HP > 0 {
		fmt.Println("\n--- Tour de combat ---")
		fmt.Printf("%s : %d HP | %s : %d HP\n", Personnage.Nom, Personnage.HP, Monstre.Nom, Monstre.HP)
		fmt.Printf("Vitesse : %d | Vitesse : %d\n", Personnage.Vitesse, Monstre.Vitesse) // Affiche la vitesse

		// Détermine qui joue en premier
		joueurPremier := Personnage.Vitesse >= Monstre.Vitesse

		if joueurPremier {
			// Tour du joueur
			fmt.Println("À toi de jouer !")
			fmt.Println("1) Attaquer")
			fmt.Println("2) Défendre")
			fmt.Println("3) Utiliser un objet")
			fmt.Println("4) Utiliser un pouvoir")
			fmt.Println("5) Fuir")
			fmt.Print("Ton choix: ")
			choix, _ := reader.ReadString('\n')
			choix = strings.TrimSpace(choix)

			switch choix {
			case "1":
				Attaquer(Personnage, Monstre)
			case "2":
				Defendre(Personnage)
			case "3":
				UtiliserObjet(Personnage)
			case "4":
				UtiliserPouvoir(Personnage, Monstre)
			case "5":
				Fuir(Personnage)
				return
			default:
				fmt.Println("Choix invalide.")
			}

			// Si le monstre est encore vivant, il joue
			if Monstre.HP > 0 && Personnage.HP > 0 {
				EnnemiAttaque(Monstre, Personnage)
			}
		} else {
			// Tour du monstre
			fmt.Println("Le monstre agit en premier !")
			EnnemiAttaque(Monstre, Personnage)

			// Si le joueur est encore vivant, il joue
			if Personnage.HP > 0 && Monstre.HP > 0 {
				fmt.Println("À toi de jouer !")
				fmt.Println("1) Attaquer")
				fmt.Println("2) Défendre")
				fmt.Println("3) Utiliser un objet")
				fmt.Println("4) Utiliser un pouvoir")
				fmt.Println("5) Fuir")
				fmt.Print("Ton choix: ")
				choix, _ := reader.ReadString('\n')
				choix = strings.TrimSpace(choix)

				switch choix {
				case "1":
					Attaquer(Personnage, Monstre)
				case "2":
					Defendre(Personnage)
				case "3":
					UtiliserObjet(Personnage)
				case "4":
					UtiliserPouvoir(Personnage, Monstre)
				case "5":
					Fuir(Personnage)
					return
				default:
					fmt.Println("Choix invalide.")
				}
			}
		}

		// Décrémenter le cooldown du pouvoir si besoin
		if Personnage.PouvoirCooldown > 0 {
			Personnage.PouvoirCooldown--
		}
	}

	if Personnage.HP <= 0 {
		fmt.Println("Tu as été vaincu.")
	} else {
		fmt.Println("Tu as vaincu l'ennemi.")
>>>>>>> be7731b85df61a943c98fac909c47daabc977b75
	}
}
