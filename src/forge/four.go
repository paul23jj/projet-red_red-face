package forge

import (
	class "PROJETRED/src/class"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// --- Item spécial du Four ---
type Item struct {
	Name          string
	RequiredItems []class.Inventaire // Items nécessaires pour crafter
	Buff          func(p *class.Personnage)
}

// --- Liste des items craftables dans Le Four ---
func ItemsForge() []Item {
	return []Item{
		{
			Name: "Sacoche Burberry",
			RequiredItems: []class.Inventaire{
				{Name: "Pantalon de la Municipale", Quantity: 1},
				{Name: "Holster de la BAC", Quantity: 1},
			},
			Buff: func(p *class.Personnage) {
				p.Chance += 30
				p.Intelligence += 20
			},
		},
		{
			Name: "Casquette Gucci Fraise",
			RequiredItems: []class.Inventaire{
				{Name: "Casque de CRS", Quantity: 1},
				{Name: "Puff goût Fraise", Quantity: 1},
			},
			Buff: func(p *class.Personnage) {
				p.Vitesse += 40
			},
		},
		{
			Name: "TN",
			RequiredItems: []class.Inventaire{
				{Name: "Bottes de Big Show", Quantity: 1},
			},
			Buff: func(p *class.Personnage) {
				p.Force += 50
				p.Resistance += 20
			},
		},
	}
}

// --- Affichage de la forge ---
func showForge(items []Item) {
	fmt.Println("\n--- 🔥 Le Four (Forge) 🔥 ---")
	for i, item := range items {
		fmt.Printf("%d) %s - Requiert: ", i+1, item.Name)
		for j, req := range item.RequiredItems {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s x%d", req.Name, req.Quantity)
		}
		fmt.Println()
	}
	fmt.Println("Écris 'tess' pour retourner à la tess.")
}

// --- Achat ---
func acheterForge(p *class.Personnage, item Item) {
	// Vérifier si le joueur a tous les items requis
	for _, req := range item.RequiredItems {
		found := false
		for _, inv := range p.Saccoche {
			if inv.Name == req.Name && inv.Quantity >= req.Quantity {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("❌ Pas assez de %s pour forger %s !\n", req.Name, item.Name)
			return
		}
	}

	// Retirer les items requis de la saccoche
	for _, req := range item.RequiredItems {
		for i, inv := range p.Saccoche {
			if inv.Name == req.Name {
				p.Saccoche[i].Quantity -= req.Quantity
				if p.Saccoche[i].Quantity == 0 {
					// Supprimer l'item si quantité = 0
					p.Saccoche = append(p.Saccoche[:i], p.Saccoche[i+1:]...)
				}
				break
			}
		}
	}

	// Ajouter l'item forgé à la saccoche
	found := false
	for i, it := range p.Saccoche {
		if it.Name == item.Name {
			p.Saccoche[i].Quantity++
			found = true
			break
		}
	}
	if !found {
		p.Saccoche = append(p.Saccoche, class.Inventaire{Name: item.Name, Quantity: 1})
	}

	// Appliquer le buff
	item.Buff(p)
	fmt.Printf("✅ %s forgé avec succès dans Le Four !\n", item.Name)
}

// --- Fonction principale pour entrer dans Le Four ---
func EntrerForge(p *class.Personnage, showStats func(*class.Personnage)) {
	// Débogage : Afficher Saccoche immédiatement à l'entrée
	fmt.Println("--- Debug Saccoche (entrée dans EntrerForge) ---")
	if len(p.Saccoche) == 0 {
		fmt.Println("Saccoche vide")
	} else {
		for _, item := range p.Saccoche {
			fmt.Printf("- %s x%d\n", item.Name, item.Quantity)
		}
	}

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
			// Débogage : Afficher Saccoche à la sortie
			fmt.Println("--- Debug Saccoche (sortie) ---")
			if len(p.Saccoche) == 0 {
				fmt.Println("Saccoche vide")
			} else {
				for _, item := range p.Saccoche {
					fmt.Printf("- %s x%d\n", item.Name, item.Quantity)
				}
			}
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
