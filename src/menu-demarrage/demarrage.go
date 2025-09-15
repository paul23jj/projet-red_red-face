package menuDemarrage

import (
	Classe "PROJETRED/src/class"
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
		Classe.InitPlayer()

		for {
			fmt.Println("\n=== Menu Principal ===")
			fmt.Println("1. Aller dans le Four")
			fmt.Println("2. Aller au Marché")
			fmt.Println("3. Chercher un tête à tête")
			fmt.Println("4. Quitter")
			fmt.Print("Choisis une option : ")

			menuChoice, _ := reader.ReadString('\n')
			menuChoice = strings.TrimSpace(menuChoice)

			switch menuChoice {
			case "1":
				fmt.Println("Tu es maintenant dans le Four !")
				// Ici tu peux appeler ton module Four : four.StartFour()
			case "2":
				fmt.Println("Tu es maintenant au Marché !")
				// Ici tu peux appeler ton module Marché : marche.StartMarche()
			case "3":
				fmt.Println("Tu cherches un tête à tête...")
			case "4":
				fmt.Println("À bientôt !")
				return
			default:
				fmt.Println("Option invalide, réessaie.")
			}
		}

	} else {
		fmt.Println("Dommage... à bientôt !")
	}
}
