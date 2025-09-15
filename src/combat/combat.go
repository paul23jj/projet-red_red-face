package combat

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	class "PROJETRED/src/class"
	Inventaire "PROJETRED/src/inventaire"
	Monstre "PROJETRED/src/monstre"
)

func Combat(Personnage *class.Personnage, Monstre *Monstre.Monstre) {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	for Personnage.HP > 0 && Monstre.HP > 0 {
		fmt.Println("\n--- Tour de combat ---")
		fmt.Printf("%s : %d HP | %s : %d HP\n", Personnage.Nom, Personnage.HP, Monstre.Nom, Monstre.HP)
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
			Attaquer(Personnage, Monstre)
		case "2":
			Defendre(Personnage)
		case "3":
			UtiliserObjet(Personnage)
		case "4":
			Fuir(Personnage)
		default:
			fmt.Println("Choix invalide.")
		}
	}

	if Personnage.HP <= 0 {
		fmt.Println("Tu as été vaincu.")
	} else {
		fmt.Println("Tu as vaincu l'ennemi.")
	}
}

func Fuir(Personnage *class.Personnage) {
	panic("unimplemented")
}

func UtiliserObjet(p *class.Personnage) {
	Inventaire.AfficherSacoche()
	fmt.Print("Quel objet veux-tu utiliser ? ")
	reader := bufio.NewReader(os.Stdin)
	objet, _ := reader.ReadString('\n')
	objet = strings.TrimSpace(objet)
	Inventaire.UtiliserObjet(objet, p)
}

func Defendre(Personnage *class.Personnage) {
	fmt.Printf("%s se met en position de défense.\n", Personnage.Nom)
	// Exemple simple : augmenter temporairement la défense
	Personnage.Resistance += 5
	fmt.Println("Défense augmentée pour ce tour !")
}

// Ajout de la fonction Attaquer
func Attaquer(Personnage *class.Personnage, Monstre *Monstre.Monstre) {
	// Exemple simple d'attaque
	damage := rand.Intn(10) + 1 // dégâts aléatoires entre 1 et 10
	Monstre.HP -= damage
	fmt.Printf("%s attaque %s et inflige %d dégâts!\n", Personnage.Nom, Monstre.Nom, damage)
	if Monstre.HP < 0 {
		Monstre.HP = 0
	}
}
