package menuDemarrage

import (
	Classe "PROJETRED/src/class"
	Combat "PROJETRED/src/combat"
	Four "PROJETRED/src/forge"
	Marche "PROJETRED/src/marche"
	Monstre "PROJETRED/src/monstre"
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
				Four.StartFour()
			case "2":
				fmt.Println("Tu es maintenant au Marché !")
				Marche.StartMarche()
			case "3":
				fmt.Println("Tu cherches un tête à tête...")
				player := Classe.Personnage{Nom: "Joueur", HP: 100, Resistance: 10}
				ennemi := Monstre.Monstre{Nom: "Ennemi", HP: 80}
				Combat.Combat(&player, &ennemi)
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
