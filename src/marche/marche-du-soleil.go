package main

import (
	"fmt"
)

// Buff appliqué aux stats
type Buff struct {
	Health       int
	Energy       int
	Intelligence int
	Defense      int
}

// Structure Item
type Item struct {
	Name  string
	Price int
	Buff  Buff
}

// Exemple de structure Personnage (tu peux remplacer par la tienne)
type Character struct {
	Name         string
	Money        int
	Health       int
	Energy       int
	Intelligence int
	Defense      int
	Bag          []Item
}

// Appliquer l'effet d'un item sur un joueur
func applyBuff(c *Character, item Item) {
	c.Health += item.Buff.Health
	c.Energy += item.Buff.Energy
	c.Intelligence += item.Buff.Intelligence
	c.Defense += item.Buff.Defense
}

// Afficher le marché
func showMarket(items []Item) {
	fmt.Println("\n--- Marché du Soleil 🌞 ---")
	for i, item := range items {
		fmt.Printf("%d) %s - %d pièces | Effets: Santé %+d, Énergie %+d, Intel %+d, Défense %+d\n",
			i+1, item.Name, item.Price,
			item.Buff.Health, item.Buff.Energy, item.Buff.Intelligence, item.Buff.Defense)
	}
	fmt.Println("0) Quitter le marché")
}

// Fonction d'achat (cœur du marché)
func buyItem(c *Character, items []Item, choice int) {
	if choice < 1 || choice > len(items) {
		fmt.Println("❌ Choix invalide.")
		return
	}

	item := items[choice-1]
	if c.Money < item.Price {
		fmt.Println("❌ Pas assez de kichta.")
		return
	}

	// Achat réussi
	c.Money -= item.Price
	c.Bag = append(c.Bag, item)
	applyBuff(c, item)
	fmt.Printf("✅ %s a acheté %s pour %d pièces !\n", c.Name, item.Name, item.Price)
}

// Exemple de fonction principale
func main() {
	items := []item{
		{"rtx 5070", 500, "+100 puissance graphique"},
		{"red bull", 10, "+20 santé"},
		{"ventoline", 25, "+30 respiration"},
		{"hérisson", 60, "+15 défense (piquant)"},
		{"bissap", 15, "+10 vitalité"},
		{"seringue", 5, "+5 soin rapide"},
		{"eau", 2, "+5 santé"},
		{"puff", 20, "-5 santé, +15 détente"},
		{"snus", 15, "-3 santé, +10 concentration"},
		{"nerd", 8, "+5 intelligence"},
		{"manuel de soumission", 125, "+5 resistance"},
		{"shambala", 150, "+5 chance, +3vitesse"},
		{"chicha", 30, "-5 santé, +3 vitesse"},
	}
	player := player{money: 100, inv: []item{}}
	fmt.Println(items[0], player)
}
