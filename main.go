package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Déclaration de la structure Player
type Player struct {
	Nom    string `yaml:"nom"`
	Pseudo string `yaml:"pseudo"`
	Age    int    `yaml:"age"`
	Health int    `yaml:"health"`
	Mana   int    `yaml:"mana"`
}

// Méthode save pour sauvegarder les données du joueur dans un fichier YAML
func (p *Player) save() error {
	// Vérifie si le nom est vide
	if p.Nom == "" {
		return fmt.Errorf("le nom du joueur ne peut pas être vide")
	}

	// Conversion de la structure Player en YAML
	data, err := yaml.Marshal(p)
	if err != nil {
		return fmt.Errorf("erreur lors de la conversion en YAML : %v", err)
	}

	// Nom du fichier basé sur le champ Nom du joueur
	filename := p.Nom + ".yml"

	// Écriture des données YAML dans le fichier
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture dans le fichier : %v", err)
	}

	return nil
}

func main() {
	// Création d'un exemple de joueur
	player := Player{
		Nom:    "Dupont",
		Pseudo: "Dupu",
		Age:    25,
		Health: 100,
		Mana:   50,
	}

	// Sauvegarde des données du joueur dans un fichier
	if err := player.save(); err != nil {
		log.Fatalf("Erreur de sauvegarde : %v", err)
	} else {
		fmt.Println("Données du joueur sauvegardées avec succès.")
	}
}
