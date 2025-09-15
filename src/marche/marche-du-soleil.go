<<<<<<< HEAD
package marcheDuSoleil
=======
package marche
>>>>>>> refs/remotes/origin/main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------- Structures ----------------

type Inventaire struct {
	name     string
	quantity int
}

type Personnage struct {
	classe       string
	hp           int
	max_hp       int
	vitesse      int
	force        int
	intelligence int
	resistance   int
	chance       int
	kishta       int
	inventaire   []Inventaire
}

type Item struct {
	Name         string
	Price        int
	BuffNormal   func(p *Personnage)
	BuffFavori   func(p *Personnage)
	FavoriClasse string // si vide -> pas d’item favori
	PriceFavori  int
}

// ---------------- Fonctions ----------------

// Applique un buff de soin avec limite au max_hp
func heal(p *Personnage, amount int) {
	p.hp += amount
	if p.hp > p.max_hp {
		p.hp = p.max_hp
	}
}

// Affiche le marché
func showMarket(items []Item) {
	fmt.Println("\n--- 🌞 Marché du Soleil 🌞 ---")
	for i, item := range items {
		fmt.Printf("%d) %s - %d kishta\n", i+1, item.Name, item.Price)
	}
	fmt.Println("Écris 'tess' pour retourner à la tess.")
}

// Affiche les stats du joueur
func showStats(p *Personnage) {
	fmt.Printf("\n--- Stats de ton perso (%s) ---\n", p.classe)
	fmt.Printf("HP: %d/%d | Force: %d | Vitesse: %d | Intel: %d | Résistance: %d | Chance: %d | Kishta: %d\n",
		p.hp, p.max_hp, p.force, p.vitesse, p.intelligence, p.resistance, p.chance, p.kishta)
	fmt.Println("Inventaire :")
	if len(p.inventaire) == 0 {
		fmt.Println(" (vide)")
	} else {
		for _, it := range p.inventaire {
			fmt.Printf(" - %s x%d\n", it.name, it.quantity)
		}
	}
}

// Achat d’un item
func acheterItem(p *Personnage, item Item) {
	prix := item.Price
	buff := item.BuffNormal

	// Si c’est l’item favori du perso
	if item.FavoriClasse == p.classe {
		prix = item.PriceFavori
		buff = item.BuffFavori
	}

	if p.kishta < prix {
		fmt.Println("❌ Pas assez de kishta !")
		return
	}

	// Retirer l’argent
	p.kishta -= prix

	// Ajouter à l’inventaire
	found := false
	for i, it := range p.inventaire {
		if it.name == item.Name {
			p.inventaire[i].quantity++
			found = true
			break
		}
	}
	if !found {
		p.inventaire = append(p.inventaire, Inventaire{name: item.Name, quantity: 1})
	}

	// Appliquer le buff
	buff(p)

	fmt.Printf("✅ Tu as acheté %s pour %d kishta !\n", item.Name, prix)
}

// ---------------- Main ----------------

func main() {
	// Exemple : un joueur Russe
	p := Personnage{
		classe:       "Russe",
		hp:           100,
		max_hp:       100,
		vitesse:      3,
		force:        10,
		intelligence: 3,
		resistance:   7,
		chance:       3,
		kishta:       100, // argent de départ
		inventaire:   []Inventaire{},
	}

	// Liste des items
	items := []Item{
		{"Hérisson", 40,
			func(p *Personnage) { p.resistance += 10 },
			func(p *Personnage) { p.resistance += 20 },
			"Nomade", 20},
		{"Vodka", 30,
			func(p *Personnage) { p.force += 10; p.hp -= 5 },
			func(p *Personnage) { p.force += 20; p.hp -= 5 },
			"Russe", 15},
		{"Manuel de soumission", 50,
			func(p *Personnage) { p.intelligence += 15 },
			func(p *Personnage) { p.intelligence += 25 },
			"tchetchene", 25},
		{"Bissap", 25,
			func(p *Personnage) { heal(p, 15) },
			func(p *Personnage) { heal(p, 30) },
			"Malien", 12},
		{"Shamballa", 40,
			func(p *Personnage) { p.chance += 10 },
			func(p *Personnage) { p.chance += 20 },
			"Bresilien", 20},
		{"Red bull", 15,
			func(p *Personnage) { p.vitesse += 10; p.hp -= 5 },
			nil, "", 0},
		{"Ventoline", 20,
			func(p *Personnage) { p.vitesse += 15 },
			nil, "", 0},
		{"Seringue", 5,
			func(p *Personnage) { heal(p, 10) },
			nil, "", 0},
		{"Eau", 2,
			func(p *Personnage) { heal(p, 5) },
			nil, "", 0},
		{"Puff", 20,
			func(p *Personnage) { p.hp -= 5 }, // détente RP
			nil, "", 0},
		{"Snus", 15,
			func(p *Personnage) { p.hp -= 3; p.intelligence += 10 },
			nil, "", 0},
		{"Nerd", 8,
			func(p *Personnage) { p.intelligence += 5 },
			nil, "", 0},
		{"RTX 5070", 80,
			func(p *Personnage) { p.intelligence += 50 },
			nil, "", 0},
	}

	// Scanner
	scanner := bufio.NewScanner(os.Stdin)

	for {
		showStats(&p)
		showMarket(items)

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

		acheterItem(&p, items[num-1])
	}
}
