package combat

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	class "PROJETRED/src/class"
	Monstre "PROJETRED/src/monstre"
)

func Combat(player *class.Personnage, enemy *Monstre.Monstre) {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	for player.HP > 0 && enemy.HP > 0 {
		fmt.Println("\n--- Tour de combat ---")
		fmt.Printf("%s : %d HP | %s : %d HP\n", player.Nom, player.HP, enemy.Nom, enemy.HP)
		fmt.Println("Actions disponibles :")
		fmt.Println("1) Attaquer")
		fmt.Println("2) Défendre")
		fmt.Println("3) Utiliser un objet")
		fmt.Println("4) Fuir")
		fmt.Print("Ton choix: ")
		choix, _ := reader.ReadString('\n')
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":
			Attaquer(player, enemy)
		case "2":
			Defendre(player)
		case "3":
			UtiliserObjet(player)
		case "4":
			Fuir(player)
		default:
			fmt.Println("Choix invalide.")
		}
	}

	if player.HP <= 0 {
		fmt.Println("Tu as été vaincu.")
	} else {
		fmt.Println("Tu as vaincu l'ennemi.")
	}
}

func Fuir(player *class.Personnage) {
	panic("unimplemented")
}

func UtiliserObjet(player *class.Personnage) {
	panic("unimplemented")
}

// func Defendre removed to fix redeclaration error
func Defendre(player *class.Personnage) {
	fmt.Printf("%s se met en position de défense.\n", player.Nom)
	// Exemple simple : augmenter temporairement la défense
	player.Resistance += 5
	fmt.Println("Défense augmentée pour ce tour !")
}

// Ajout de la fonction Attaquer
func Attaquer(player *class.Personnage, enemy *Monstre.Monstre) {
	// Exemple simple d'attaque
	damage := rand.Intn(10) + 1 // dégâts aléatoires entre 1 et 10
	enemy.HP -= damage
	fmt.Printf("%s attaque %s et inflige %d dégâts!\n", player.Nom, enemy.Nom, damage)
	if enemy.HP < 0 {
		enemy.HP = 0
	}
}
