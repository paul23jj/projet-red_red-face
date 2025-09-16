package menuDemarrage

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func gererFour() {
    for {
        fmt.Println("\nBienvenue dans le Four !")
        fmt.Println("Que veux-tu faire ?")
        fmt.Println("1. Forger une arme")
        fmt.Println("2. Reposer")
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
            fmt.Println("Tu as forgé une épée !")
        case "2":
            fmt.Println("Tu te reposes et récupères 20 HP.")
        case "3":
            fmt.Println("Retour au menu principal...")
            return
        default:
            fmt.Println("Option invalide, réessaie.")
        }
        reader.ReadString('\n') // Vider le buffer
    }
}
