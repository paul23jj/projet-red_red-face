package four

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// --- Structures importées depuis main ---
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
    Name          string
    RequiredItems []Inventaire // Items nécessaires pour crafter
    Buff          func(p *Personnage)
}

// --- Liste des items craftables dans Le Four ---
func ItemsForge() []Item {
    return []Item{
        {
            Name: "Sacoche Burberry",
            RequiredItems: []Inventaire{
                {Name: "Pantalon de la Municipale", Quantity: 1},
                {Name: "Holster de la BAC", Quantity: 1},
            },
            Buff: func(p *Personnage) {
                p.Chance += 30
                p.Intelligence += 20
            },
        },
        {
            Name: "Casquette Gucci Fraise",
            RequiredItems: []Inventaire{
                {Name: "Casque de CRS", Quantity: 1},
                {Name: "Puff goût Fraise", Quantity: 1},
            },
            Buff: func(p *Personnage) {
                p.Vitesse += 40
            },
        },
        {
            Name: "TN",
            RequiredItems: []Inventaire{
                {Name: "Bottes de Big Show", Quantity: 1},
            },
            Buff: func(p *Personnage) {
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
func acheterForge(p *Personnage, item Item) {
    // Vérifier si le joueur a tous les items requis
    for _, req := range item.RequiredItems {
        found := false
        for _, inv := range p.Inventaire {
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

    // Retirer les items requis de l'inventaire
    for _, req := range item.RequiredItems {
        for i, inv := range p.Inventaire {
            if inv.Name == req.Name {
                p.Inventaire[i].Quantity -= req.Quantity
                if p.Inventaire[i].Quantity == 0 {
                    // Supprimer l'item si quantité = 0
                    p.Inventaire = append(p.Inventaire[:i], p.Inventaire[i+1:]...)
                }
                break
            }
        }
    }

    // Ajouter l'item forgé à l'inventaire
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
