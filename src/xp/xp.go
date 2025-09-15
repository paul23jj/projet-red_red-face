package xp

import (
	"fmt"
	class "PROJETRED/src/class"
)

// GainXP ajoute de l’XP au joueur et gère la montée de niveau
func GainXP(player *class.Personnage, amount int) {
	player.Niveau += amount
	fmt.Printf("✨ %s a gagné %d points d'expérience !\n", player.Nom, amount)

	// Exemple de règle : tous les 10 XP → le joueur augmente de niveau
	if player.Niveau%10 == 0 {
		LevelUp(player)
	}
}

// LevelUp augmente les stats du joueur
func LevelUp(player *class.Personnage) {
	fmt.Printf("🎉 %s passe au niveau %d !\n", player.Nom, player.Niveau/10+1)

	// Bonus de stats à chaque montée de niveau
	player.HP += 10
	player.MaxHP += 10
	player.Force += 2
	player.Resistance += 2

	fmt.Printf("💪 Nouvelles stats : %d HP | %d Force | %d Résistance\n",
		player.HP, player.Force, player.Resistance)
}
