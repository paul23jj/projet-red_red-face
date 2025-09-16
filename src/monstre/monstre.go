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
}

func (m *Monstre) EnnemiAttaque(Monstre *Monstre, Personnage *class.Personnage) {
    panic("unimplemented")
}

// Fonction qui crée un monstre aléatoire
func GenererMonstre() Monstre {
    // Différents types de monstres possibles
    monstres := []Monstre{
        {"La municipale", 30, 3, 3, 3},
        {"La nationale", 50, 5, 5, 5},
        {"La bac", 70, 7, 7, 7},
        {"Le crs", 100, 10, 10, 10},
        {"Le big show", 200, 20, 20, 20},
    }

    // Choisir un monstre au hasard
    rand.Seed(time.Now().UnixNano())
    choix := rand.Intn(len(monstres))

    return monstres[choix]
}

func Monster() {
    // Exemple : générer un monstre au hasard
    ennemi := GenererMonstre()
    fmt.Printf(" Un %s apparaît ! \n HP: %d,\n Force: %d,\n Défense: %d\n",
        ennemi.Nom, ennemi.HP, ennemi.Force, ennemi.Defense)
}
