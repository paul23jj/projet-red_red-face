package menuDemarrage

import (
	Classe "PROJETRED/src/class"
	"PROJETRED/src/four"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartMenu() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Bienvenue dans Projet-Red ===")
	fmt.Print("Veux-tu rentrer dans la tess ? (oui/non) : ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(strings.ToLower(choice))

	if choice == "oui" {
		player := Classe.InitPlayer() // initialisation du joueur

		scanner := bufio.NewScanner(os.Stdin)

		for {
			fmt.Println("\n=== Menu Principal ===")
			fmt.Println("1) Aller au Marché du Soleil 🏪")
			fmt.Println("2) Aller voir Le Four 🔥")
			fmt.Println("3) Quitter la tess")
			fmt.Print("Choix : ")

			scanner.Scan()
			choix := strings.TrimSpace(scanner.Text())

			switch choix {
			case "1":
				Classe.EntrerMarche(player) // marche déjà existant
			case "2":
				four.EntrerForge(player, Classe.ShowStats) // forge
			case "3":
				fmt.Println("👉 Tu es retourné à la tess. À bientôt !")
				return
			default:
				fmt.Println("Choix invalide.")
			}
		}

	} else {
		fmt.Println("Dommage... à bientôt !")
	}
}
