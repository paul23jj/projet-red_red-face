package inventaire

import (
	class "PROJETRED/src/class"
	Monstre "PROJETRED/src/monstre"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func UtiliserObjetParNumero(joueur *class.Personnage, ennemi *Monstre.Monstre) {
	if len(joueur.Saccoche) == 0 {
		fmt.Println("Ta sacoche est vide !")
		return
	}
	fmt.Println("\nObjets disponibles :")
	for i, it := range joueur.Saccoche {
		fmt.Printf("%d) %s x%d\n", i+1, it.Name, it.Quantity)
	}
	fmt.Print("Tape le numéro de l'objet à utiliser : ")
	reader := bufio.NewReader(os.Stdin)
	choixStr, _ := reader.ReadString('\n')
	choixStr = strings.TrimSpace(choixStr)
	choix, err := strconv.Atoi(choixStr)
	if err != nil || choix < 1 || choix > len(joueur.Saccoche) {
		fmt.Println("Choix invalide.")
		return
	}
	obj := &joueur.Saccoche[choix-1]
	if obj.Quantity <= 0 {
		fmt.Println("Tu n'en as plus !")
		return
	}
	switch obj.Name {
	case "Red bull":
		joueur.HP += 20
		obj.Quantity--
		fmt.Printf("%s utilise un Red Bull (+20 PV). PV actuels: %d\n", joueur.Nom, joueur.HP)
	case "Eau":
		joueur.HP += 10
		obj.Quantity--
		fmt.Printf("%s utilise une Eau (+10 PV). PV actuels: %d\n", joueur.Nom, joueur.HP)
	case "Bissap":
		joueur.HP += 15
		obj.Quantity--
		fmt.Printf("%s utilise un Bissap (+15 PV). PV actuels: %d\n", joueur.Nom, joueur.HP)
	case "Seringue":
		joueur.HP += 25
		obj.Quantity--
		fmt.Printf("%s utilise une Seringue (+25 PV). PV actuels: %d\n", joueur.Nom, joueur.HP)
	case "Vodka":
		joueur.Force += 5
		obj.Quantity--
		fmt.Printf("%s utilise une Vodka (+5 Force, -5 PV). Force actuelle: %d, PV actuels: %d\n", joueur.Nom, joueur.Force, joueur.HP-5)
	case "Hérisson":
		joueur.Resistance += 5
		obj.Quantity--
		fmt.Printf("%s utilise un Hérisson (+5 Résistance). Résistance actuelle: %d\n", joueur.Nom, joueur.Resistance)
	case "Manuel de soumission":
		joueur.Intelligence += 5
		obj.Quantity--
		fmt.Printf("%s utilise un Manuel de soumission (+5 Intelligence). Intelligence actuelle: %d\n", joueur.Nom, joueur.Intelligence)
	case "Snus":
		joueur.Intelligence += 3
		obj.Quantity--
		fmt.Printf("%s utilise un Snus (+3 Intelligence, -3 PV). Intelligence actuelle: %d, PV actuels: %d\n", joueur.Nom, joueur.Intelligence, joueur.HP-3)
	case "Puff goût fraise":
		joueur.HP -= 5
		obj.Quantity--
		fmt.Printf("%s utilise un Puff goût fraise (-5 PV). PV actuels: %d\n", joueur.Nom, joueur.HP)
	case "RTX 5070":
		joueur.Intelligence += 20
		obj.Quantity--
		fmt.Printf("%s utilise une RTX 5070 (+20 Intelligence). Intelligence actuelle: %d\n", joueur.Nom, joueur.Intelligence)
	case "Ventoline":
		joueur.Vitesse += 10
		obj.Quantity--
		fmt.Printf("%s utilise une Ventoline (+10 Vitesse). Vitesse actuelle: %d\n", joueur.Nom, joueur.Vitesse)
	case "Shamballa":
		joueur.Chance += 5
		obj.Quantity--
		fmt.Printf("%s utilise un Shamballa (+5 Chance). Chance actuelle: %d\n", joueur.Nom, joueur.Chance)
	case "Pain":
		var IsPain bool = false
		var count int = 3
		if IsPain {
			if count != 0 {
				count--
				ennemi.HP -= 10
			} else {
				IsPain = false
			}
			fmt.Printf("%s subit les dégâts des pigeons (-10 PV). PV actuels: %d\n", ennemi.Nom, ennemi.HP)
		}
	default:
		fmt.Println("❌ Objet non utilisable.")
	}
}
