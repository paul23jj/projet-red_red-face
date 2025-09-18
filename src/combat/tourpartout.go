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

func TourPartoutCombat(Personnage *class.Personnage, Monstre *Monstre.Monstre) {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	for Personnage.HP > 0 && Monstre.HP > 0 {
		fmt.Println("\n--- Tour de combat ---")
		fmt.Printf("%s : %d HP | %s : %d HP\n", Personnage.Nom, Personnage.HP, Monstre.Nom, Monstre.HP)
		fmt.Printf("Vitesse : %d | Vitesse : %d\n", Personnage.Vitesse, Monstre.Vitesse)

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
				UtiliserObjetParNumero(Personnage, Monstre)
			case "4":
				UtiliserPouvoir(Personnage, Monstre)
			case "5":
				Fuir(Personnage)
				return
			default:
				fmt.Println("Choix invalide.")
			}
			inv.PainDommage(Monstre)
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
					UtiliserObjetParNumero(Personnage, Monstre)
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

	// Fin du combat
	if Personnage.HP <= 0 {
		fmt.Println("GAV pour toi")
	} else {
		fmt.Printf("Tu as piétiner %s !\n", Monstre.Nom)
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
