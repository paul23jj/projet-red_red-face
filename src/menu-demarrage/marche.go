package menuDemarrage

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func gererMarche() {
    for {
        fmt.Println("\nBienvenue au Marché !")
        fmt.Println("Que veux-tu acheter ?")
        fmt.Println("1. Potion de soin (50 pièces)")
        fmt.Println("2. Armure légère (100 pièces)")
        fmt.Println("3. Retourner au menu principal")
        fmt.Print("Choisis une option : ")

        reader := bufio.NewReader(os.Stdin)
        choice, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Erreur de lecture, réessaie.")
            reader.ReadString('\n') // Vider le buffer
            continue
        }
        choice = strings.TrimSpace(choice)

        switch choice {
        case "1":
            fmt.Println("Tu as acheté une potion de soin !")
        case "2":
            fmt.Println("Tu as acheté une armure légère !")
        case "3":
            fmt.Println("Retour au menu principal...")
            return
        default:
            fmt.Println("Option invalide, réessaie.")
        }
        reader.ReadString('\n') // Vider le buffer
    }
}
