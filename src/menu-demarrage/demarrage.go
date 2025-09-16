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
)

var player class.Personnage
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

	player = class.InitPlayer()

	if player.Saccoche == nil {
		player.Saccoche = []class.Inventaire{}
	}

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

		case "3":
			fmt.Println("Voici ta sacoche :")
			inventaire.AfficherSacoche(&player)
		case "4":
			fmt.Println("Tu cherches un tête à tête...")
			currentMonstre = monstre.GenererMonstre()
			combat.TourPartoutCombat(&player, &currentMonstre)
		case "5":
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
		Classe:       player.Classe,
		Hp:           player.HP,
		MaxHp:        player.MaxHP,
		Vitesse:      player.Vitesse,
		Force:        player.Force,
		Intelligence: player.Intelligence,
		Resistance:   player.Resistance,
		Chance:       player.Chance,
		Kishta:       player.Kishta,
		Inventaire:   []four.Inventaire{},
	}

	// Appeler la fonction du Four
	four.EntrerForge(pFour, showStats)

	// Synchroniser les changements vers player
	player.HP = pFour.Hp
	player.MaxHP = pFour.MaxHp
	player.Vitesse = pFour.Vitesse
	player.Force = pFour.Force
	player.Intelligence = pFour.Intelligence
	player.Resistance = pFour.Resistance
	player.Chance = pFour.Chance
	player.Kishta = pFour.Kishta
	// Synchroniser l'inventaire
	player.Saccoche = []class.Inventaire{}
}
