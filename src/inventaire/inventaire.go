package inventaire

import (
	"fmt"
)

var Sacoche = map[string]int{}

func AfficherSacoche() {
	fmt.Println("\n📦 Sacoche :")
	if len(Sacoche) == 0 {
		fmt.Println("(vide)")
		return
	}
	for objet, qte := range Sacoche {
		fmt.Printf("- %s : %d\n", objet, qte)
	}
}

func AjouterObjet(nom string, quantite int) {
	Sacoche[nom] += quantite
	fmt.Printf("✅ %d %s ajouté(s) à la sacoche.\n", quantite, nom)
}

func UtiliserObjet(nom string, joueur *Personnage) {
	if Sacoche[nom] > 0 {
		switch nom {
		case "red bull":
			joueur.HP += 20
			Sacoche[nom]--
			fmt.Printf("%s utilise un red bull (+20 PV). PV actuels: %d\n", joueur.Nom, joueur.HP)
		case "Eau":
			joueur.HP += 10
			Sacoche[nom]--
			fmt.Printf("%s utilise une Eau (+10 PV). PV actuels: %d\n", joueur.Nom, joueur.HP)

		default:
			fmt.Println("❌ Objet non utilisable.")
		}
	} else {
		fmt.Printf("❌ Vous n’avez pas de %s.\n", nom)
	}
}

type Personnage struct {
	Nom          string
	Classe       string
	Niveau       int
	HP           int
	Max_hp       int
	Vitesse      int
	Force        int
	Intelligence int
	Resistance   int
	Chance       int
}

type Objet struct {
	Nom      string
	Quantite int
}
