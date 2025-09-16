package inventaire

import (
	class "PROJETRED/src/class"
	"fmt"
)

var Sacoche = map[string]int{}
var MaxSlots = 10

func AfficherSacoche(joueur *class.Personnage) {
	fmt.Println("\n📦 Sacoche :")
	if len(joueur.Sacoche) == 0 {
		fmt.Println("(vide)")
		return
	}
	for objet, qte := range joueur.Sacoche {
		fmt.Printf("- %s : %d\n", objet, qte)
	}
	fmt.Printf("\n🎒 Slots utilisés : %d/%d (reste %d)\n",
		len(joueur.Sacoche), MaxSlots, MaxSlots-len(joueur.Sacoche))
}

func AjouterObjet(joueur *class.Personnage, nom string, quantite int) {
	joueur.Sacoche[nom] += quantite
	fmt.Printf("✅ %d %s ajouté(s) à la sacoche.\n", quantite, nom)
}

func UtiliserObjet(nom string, joueur *class.Personnage) {
	if joueur.Sacoche[nom] > 0 {
		switch nom {
		case "red bull":
			joueur.HP += 20
			joueur.Sacoche[nom]--
			fmt.Printf("%s utilise un red bull (+20 PV). PV actuels: %d\n", joueur.Nom, joueur.HP)
		case "Eau":
			joueur.HP += 10
			joueur.Sacoche[nom]--
			fmt.Printf("%s utilise une Eau (+10 PV). PV actuels: %d\n", joueur.Nom, joueur.HP)
		default:
			fmt.Println("❌ Objet non utilisable.")
		}
	} else {
		fmt.Printf("❌ Vous n’avez pas de %s.\n", nom)
	}
}
