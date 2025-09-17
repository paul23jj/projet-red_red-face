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
	monstre "PROJETRED/src/monstre"
	marche "PROJETRED/src/marche"
)

var Player class.Personnage
var currentMonstre monstre.Monstre

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

	Player = class.InitPlayer()

	if Player.Saccoche == nil {
		Player.Saccoche = []class.Inventaire{}
	}

	for {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Println("1. Aller dans le Four")
		fmt.Println("2. Aller au Marché")
		fmt.Println("3. Regarder la sacoche")
		fmt.Println("4. Chercher un tête à tête")
		fmt.Println("5. Regarder les stats")
		fmt.Println("6. Quitter")
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

		case "3":
			fmt.Println("Voici ta sacoche :")
			inventaire.AfficherSacoche(&Player)
		case "4":
			fmt.Println("Tu cherches un tête à tête...")
			currentMonstre = monstre.GenererMonstre()
			combat.TourPartoutCombat(&Player, &currentMonstre)
		case "5":
			fmt.Println("Voici tes stats :")
			marche.ShowStats(&Player)
		case "6":
			fmt.Println("À bientôt !")
			os.Exit(0)
		default:
			fmt.Println("Option invalide, réessaie.")
		}

		// Vider le buffer
		reader.ReadString('\n')
	}
}

func gererFour() {
	// Fonction pour afficher les stats compatible avec four.Personnage
	showStats := func(p *four.Personnage) {
		fmt.Printf("\n--- Stats de ton perso (%s) ---\n", p.Classe)
		fmt.Printf("HP: %d/%d | Force: %d | Vitesse: %d | Intel: %d | Résistance: %d | Chance: %d | Kishta: %d\n",
			p.Hp, p.MaxHp, p.Force, p.Vitesse, p.Intelligence, p.Resistance, p.Chance, p.Kishta)
		fmt.Println("Inventaire :")
		if len(p.Inventaire) == 0 {
			fmt.Println(" (vide)")
		} else {
			for _, it := range p.Inventaire {
				fmt.Printf(" - %s x%d\n", it.Name, it.Quantity)
			}
		}
	}

	// Convertir class.Personnage en four.Personnage
	pFour := &four.Personnage{
		Classe:       Player.Classe,
		Hp:           Player.HP,
		MaxHp:        Player.MaxHP,
		Vitesse:      Player.Vitesse,
		Force:        Player.Force,
		Intelligence: Player.Intelligence,
		Chance:       Player.Chance,
		Kishta:       Player.Kishta,
		Inventaire:   []four.Inventaire{},
	}

	// Appeler la fonction du Four
	four.EntrerForge(pFour, showStats)

	// Synchroniser les changements vers Player
	Player.HP = pFour.Hp
	Player.MaxHP = pFour.MaxHp
	Player.Vitesse = pFour.Vitesse
	Player.Force = pFour.Force
	Player.Intelligence = pFour.Intelligence
	Player.Resistance = pFour.Resistance
	Player.Chance = pFour.Chance
	Player.Kishta = pFour.Kishta
	// Synchroniser l'inventaire
	Player.Saccoche = []class.Inventaire{}
}
