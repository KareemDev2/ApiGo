package main

import "fmt"

// Déclaration de la structure Player
type Player struct {
	Nom    string
	Pseudo string
	Age    int
	Health int
	Mana   int
}

func main() {
	// Déclaration et initialisation de la map "players"
	players := map[string]Player{
		"Dupont": {Nom: "Tommy", Pseudo: "Tom", Age: 25, Health: 100, Mana: 50},
		"Martin": {Nom: "Jerry", Pseudo: "Jee", Age: 30, Health: 80, Mana: 60},
	}

	// Affichage des informations des joueurs
	for nom, player := range players {
		fmt.Printf("Nom: %s, Player: %+v\n", nom, player)
	}
}
