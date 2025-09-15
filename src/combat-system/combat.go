package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	class "PROJETRED/src/class"
	monster "PROJETRED/src/monster"
)

// Remove local Character type if class.Character is used

func main() {
	// Initialisation du joueur via le package class
	player := class.InitCharacter() // doit retourner un class.Character

	// Création d’un ennemi niveau 2 via le package monster
	enemy := monster.CreateMonster(2) // doit retourner un class.Character

	// Lancer un combat
	StartCombat(player, enemy)
}
func StartCombat(player class.Character, enemy class.Character) {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	for player.HP > 0 && enemy.HP > 0 {
		fmt.Println("Choisis une action :")
		fmt.Println("1) Attaquer")
		fmt.Println("2) Défendre")
		fmt.Println("3) Utiliser un objet")
		fmt.Println("4) Fuir")
		fmt.Print("Ton choix:")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			damage := max(0, player.Attack-enemy.Defense+rand.Intn(3))
			enemy.HP -= damage
			fmt.Printf("\n💥 Tu attaques %s et infliges %d dégâts !\n", enemy.Name, damage)
		case "2":
			fmt.Println("\n🛡️ Tu te prépares à encaisser le coup !")
			player.Defense += 2
		case "3":
			fmt.Println("\n💊 Tu utilises un bissap (+10 HP) !")
			player.HP += 10
		case "4":
			fmt.Println("\n🏃 Tu fuis le combat...")
			return
		default:
			fmt.Println("Commande invalide !")
			continue
		}

		if enemy.HP <= 0 {
			fmt.Printf("\n🎉 Tu as péta %s !🎉\n", enemy.Name)
			break
		}

		// Tour de l'ennemi
		enemyDamage := max(0, enemy.Attack-player.Defense+rand.Intn(3))
		player.HP -= enemyDamage
		fmt.Printf("⚔️ %s attaque et inflige %d dégâts !\n", enemy.Name, enemyDamage)

		if player.Defense > 2 {
			player.Defense -= 2 // reset bonus défense
		}
	}

	if player.HP <= 0 {
		fmt.Printf("\n💀 %s t'a vaincu...\n", enemy.Name)
		fmt.Println("Tu vas en garde à vue...")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
