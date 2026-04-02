package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Character struct {
	Name           string
	Class          string
	Level          int
	MaxHP          int
	CurrentHP      int
	Inventory      []string
	InventoryLimit int
	Skills         []string
	Money          int
}

var p1 Character

// TACHE 11 : Création interactive améliorée
func characterCreation() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrez le nom de votre héros : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	
	// Formatage : Majuscule au début, reste en minuscule
	name = strings.Title(strings.ToLower(name))

	fmt.Println("Choisissez votre classe :")
	fmt.Println("1. Humain (100 PV)")
	fmt.Println("2. Elfe (80 PV)")
	fmt.Println("3. Nain (120 PV)")
	
	var choice int
	fmt.Scan(&choice)

	maxHP := 100
	class := "Humain"

	switch choice {
	case 2:
		class = "Elfe"
		maxHP = 80
	case 3:
		class = "Nain"
		maxHP = 120
	}

	// TACHE 1 & 2 : Initialisation
	p1 = Character{
		Name:           name,
		Class:          class,
		Level:          1,
		MaxHP:          maxHP,
		CurrentHP:      maxHP / 2, // 50% des PV max
		Inventory:      []string{"Potion", "Potion", "Potion"},
		InventoryLimit: 10,
		Skills:         []string{"Coup de Poing"},
		Money:          100,
	}
	fmt.Printf("\nBienvenue, %s l'%s !\n", p1.Name, p1.Class)
}