package menuPrincipal

import (
	class "PROJETRED/src/class"
	combat "PROJETRED/src/combat"
	"fmt"
)

func MenuPrincipal(Personnage class.Personnage) {
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
			monstre := combat.Monstre{} // Use the Monstre type from the combat package
			combat.Combat(&Personnage, &monstre)

		case 2:
			fmt.Println("Inventaire:", Personnage.Inventaire)
		case 3:
			fmt.Println("À bientôt 👋")
			return
		default:
			fmt.Println("Choix invalide !")
		}
	}
}
