package combat

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"PROJETRED/src/class"
	"PROJETRED/src/inventaire"
	"PROJETRED/src/monstre"
)

func Combat(p *class.Personnage, m *monstre.Monstre) {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	for p.HP > 0 && m.HP > 0 {
		fmt.Println("\n--- Tour de combat ---")
		fmt.Printf("%s : %d/%d HP | %s : %d/%d HP\n", p.Nom, p.HP, p.MaxHP, m.Nom, m.HP, m.MaxHP)
		fmt.Println("Actions disponibles :")
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
			damage := rand.Intn(p.Force) + 1
			m.HP -= damage
			fmt.Printf("%s attaque %s et inflige %d dégâts!\n", p.Nom, m.Nom, damage)
		case "2":
			p.Resistance += 5
			fmt.Printf("%s se met en défense (+5 Résistance ce tour)!\n", p.Nom)
		case "3":
			inventaire.AfficherSacoche()
			fmt.Print("Quel objet veux-tu utiliser ? ")
			objet, _ := reader.ReadString('\n')
			objet = strings.TrimSpace(objet)
			inventaire.UtiliserObjet(objet, p)
		case "4":
			if len(p.Pouvoirs) > 0 {
				class.UtiliserPouvoir(p, p.Pouvoirs[0], m)
			} else {
				fmt.Println("Pas de pouvoir disponible.")
			}
		case "5":
			if rand.Intn(100) < 50 {
				fmt.Println("Fuite réussie !")
				return
			} else {
				fmt.Println("Fuite échouée ! L'ennemi contre-attaque.")
			}
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		if m.HP <= 0 {
			fmt.Printf("Tu as vaincu %s !\n", m.Nom)
			xp.GainXP(p, 5) // Gagne 5 XP par combat
			p.Kishta += 20  // Récompense en Kishta
			return
		}

		m.EnnemiAttaque(p)
		if p.HP <= 0 {
			fmt.Println("Tu as été vaincu...")
			return
		}
	}
}
