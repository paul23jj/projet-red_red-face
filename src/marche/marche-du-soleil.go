package marche

import (
	class "PROJETRED/src/class"
	inventaire "PROJETRED/src/inventaire"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Inventaire struct {
	Name     string
	Quantity int
}

type Item struct {
	Name         string
	Price        int
	BuffNormal   func(p *class.Personnage)
	BuffFavori   func(p *class.Personnage)
	FavoriClasse string // si vide -> pas d’item favori
	PriceFavori  int
}

// ---------------- Fonctions ----------------

// Applique un buff de soin avec limite au max_hp
func Heal(p *class.Personnage, amount int) {
	p.HP += amount
	if p.HP > p.MaxHP {
		p.HP = p.MaxHP
	}
	// Si le choix ne correspond à aucun item, retourner un item vide
}

// Affiche le marché
func ShowMarket(items []Item) {
	fmt.Println("\n--- 🌞 Marché du Soleil 🌞 ---")
	for i, item := range items {
		fmt.Printf("%d) %s - %d kishta\n", i+1, item.Name, item.Price)
	}
	fmt.Println("Écris 'tess' pour retourner à la tess.")
}

// Affiche les stats du joueur
func ShowStats(p *class.Personnage) {
	fmt.Printf("\n--- Stats de ton perso (%s) ---\n", p.Classe)
	fmt.Printf("HP: %d/%d | Force: %d | Vitesse: %d | Intel: %d | Résistance: %d | Chance: %d | Kishta: %d\n",
		p.HP, p.MaxHP, p.Force, p.Vitesse, p.Intelligence, p.Resistance, p.Chance, p.Kishta)
	fmt.Println("Saccoche :")
	if len(p.Saccoche) == 0 {
		fmt.Println("Saccoche vide !")
	} else {
		for _, it := range p.Saccoche {
			fmt.Printf(" - %s x%d\n", it.Name, it.Quantity)
		}
		// Si aucun choix valide, retourner un item vid
	}
}

// Achat d’un item
func acheterItem(p *class.Personnage, Item Item) {
	prix := Item.Price
	buff := Item.BuffNormal

	// Si c’est l’item favori du perso
	if Item.FavoriClasse == p.Classe {
		prix = Item.PriceFavori
		buff = Item.BuffFavori
	}

	if p.Kishta < prix {
		fmt.Println("❌ Pas assez de kishta !")
		return
	}

	// Vérifier la limite de slots
	if len(p.Saccoche) >= inventaire.MaxSlots {
		// sauf si l’objet existe déjà (stackable)
		found := false
		for _, it := range p.Saccoche {
			if it.Name == Item.Name {
				found = true
				break
			}
		}
		if !found && Item.Name != "Sacoche +" {
			fmt.Println("❌ Sacoche pleine, impossible d’ajouter de nouveaux objets.")
			return
		}
	}

	// Retirer l’argent
	p.Kishta -= prix

	// Ajouter à l’inventaire
	found := false
	for i, it := range p.Saccoche {
		if it.Name == Item.Name {
			p.Saccoche[i].Quantity++
			found = true
			break
		}
	}
	if !found {
		p.Saccoche = append(p.Saccoche, class.Inventaire{Name: Item.Name, Quantity: 1})
	}

	// Appliquer le buff
	if buff != nil {
		buff(p)
	}

	fmt.Printf("✅ Tu as acheté %s pour %d kishta !\n", Item.Name, prix)
}
func MarcheDuSoleil(p *class.Personnage) {
	items := []Item{
		{"Hérisson", 40, func(p *class.Personnage) { p.Resistance += 10 }, func(p *class.Personnage) { p.Resistance += 20 }, "Nomade", 20},
		{"Vodka", 30, func(p *class.Personnage) { p.Force += 10; p.HP -= 5 }, func(p *class.Personnage) { p.Force += 20; p.HP -= 5 }, "Russe", 15},
		{"Manuel de soumission", 50, func(p *class.Personnage) { p.Intelligence += 15 }, func(p *class.Personnage) { p.Intelligence += 25 }, "Tchetchene", 25},
		{"Bissap", 25, func(p *class.Personnage) { Heal(p, 15) }, func(p *class.Personnage) { Heal(p, 30) }, "Malien", 12},
		{"Shamballa", 40, func(p *class.Personnage) { p.Chance += 10 }, func(p *class.Personnage) { p.Chance += 20 }, "Bresilien", 20},
		{"Red bull", 15, func(p *class.Personnage) { p.Vitesse += 10; p.HP -= 5 }, nil, "", 0},
		{"Ventoline", 20, func(p *class.Personnage) { p.Vitesse += 15 }, nil, "", 0},
		{"Seringue", 5, func(p *class.Personnage) { Heal(p, 10) }, nil, "", 0},
		{"Eau", 2, func(p *class.Personnage) { Heal(p, 5) }, nil, "", 0},
		{"Puff goût fraise", 20, func(p *class.Personnage) { p.HP -= 5 }, nil, "", 0},
		{"Snus", 15, func(p *class.Personnage) { p.HP -= 3; p.Intelligence += 10 }, nil, "", 0},
		{"RTX 5070", 80, func(p *class.Personnage) { p.Intelligence += 50 }, nil, "", 0},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		ShowStats(p)
		ShowMarket(items)

		fmt.Print("\nQue veux-tu acheter ? (numéro ou 'tess') : ")
		scanner.Scan()
		choix := strings.TrimSpace(scanner.Text())

		if choix == "tess" {
			fmt.Println("👉 Tu es retourné à la tess.")
			break
		}

		num, err := strconv.Atoi(choix)
		if err != nil || num < 1 || num > len(items) {
			fmt.Println("Choix invalide.")
			continue
		}
		acheterItem(p, items[num-1])
	}

}
