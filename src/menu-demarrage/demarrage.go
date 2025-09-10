package menuDemarrage

import (
	"PROJETRED/class"
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
		class.InitPlayer()
	} else {
		fmt.Println("Dommage... à bientôt !")
	}
}
