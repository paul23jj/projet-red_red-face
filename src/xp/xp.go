package xp

import (
	class "PROJETRED/src/class"
	"fmt"
)

// GainXP ajoute de l’XP au joueur et gère la montée de niveau
func GainXP(player *class.Personnage, amount int) {
	player.Niveau += amount
	fmt.Printf("✨ %s a gagné %d points d'expérience ! ✨\n", player.Nom, amount)

	// Tous les 100 XP → le joueur augmente de niveau
	if player.Niveau >= 100 {
		LevelUp(player)
	}
}

// LevelUp augmente les stats du joueur
func LevelUp(player *class.Personnage) {
	player.Niveau++
	fmt.Printf("🎉 %s passe au niveau %d ! 🎉\n", player.Nom, player.Niveau)

	// Bonus de stats à chaque montée de niveau
	player.HP += 10
	player.MaxHP += 10
	player.Force += 2
	player.Resistance += 2

	fmt.Printf("💪 Nouvelles stats : %d (+10) HP | %d (+2) Force | %d (+2) Résistance 💪\n",
		player.HP, player.Force, player.Resistance)
}
