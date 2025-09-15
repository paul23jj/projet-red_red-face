package four

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// --- Structures importées depuis main ---
// On suppose que tu as déjà ces structs dans ton projet.
// Si elles sont dans un autre package, importe-les correctement.
type Inventaire struct {
	Name     string
	Quantity int
}

type Personnage struct {
	Classe       string
	Hp           int
	MaxHp        int
	Vitesse      int
	Force        int
	Intelligence int
	Resistance   int
	Chance       int
	Kishta       int
	Inventaire   []Inventaire
}

// --- Item spécial du Four ---
type Item struct {
	Name  string
	Price int
	Buff  func(p *Personnage)
}

// --- Liste des items craftables dans Le Four ---
func ItemsForge() []Item {
	return []Item{
		{"Sacoche Burberry", 120, func(p *Personnage) {
			p.Chance += 30
			p.Intelligence += 20
		}},
		{"Casquette Gucci Fraise", 100, func(p *Personnage) {
			p.Vitesse += 40
		}},
		{"TN", 150, func(p *Personnage) {
			p.Force += 50
			p.Resistance += 20
		}},
	}
}

// --- Affichage de la forge ---
func showForge(items []Item) {
	fmt.Println("\n--- 🔥 Le Four (Forge) 🔥 ---")
	for i, item := range items {
		fmt.Printf("%d) %s - %d kishta\n", i+1, item.Name, item.Price)
	}
	fmt.Println("Écris 'tess' pour retourner à la tess.")
}

// --- Achat ---
func acheterForge(p *Personnage, item Item) {
	if p.Kishta < item.Price {
		fmt.Println("❌ Pas assez de kishta pour cet objet !")
		return
	}
	p.Kishta -= item.Price

	// Ajouter à l’inventaire
	found := false
	for i, it := range p.Inventaire {
		if it.Name == item.Name {
			p.Inventaire[i].Quantity++
			found = true
			break
		}
	}
	if !found {
		p.Inventaire = append(p.Inventaire, Inventaire{Name: item.Name, Quantity: 1})
	}

	// Appliquer le buff
	item.Buff(p)
	fmt.Printf("✅ %s forgé avec succès dans Le Four !\n", item.Name)
}

// --- Fonction principale pour entrer dans Le Four ---
func EntrerForge(p *Personnage, showStats func(*Personnage)) {
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
