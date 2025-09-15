<<<<<<< HEAD
package menuPrincipal

import (
	class "PROJETRED/src/class"
	combat "PROJETRED/src/combat-system"
	monster "PROJETRED/src/monster"
=======
package MenuPrincipal

import (
	class "PROJETRED/src/class"

>>>>>>> refs/remotes/origin/main
	"fmt"
)

func MenuPrincipal(player class.Personnage) {
	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1) Aller se battre")
		fmt.Println("2) Voir inventaire")
		fmt.Println("3) Quitter")
		fmt.Print("Choix: ")
<<<<<<< HEAD
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
=======

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:

>>>>>>> refs/remotes/origin/main
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
