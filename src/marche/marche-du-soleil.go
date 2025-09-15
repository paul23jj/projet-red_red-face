package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Buff appliqué aux stats
type Buff struct {
	Health      int
	Energy      int
	Intelligence int
	Defense     int
}

// Structure Item
type Item struct {
	Name  string
	Price int
	Buff  Buff
}

// Exemple de structure Personnage (tu peux remplacer par la tienne)
type Character struct {
	Name        string
	Money       int
	Health      int
	Energy      int
	Intelligence int
	Defense     int
	Bag         []Item
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
	// Liste des items disponibles au marché
	items := []Item{
		{"rtx 5070", 500, Buff{Health: 0, Energy: 0, Intelligence: +50, Defense: 0}},
		{"red bull", 10, Buff{Health: 0, Energy: +20, Intelligence: 0, Defense: 0}},
		{"ventoline", 25, Buff{Health: +10, Energy: +15, Intelligence: 0, Defense: 0}},
		{"hérisson", 60, Buff{Health: 0, Energy: 0, Intelligence: 0, Defense: +15}},
		{"bissap", 15, Buff{Health: +10, Energy: +5, Intelligence: 0, Defense: 0}},
		{"seringue", 5, Buff{Health: +5, Energy: 0, Intelligence: 0, Defense: 0}},
		{"eau", 2, Buff{Health: +2, Energy: +2, Intelligence: 0, Defense: 0}},
		{"puff", 20, Buff{Health: -5, Energy: +15, Intelligence: 0, Defense: 0}},
		{"snus", 15, Buff{Health: -3, Energy: +5, Intelligence: +10, Defense: 0}},
		{"nerd", 8, Buff{Health: 0, Energy: 0, Intelligence: +5, Defense: 0}},
	}

	// Exemple de personnage (tu remplaceras par les tiens)
	player := Character{
		Name:        "Aventurier",
		Money:       200,
		Health:      100,
		Energy:      50,
		Intelligence: 10,
		Defense:     5,
		Bag:         []Item{},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMarket(items)
		fmt.Printf("\n%s | 💰 Argent: %d | ❤️ Santé: %d | ⚡ Énergie: %d | 🧠 Intel: %d | 🛡️ Défense: %d\n",
			player.Name, player.Money, player.Health, player.Energy, player.Intelligence, player.Defense)

		fmt.Print("\nQue veux-tu acheter ? (numéro) : ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "0" {
			fmt.Println("👋 Merci d'avoir visité le Marché du Soleil 🌞 !")
			break
		}

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("❌ Entrée invalide.")
			continue
		}

		buyItem(&player, items, choice)
	}
}
