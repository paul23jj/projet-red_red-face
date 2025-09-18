package menuDemarrage

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	class "PROJETRED/src/class"
	combat "PROJETRED/src/combat"
	four "PROJETRED/src/forge"
	inventaire "PROJETRED/src/inventaire"
	marche "PROJETRED/src/marche"
	monstre "PROJETRED/src/monstre"
)

var Player class.Personnage
var currentMonstre monstre.Monstre

func StartMenu() {
	// Scanner unique pour tout le menu
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Bienvenue dans Projet-Red ===")
	fmt.Print("Veux-tu rentrer dans la tess ? (oui/non) : ")
	scanner.Scan()
	choice := strings.TrimSpace(strings.ToLower(scanner.Text()))

	if choice != "oui" {
		fmt.Println("Dommage... à bientôt !")
		return
	}

	Player = class.InitPlayer()

	if Player.Saccoche == nil {
		Player.Saccoche = []class.Inventaire{}
	}

	for {
		fmt.Println("\n=== La Tess ===")
		fmt.Println("1. Aller dans le Four")
		fmt.Println("2. Aller au Marché")
		fmt.Println("3. Regarder la sacoche")
		fmt.Println("4. Chercher un tête à tête")
		fmt.Println("5. Regarder les stats")
		fmt.Println("6. Quitter")
		fmt.Print("Choisis une option : ")

		scanner.Scan()
		menuChoice := strings.TrimSpace(scanner.Text())

		switch menuChoice {
		case "1":
			fmt.Println("Tu es maintenant dans le Four !")
			gererFour(&Player)
		case "2":
			fmt.Println("Tu es maintenant au Marché !")
			gererMarche(scanner)
		case "3":
			fmt.Println("Voici ta sacoche :")
			inventaire.AfficherSacoche(&Player)
		case "4":
			fmt.Println("Tu cherches un tête à tête...")
			currentMonstre = monstre.GenererMonstre()
			combat.TourPartoutCombat(&Player, &currentMonstre)
		case "5":
			fmt.Println("Voici tes stats :")
			marche.ShowStats(&Player, scanner) // <-- scanner partagé
		case "6":
			fmt.Println("À bientôt !")
			os.Exit(0)
		default:
			fmt.Println("Option invalide, réessaie.")
		}
	}
}

func gererFour(p *class.Personnage) {
	// Convertir class.Personnage en four.Personnage

	fmt.Println(p.Saccoche)

	// Appeler la fonction du Four
	four.EntrerForge(p)

	fmt.Println(p.Saccoche)
	fmt.Println(Player.Saccoche)

}

func gererMarche(scanner *bufio.Scanner) {
	// On passe le scanner à MarcheDuSoleil pour utiliser un scanner unique
	marche.MarcheDuSoleil(&Player, scanner)
}
