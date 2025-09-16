package menuDemarrage

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	Classe "PROJETRED/src/class"
	Combat "PROJETRED/src/combat"
	Inventaire "PROJETRED/src/inventaire"
	Monstre "PROJETRED/src/monstre"
)

var player Classe.Personnage
var monstre Monstre.Monstre

func StartMenu() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Bienvenue dans Projet-Red ===")
	fmt.Print("Veux-tu rentrer dans la tess ? (oui/non) : ")
	choice, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erreur de lecture, à bientôt !")
		return
	}
	choice = strings.TrimSpace(strings.ToLower(choice))

	if choice != "oui" {
		fmt.Println("Dommage... à bientôt !")
		return
	}

	player = Classe.InitPlayer()

	for {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Println("1. Aller dans le Four")
		fmt.Println("2. Aller au Marché")
		fmt.Println("3. Regarder la sacoche")
		fmt.Println("4. Chercher un tête à tête")
		fmt.Println("5. Quitter")
		fmt.Print("Choisis une option : ")

		menuChoice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erreur de lecture, réessaie.")
			continue
		}
		menuChoice = strings.TrimSpace(menuChoice)

		switch menuChoice {
		case "1":
			fmt.Println("Tu es maintenant dans le Four !")
			gererFour()
		case "2":
			fmt.Println("Tu es maintenant au Marché !")
			gererMarche()
		case "3":
			fmt.Println("Voici ta sacoche :")
			Inventaire.AfficherSacoche(&player)
		case "4":
			fmt.Println("Tu cherches un tête à tête...")
			monstre = Monstre.GenererMonstre()
			Combat.CombatMain(&player, &monstre)
		case "5":
			fmt.Println("À bientôt !")
			os.Exit(0)
		default:
			fmt.Println("Option invalide, réessaie.")
		}

		// Vider le buffer en lisant les caractères restants jusqu'à la fin de la ligne
		reader.ReadString('\n')
	}
}

func gererFour() {
	fmt.Println("Bienvenue dans le Four !")
	// Ajoute ici la logique pour le Four
}

func gererMarche() {
	fmt.Println("Bienvenue au Marché !")
	// Ajoute ici la logique pour le Marché
}
