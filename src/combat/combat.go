package combat

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	class "PROJETRED/src/class"
	Inventaire "PROJETRED/src/inventaire"
	Monstre "PROJETRED/src/monstre"
)

func Combat(Personnage *class.Personnage, Monstre *Monstre.Monstre) {
	fmt.Println(Personnage)
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	for Personnage.HP > 0 && Monstre.HP > 0 {
		fmt.Println("\n--- Tour de combat ---")
		fmt.Printf("%s : %d HP | %s : %d HP\n", Personnage.Nom, Personnage.HP, Monstre.Nom, Monstre.HP)
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
			Attaquer(Personnage, Monstre)
		case "2":
			Defendre(Personnage)
		case "3":
			UtiliserObjet(Personnage)
		case "4":
			UtiliserPouvoir(Personnage, Monstre)
		case "5":
			Fuir(Personnage)
			return
		default:
			fmt.Println("Choix invalide.")
		}

		// L'ennemi attaque si les deux sont encore vivants
		if Monstre.HP > 0 && Personnage.HP > 0 {
			EnnemiAttaque(Monstre, Personnage)
		}
		// Décrémenter le cooldown du pouvoir si besoin
		if Personnage.PouvoirCooldown > 0 {
			Personnage.PouvoirCooldown--
		}
	}
	if Personnage.HP <= 0 {
		fmt.Println("Tu as été vaincu.")
	} else {
		fmt.Println("Tu as vaincu l'ennemi.")
	}
}

func Fuir(Personnage *class.Personnage) {
	fmt.Printf("%s prend la fuite !\n", Personnage.Nom)
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
	Personnage.Resistance += 5
	fmt.Println("Défense augmentée pour ce tour !")
}

func Attaquer(Personnage *class.Personnage, Monstre *Monstre.Monstre) {
	damage := rand.Intn(10) + 1 // dégâts aléatoires entre 1 et 10
	if Personnage.Nom == "Kantin" {
		damage *= 10 // boost 10x pour la démo
	}
	Monstre.HP -= damage
	fmt.Printf("%s attaque %s et inflige %d dégâts!\n", Personnage.Nom, Monstre.Nom, damage)
	if Monstre.HP < 0 {
		Monstre.HP = 0
	}
}

func EnnemiAttaque(monstre *Monstre.Monstre, joueur *class.Personnage) {
	damage := rand.Intn(8) + 1 // dégâts aléatoires entre 1 et 8
	joueur.HP -= damage
	fmt.Printf("%s attaque %s et inflige %d dégâts!\n", monstre.Nom, joueur.Nom, damage)
	if joueur.HP < 0 {
		joueur.HP = 0
	}
}

func UtiliserPouvoir(Personnage *class.Personnage, cible *Monstre.Monstre) {
	fmt.Println(Personnage)
	if len(Personnage.Pouvoirs) == 0 {
		fmt.Println("Tu n'as pas de pouvoir spécial.")
		return
	}
	if Personnage.PouvoirCooldown > 0 {
		fmt.Printf("⏳ Ton pouvoir sera prêt dans %d tour(s).\n", Personnage.PouvoirCooldown)
		return
	}
	fmt.Println("Pouvoirs disponibles :")
	for i, pouvoir := range Personnage.Pouvoirs {
		fmt.Printf("%d) %s\n", i+1, pouvoir)
	}
	fmt.Print("Choisis un pouvoir : ")
	reader := bufio.NewReader(os.Stdin)
	choix, _ := reader.ReadString('\n')
	choix = strings.TrimSpace(choix)
	index, err := strconv.Atoi(choix)
	if err != nil || index < 1 || index > len(Personnage.Pouvoirs) {
		fmt.Println("Choix invalide.")
		return
	}

	// Appliquer l'effet du pouvoir choisi
	pouvoir := Personnage.Pouvoirs[index-1]
	switch pouvoir {
	case "lancer de cuivre":
		fmt.Println("💥 Tu lances un cuivre !")
		cible.HP -= int(float64(Personnage.Force) * 1.5)
	case "ak47":
		fmt.Println("🔫 Tu tires avec l'AK47 !")
		cible.HP -= int(float64(Personnage.Force) * 2.0)
	case "corps à corps":
		fmt.Println("🥊 Attaque corps à corps !")
		cible.HP -= int(float64(Personnage.Force) * 1.3)
	case "magie noire":
		fmt.Println("🧙 Tu utilises la magie noire !")
		cible.HP -= int(float64(Personnage.Intelligence) * 2.0)
	case "joga bonito":
		fmt.Println("⚽ Tu esquives gracieusement !")
		Personnage.Vitesse *= 2
	default:
		fmt.Println("Pouvoir inconnu.")
		return
	}
	Personnage.PouvoirCooldown = 3
	fmt.Println("⏳ Pouvoir utilisé ! Il sera à nouveau disponible dans 3 tours.")
}
