package four

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"PROJETRED/src/class"
)

type Item struct {
	Name  string
	Price int
	Buff  func(p *class.Personnage)
}


func ItemsForge() []Item {
	return []Item{
		{"Sacoche Burberry", 120, func(p *class.Personnage) {
			p.Chance += 30
			p.Intelligence += 20
		}},
		{"Casquette Gucci Fraise", 100, func(p *class.Personnage) {
			p.Vitesse += 40
		}},
		{"TN", 150, func(p *class.Personnage) {
			p.Force += 50
			p.Resistance += 20
		}},
	}
}


func showForge(items []Item) {
	fmt.Println("\n--- 🔥 Le Four (Forge) 🔥 ---")
	for i, item := range items {
		fmt.Printf("%d) %s - %d kishta\n", i+1, item.Name, item.Price)
	}
	fmt.Println("Écris 'tess' pour retourner à la tess.")
}

func acheterForge(p *class.Personnage, item Item) {
	if p.Kishta < item.Price {
		fmt.Println("❌ Pas assez de kishta pour cet objet !")
		return
	}
	p.Kishta -= item.Price


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


	item.Buff(p)
	fmt.Printf("✅ %s forgé avec succès dans Le Four !\n", item.Name)
}

func EntrerForge(p *class.Personnage, showStats func(*class.Personnage)) {
	items := ItemsForge()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		showStats(p)
		showForge(items)

		fmt.Print("Choix : ")
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
		acheterForge(p, items[num-1])
	}
}
