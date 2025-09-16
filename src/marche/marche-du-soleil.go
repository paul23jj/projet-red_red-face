 package marche

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"PROJETRED/src/class"
)

type Item struct {
	Name         string
	Price        int
	BuffNormal   func(p *class.Personnage)
	BuffFavori   func(p *class.Personnage)
	FavoriClasse string
	PriceFavori  int
}

func heal(p *class.Personnage, amount int) {
	p.HP += amount
	if p.HP > p.MaxHP {
		p.HP = p.MaxHP
	}
}


func showMarket(items []Item) {
	fmt.Println("\n--- 🌞 Marché du Soleil 🌞 ---")
	for i, item := range items {
		fmt.Printf("%d) %s - %d kishta\n", i+1, item.Name, item.Price)
	}
	fmt.Println("Écris 'tess' pour retourner à la tess.")
}

func showStats(p *class.Personnage) {
	fmt.Printf("\n--- Stats de ton perso (%s) ---\n", p.Classe)
	fmt.Printf("HP: %d/%d | Force: %d | Vitesse: %d | Intel: %d | Résistance: %d | Chance: %d | Kishta: %d\n",
		p.HP, p.MaxHP, p.Force, p.Vitesse, p.Intelligence, p.Resistance, p.Chance, p.Kishta)
	fmt.Println("Inventaire :")
	if len(p.Inventaire) == 0 {
		fmt.Println(" (vide)")
	} else {
		for _, it := range p.Inventaire {
			fmt.Printf(" - %s x%d\n", it.Name, it.Quantity)
		}
	}
}

func acheterItem(p *class.Personnage, item Item) {
	prix := item.Price
	buff := item.BuffNormal

	if item.FavoriClasse == p.Classe {
		prix = item.PriceFavori
		buff = item.BuffFavori
	}

	if p.Kishta < prix {
		fmt.Println("❌ Pas assez de kishta !")
		return
	}

	p.Kishta -= prix
	found := false
	for i, it := range p.Inventaire {
		if it.Name == item.Name {
			p.Inventaire[i].Quantity++
			found = true
			break
		}
	}
	if !found {
		p.Inventaire = append(p.Inventaire, class.Inventaire{Name: item.Name, Quantity: 1})
	}


	buff(p)

	fmt.Printf("✅ Tu as acheté %s pour %d kishta !\n", item.Name, prix)
}

func EntrerMarche(p *class.Personnage) {
	items := []Item{
		{"Hérisson", 40, func(p *class.Personnage) { p.Resistance += 10 },
			func(p *class.Personnage) { p.Resistance += 20 }, "Nomade", 20},
		{"Vodka", 30, func(p *class.Personnage) { p.Force += 10; p.HP -= 5 },
			func(p *class.Personnage) { p.Force += 20; p.HP -= 5 }, "Russe", 15},
		{"Manuel de soumission", 50, func(p *class.Personnage) { p.Intelligence += 15 },
			func(p *class.Personnage) { p.Intelligence += 25 }, "Tchetchene", 25},
		{"Bissap", 25, func(p *class.Personnage) { heal(p, 15) },
			func(p *class.Personnage) { heal(p, 30) }, "Malien", 12},
		{"Shamballa", 40, func(p *class.Personnage) { p.Chance += 10 },
			func(p *class.Personnage) { p.Chance += 20 }, "Bresilien", 20},
		{"Red bull", 15, func(p *class.Personnage) { p.Vitesse += 10; p.HP -= 5 }, nil, "", 0},
		{"Ventoline", 20, func(p *class.Personnage) { p.Vitesse += 15 }, nil, "", 0},
		{"Seringue", 5, func(p *class.Personnage) { heal(p, 10) }, nil, "", 0},
		{"Eau", 2, func(p *class.Personnage) { heal(p, 5) }, nil, "", 0},
		{"Puff", 20, func(p *class.Personnage) { p.HP -= 5 }, nil, "", 0},
		{"Snus", 15, func(p *class.Personnage) { p.HP -= 3; p.Intelligence += 10 }, nil, "", 0},
		{"Nerd", 8, func(p *class.Personnage) { p.Intelligence += 5 }, nil, "", 0},
		{"RTX 5070", 80, func(p *class.Personnage) { p.Intelligence += 50 }, nil, "", 0},
	}


	scanner := bufio.NewScanner(os.Stdin)

	for {
		showStats(p)
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
		acheterItem(p, items[num-1])
	}
}
