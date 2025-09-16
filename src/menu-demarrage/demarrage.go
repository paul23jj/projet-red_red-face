package menuDemarrage

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    class "PROJETRED/src/class"
    combat "PROJETRED/src/combat"
    inventaire "PROJETRED/src/inventaire"
    monstre "PROJETRED/src/monstre"
    four "PROJETRED/src/four"
    marche "PROJETRED/src/marche"
)

var player class.Personnage
var monstre monstre.Monstre

func StartMenu() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("=== Bienvenue dans Projet-Red ===")
    fmt.Print("Veux-tu rentrer dans la tess ? (oui/non) : ")
    choice, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Erreur de lecture, à bientôt !")
        return
    }
    choice = strings.TrimSpace(strings.ToLower(choice))

    if choice != "oui" {
        fmt.Println("Dommage... à bientôt !")
        return
    }

    player = class.InitPlayer()

    // Initialiser Sacoche pour éviter les erreurs nil map
    if player.Sacoche == nil {
        player.Sacoche = make(map[string]int)
    }
    // Synchroniser Inventaire (slice) avec Sacoche (map)
    for _, item := range player.Inventaire {
        player.Sacoche[item.Name] = item.Quantity
    }

    for {
        fmt.Println("\n=== Menu Principal ===")
        fmt.Println("1. Aller dans le Four")
        fmt.Println("2. Aller au Marché")
        fmt.Println("3. Regarder la sacoche")
        fmt.Println("4. Chercher un tête à tête")
        fmt.Println("5. Quitter")
        fmt.Print("Choisis une option : ")

<<<<<<< HEAD
        menuChoice, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Erreur de lecture, réessaie.")
            continue
        }
        menuChoice = strings.TrimSpace(menuChoice)
        switch menuChoice {
        case "1":
            fmt.Println("Tu es maintenant dans le Four !")
            gererFour()
        case "2":
            fmt.Println("Tu es maintenant au Marché !")
            gererMarche()
        case "3":
            fmt.Println("Voici ta sacoche :")
            inventaire.AfficherSacoche(&player)
        case "4":
            fmt.Println("Tu cherches un tête à tête...")
            monstre = monstre.GenererMonstre()
            combat.Combat(&player, &monstre) // Ou combat.TourPartoutCombat si tu veux utiliser la nouvelle version
        case "5":
            fmt.Println("À bientôt !")
            os.Exit(0)
        default:
            fmt.Println("Option invalide, réessaie.")
        }
=======
		switch menuChoice {
		case "1":
			fmt.Println("Tu es maintenant dans le Four !")
			gererFour()
		case "2":
			fmt.Println("Tu es maintenant au Marché !")
			gererMarche()
		case "3":
			fmt.Println("Voici ta sacoche :")
			Inventaire.AfficherSacoche(&player)
		case "4":
			fmt.Println("Tu cherches un tête à tête...")
			monstre = Monstre.GenererMonstre()
			Combat.CombatMain(&player, &monstre)
		case "5":
			fmt.Println("À bientôt !")
			os.Exit(0)
		default:
			fmt.Println("Option invalide, réessaie.")
		}
>>>>>>> be7731b85df61a943c98fac909c47daabc977b75

        // Vider le buffer
        reader.ReadString('\n')
    }
}

func gererFour() {
    // Fonction pour afficher les stats compatible avec four.Personnage
    showStats := func(p *four.Personnage) {
        fmt.Printf("\n--- Stats de ton perso (%s) ---\n", p.Classe)
        fmt.Printf("HP: %d/%d | Force: %d | Vitesse: %d | Intel: %d | Résistance: %d | Chance: %d | Kishta: %d\n",
            p.Hp, p.MaxHp, p.Force, p.Vitesse, p.Intelligence, p.Resistance, p.Chance, p.Kishta)
        fmt.Println("Inventaire :")
        if len(p.Inventaire) == 0 {
            fmt.Println(" (vide)")
        } else {
            for _, it := range p.Inventaire {
                fmt.Printf(" - %s x%d\n", it.Name, it.Quantity)
            }
        }
    }

    // Convertir class.Personnage en four.Personnage
    pFour := &four.Personnage{
        Classe:       player.Classe,
        Hp:           player.HP,
        MaxHp:        player.MaxHP,
        Vitesse:      player.Vitesse,
        Force:        player.Force,
        Intelligence: player.Intelligence,
        Resistance:   player.Resistance,
        Chance:       player.Chance,
        Kishta:       player.Kishta, // Utiliser la valeur de Kishta existante ou 100 si non défini
        Inventaire:   []four.Inventaire{},
    }

    // Copier les items de Sacoche vers Inventaire pour compatibilité
    for name, quantity := range player.Sacoche {
        pFour.Inventaire = append(pFour.Inventaire, four.Inventaire{Name: name, Quantity: quantity})
    }

    // Appeler la fonction du Four
    four.EntrerForge(pFour, showStats)

    // Synchroniser les changements vers player
    player.HP = pFour.Hp
    player.MaxHP = pFour.MaxHp
    player.Vitesse = pFour.Vitesse
    player.Force = pFour.Force
    player.Intelligence = pFour.Intelligence
    player.Resistance = pFour.Resistance
    player.Chance = pFour.Chance
    player.Kishta = pFour.Kishta
    // Synchroniser l'inventaire
    player.Sacoche = make(map[string]int) // Réinitialiser pour éviter les doublons
    for _, item := range pFour.Inventaire {
        player.Sacoche[item.Name] = item.Quantity
    }
}

func gererMarche() {
    // Fonction pour afficher les stats compatible avec marche.Personnage
    showStatsMarche := func(p *marche.Personnage) {
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

    // Convertir class.Personnage en marche.Personnage
    pMarche := &marche.Personnage{
        classe:       strings.ToLower(player.Classe), // marche utilise des minuscules
        hp:           player.HP,
        max_hp:       player.MaxHP,
        vitesse:      player.Vitesse,
        force:        player.Force,
        intelligence: player.Intelligence,
        resistance:   player.Resistance,
        chance:       player.Chance,
        kishta:       player.Kishta, // Utiliser la valeur existante ou 100 si non défini
        inventaire:   []marche.Inventaire{},
    }

    // Copier les items de Sacoche vers inventaire pour compatibilité
    for name, quantity := range player.Sacoche {
        pMarche.inventaire = append(pMarche.inventaire, marche.Inventaire{name: name, quantity: quantity})
    }

    // Appeler une boucle similaire à celle du package marche
    items := marche.Items() // Assumer que le package marche a une fonction Items() exportée
    scanner := bufio.NewScanner(os.Stdin)
    for {
        showStatsMarche(pMarche)
        marche.showMarket(items)

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

        marche.acheterItem(pMarche, items[num-1])
    }

    // Synchroniser les changements vers player
    player.HP = pMarche.hp
    player.MaxHP = pMarche.max_hp
    player.Vitesse = pMarche.vitesse
    player.Force = pMarche.force
    player.Intelligence = pMarche.intelligence
    player.Resistance = pMarche.resistance
    player.Chance = pMarche.chance
    player.Kishta = pMarche.kishta
    // Synchroniser l'inventaire
    player.Sacoche = make(map[string]int) // Réinitialiser pour éviter les doublons
    for _, item := range pMarche.inventaire {
        player.Sacoche[item.name] = item.quantity
    }
}
