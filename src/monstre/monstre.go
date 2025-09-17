package Monstre

import (
	class "PROJETRED/src/class"
	"fmt"
	"math/rand"
	"time"
)

// Structure du monstre
type Monstre struct {
	Nom     string
	HP      int
	Force   int
	Defense int
	Vitesse int
	XPValue int                // Valeur d'XP donnée à la victoire
	Loot    []class.Inventaire // objets que ce monstre peut drop
}

// Méthode d'attaque du monstre (à compléter plus tard)
func (m *Monstre) EnnemiAttaque(cible *class.Personnage) {
	// Exemple simple : dégâts = force - une partie de la résistance du joueur
	degats := m.Force - (cible.Resistance / 2)
	if degats < 0 {
		degats = 0
	}
	cible.HP -= degats
	fmt.Printf("%s attaque %s et inflige %d dégâts !\n", m.Nom, cible.Nom, degats)
	if cible.HP < 0 {
		cible.HP = 0
	}
}

// Fonction qui crée un monstre aléatoire
func GenererMonstre() Monstre {
	// Différents types de monstres possibles
	monstres := []Monstre{
		{"La municipale", 30, 3, 3, 3, 10, []class.Inventaire{{Name: "Pantalon de la Municipale", Quantity: 1}}},
		{"La nationale", 50, 5, 5, 5, 20, []class.Inventaire{{Name: "Holster de la BAC", Quantity: 1}}},
		{"La bac", 70, 7, 7, 7, 30, []class.Inventaire{{Name: "Gilet de la BAC", Quantity: 1}}},
		{"Le crs", 100, 10, 10, 10, 50, []class.Inventaire{{Name: "Casque de CRS", Quantity: 1}}},
		{"Le big show", 200, 20, 20, 20, 100, []class.Inventaire{{Name: "Bottes de Big Show", Quantity: 1}}},
	}

	// Choisir un monstre au hasard
	rand.Seed(time.Now().UnixNano())
	choix := rand.Intn(len(monstres))

	return monstres[choix]
}

// Fonction qui affiche un monstre généré
func Monster() Monstre {
	ennemi := GenererMonstre()
	fmt.Printf("⚔️ Un %s apparaît !\nHP: %d\nForce: %d\nDéfense: %d\n",
		ennemi.Nom, ennemi.HP, ennemi.Force, ennemi.Defense)
	return ennemi
}

// Gestion du loot (20% de chance de drop)
func (m *Monstre) DropLoot() *class.Inventaire {
	if len(m.Loot) == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())

	// 20% de chance de drop
	if rand.Intn(100) < 20 {
		idx := rand.Intn(len(m.Loot))
		return &m.Loot[idx] // retourne un vrai élément de la slice
	}

	return nil
}
