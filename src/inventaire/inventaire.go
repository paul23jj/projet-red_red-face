package inventaire

import (
	class "PROJETRED/src/class"
	"fmt"
)

var Sacoche = map[string]int{}
var MaxSlots = 10

func AfficherSacoche(joueur *class.Personnage) {
	fmt.Println("\n📦 Sacoche :")
	if len(joueur.Saccoche) == 0 {
		fmt.Println("(vide)")
		return
	}
	for _, it := range joueur.Saccoche {
		fmt.Printf("- %s : %d\n", it.Name, it.Quantity)
	}
	fmt.Printf("\n🎒 Slots utilisés : %d/%d (reste %d)\n",
		len(joueur.Saccoche), MaxSlots, MaxSlots-len(joueur.Saccoche))
}

func AjouterObjet(joueur *class.Personnage, nom string, quantite int) {
	for i, it := range joueur.Saccoche {
		if it.Name == nom {
			joueur.Saccoche[i].Quantity += quantite
			fmt.Printf("✅ %d %s ajouté(s) à la sacoche.\n", quantite, nom)
			return
		}
	}
	joueur.Saccoche = append(joueur.Saccoche, class.Inventaire{Name: nom, Quantity: quantite})
	fmt.Printf("✅ %d %s ajouté(s) à la sacoche.\n", quantite, nom)
}

func UtiliserObjet(nom string, joueur *class.Personnage) {
	for i, it := range joueur.Saccoche {
		if it.Name == nom && it.Quantity > 0 {
			switch nom {
			case "Red Bull":
				joueur.HP += 20
				joueur.Saccoche[i].Quantity--
				fmt.Printf("%s utilise un red bull (+20 PV). PV actuels: %d\n", joueur.Nom, joueur.HP)
			case "Eau":
				joueur.HP += 10
				joueur.Saccoche[i].Quantity--
				fmt.Printf("%s utilise une Eau (+10 PV). PV actuels: %d\n", joueur.Nom, joueur.HP)
			default:
				fmt.Println("❌ Objet non utilisable.")
			}
			return
		}
	}
	fmt.Printf("❌ Vous n’avez pas de %s.\n", nom)
}
