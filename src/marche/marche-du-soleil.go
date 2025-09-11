package marche

import (
	"fmt"
)

type item struct {
	name  string
	price int
	buff  string
}

type player struct {
	money int
	inv   []item
}

var items []item = []item{
	{"rtx 5070", 500, "+100 puissance graphique"},
	{"red bull", 10, "+20 énergie"},
	{"ventoline", 25, "+30 respiration"},
}

func showMarket(player player) {
	fmt.Println("\n--- Marché du Soleil 🌞---")
	for i, item := range items {
		fmt.Printf("%d) %s - %d kishta | Effet: %s\n", i+1, item.name, item.buff)
	}
	fmt.Println("0) Quitter le marché")
}

func showInventory(player player) {
	fmt.Println("\n--- ta saccoche ---")
	if len(player.inv) == 0 {
		fmt.Println("ta saccoche elle est vide mon gaté")
	} else {
		for _, item := range player.inv {
			fmt.Printf("- %s (Effet: %s)\n", item.name, item.buff)
		}
	}
	fmt.Printf("kishta restante: %d kishta\n", player.money)
}

func main() {
	items := []item{
		{"rtx 5070", 500, "+100 puissance graphique"},
		{"red bull", 10, "+20 énergie"},
		{"ventoline", 25, "+30 respiration"},
		{"hérisson", 60, "+15 défense (piquant)"},
		{"bissap", 15, "+10 HP"},
		{"seringue", 5, "+5 soin rapide"},
		{"eau", 2, "+5 hydratation"},
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


scanner := bufio.NewScanner(os.Stdin)
	for {
		showMarket(items)
		showInventory(player)

		fmt.Print("\nChoisis un article à acheter (numéro) ou 0 pour quitter: ")
		scanner.Scan()
		var choice int
		fmt.Sscanf(scanner.Text(), "%d", &choice)
        input := strings.TrimSpace(scanner.Text())
		if choice == 0 {
			break
		} else if choice > 0 && choice <= len(items) {
			item := items[choice-1]
			if player.money >= item.price {
				player.money -= item.price
				player.inv = append(player.inv, item)
				fmt.Printf("Tu as acheté %s pour %d kishta. effet appliqué: %s\n", item.name, item.price, item.buff)
			} else {
				fmt.Println("Tu n'as pas assez de kishta.")
			}
		} else {
			fmt.Println("Choix invalide.")
		}
	}
