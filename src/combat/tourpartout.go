package combat

import (
	class "PROJETRED/src/class"
	inv "PROJETRED/src/inventaire"
	Monstre "PROJETRED/src/monstre"
	xp "PROJETRED/src/xp"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Attaquer(Personnage *class.Personnage, Monstre *Monstre.Monstre) {
	// Exemple d'implémentation simple de l'attaque
	degats := Personnage.Force - Monstre.Defense
	if degats < 0 {
		degats = 0
	}
	Monstre.HP -= degats
	fmt.Printf("%s attaque %s et inflige %d dégâts ! PV restant du monstre : %d\n", Personnage.Nom, Monstre.Nom, degats, Monstre.HP)
}

// Fuir permet au personnage de tenter de fuir le combat
func Fuir(Personnage *class.Personnage, Monstre *Monstre.Monstre) {
	chance := rand.Intn(100)
	if chance < 50 {
		fmt.Printf("%s réussit à fuir le combat contre %s !\n", Personnage.Nom, Monstre.Nom)
	} else {
		fmt.Printf("%s tente de fuir mais échoue ! Le combat continue...\n", Personnage.Nom)
	}
}

// UtiliserPouvoir permet au personnage d'utiliser son pouvoir spécial contre le monstre
func UtiliserPouvoir(Personnage *class.Personnage, Monstre *Monstre.Monstre) {
	if Personnage.PouvoirCooldown > 0 {
		fmt.Printf("%s ne peut pas utiliser son pouvoir spécial pour %d tour(s) !\n", Personnage.Nom, Personnage.PouvoirCooldown)
		return
	}
	// Exemple d'effet de pouvoir spécial : inflige des dégâts supplémentaires
	degats := Personnage.Force*2 - Monstre.Defense
	if degats < 0 {
		degats = 0
	}
	Monstre.HP -= degats
	fmt.Printf("%s utilise son pouvoir spécial et inflige %d dégâts à %s ! PV restant du monstre : %d\n", Personnage.Nom, degats, Monstre.Nom, Monstre.HP)
	Personnage.PouvoirCooldown = 3 // cooldown de 3 tours
}

// Defendre permet au personnage de réduire les dégâts subis au prochain tour
func Defendre(Personnage *class.Personnage) {
	Personnage.Resistance += 5
	fmt.Printf("%s se défend et augmente sa résistance de 5 pour ce tour !\n", Personnage.Nom)
}

// EnnemiAttaque permet au monstre d'attaquer le personnage
func EnnemiAttaque(Monstre *Monstre.Monstre, Personnage *class.Personnage) {
	degats := Monstre.Force - Personnage.Resistance
	if degats < 0 {
		degats = 0
	}
	Personnage.HP -= degats
	fmt.Printf("%s attaque %s et inflige %d dégâts ! PV restant du personnage : %d\n", Monstre.Nom, Personnage.Nom, degats, Personnage.HP)
}

func Combatmain(Personnage *class.Personnage, Monstre *Monstre.Monstre) {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	// Effets sur plusieurs tours
	painEffet := 0
	seringueEffet := 0

	for Personnage.HP > 0 && Monstre.HP > 0 {
		fmt.Println("\n--- Tour de combat ---")
		fmt.Printf("%s : %d HP | %s : %d HP\n", Personnage.Nom, Personnage.HP, Monstre.Nom, Monstre.HP)
		fmt.Printf("Vitesse : %d | Vitesse : %d\n", Personnage.Vitesse, Monstre.Vitesse)

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
				objetUtilise := inv.UtiliserObjetParNumero(Personnage, Monstre)
				if objetUtilise == "Pain" {
					painEffet = 3
				}
				if objetUtilise == "Seringue" {
					seringueEffet = 3
				}
			case "4":
				UtiliserPouvoir(Personnage, Monstre)
			case "5":
				Fuir(Personnage, Monstre)
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
					objetUtilise := inv.UtiliserObjetParNumero(Personnage, Monstre)
					if objetUtilise == "Pain" {
						painEffet = 3
					}
					if objetUtilise == "Seringue" {
						seringueEffet = 3
					}
				case "4":
					UtiliserPouvoir(Personnage, Monstre)
				case "5":
					Fuir(Personnage, Monstre)
					return
				default:
					fmt.Println("Choix invalide.")
				}
			}
		}

		// Effet du pain à chaque tour
		if painEffet > 0 {
			Monstre.HP -= 10
			painEffet--
			if Monstre.HP < 0 {
				Monstre.HP = 0
			}
			fmt.Printf("Le pain inflige 10 dégâts à %s ! PV restant : %d\n", Monstre.Nom, Monstre.HP)
		}

		// Effet de la seringue à chaque tour
		if seringueEffet > 0 {
			Personnage.HP += 5
			seringueEffet--
			if Personnage.HP > Personnage.MaxHP {
				Personnage.HP = Personnage.MaxHP
			}
			fmt.Printf("%s récupère 5 PV grâce à la seringue ! PV actuels : %d\n", Personnage.Nom, Personnage.HP)
		}

		// Décrémenter le cooldown du pouvoir si besoin
		if Personnage.PouvoirCooldown > 0 {
			Personnage.PouvoirCooldown--
		}
	}

	// Fin du combat
	if Personnage.HP <= 0 {
		fmt.Println("GAV pour toi")
	} else {
		fmt.Printf("Tu as piétiné %s !\n", Monstre.Nom)
		xp.GainXP(Personnage, Monstre.XPValue)

		// Vérifier le loot
		loot := Monstre.DropLoot()
		if loot != nil {
			Personnage.Saccoche = append(Personnage.Saccoche, *loot)
			fmt.Printf("💰 %s a drop un item : %s !\n", Monstre.Nom, loot.Name)
		} else {
			fmt.Printf("😔 Dommage, %s n’a rien drop cette fois ci...\n", Monstre.Nom)
		}
	}
}
