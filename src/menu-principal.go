package main
import (
	"fmt"
	class "PROJETRED/src/class"
	combat "PROJETRED/src/combat-system"
	monster "PROJETRED/src/monster"
)
func MenuPrincipal(player class.Personnage) {
    for {
        fmt.Println("\n=== MENU ===")
        fmt.Println("1) Aller se battre")
        fmt.Println("2) Voir inventaire")
        fmt.Println("3) Quitter")
        fmt.Print("Choix: ")

        var choix int
        fmt.Scan(&choix)

        switch choix {
        case 1:
            enemy := monster.GenererMonstre()

            playerCharacter := class.Character{
                Name:    player.Nom,
                HP:      player.HP,
                Attack:  player.Force,
                Defense: player.Resistance,
            }

            combat.StartCombat(playerCharacter, enemy)
        case 2:
            fmt.Println("Inventaire:", player.Inventaire)
        case 3:
            fmt.Println("À bientôt 👋")
            return
        default:
            fmt.Println("Choix invalide !")
        }
    }
}
