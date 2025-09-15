package menuDemarrage

import (
	Class "PROJETRED/src/class"
	menu "PROJETRED/src/menu-principal"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartMenu() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Bienvenue dans Projet-Red ===")
	fmt.Print("Veux-tu rentrer dans la tess ? (oui/non) : ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(strings.ToLower(choice))

	if choice == "oui" {
		player := Class.InitPlayer() // ✅ ici tu récupères ton joueur
		menu.MenuPrincipal(player)   // ✅ tu envoies le joueur vers ton menu principal
	} else {
		fmt.Println("Dommage... à bientôt !")
	}
}
