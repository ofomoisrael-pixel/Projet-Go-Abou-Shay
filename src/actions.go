package main

import "fmt"

// TACHE 3 : Stats
func displayInfo() {
	fmt.Println("\n--- PERSONNAGE ---")
	fmt.Printf("Nom: %s | Classe: %s | Niv: %d\n", p1.Name, p1.Class, p1.Level)
	fmt.Printf("Santé: %d/%d | Or: %d\n", p1.CurrentHP, p1.MaxHP, p1.Money)
	fmt.Printf("Sorts: %v\n", p1.Skills)
}

// TACHE 8 : Mort
func isDead() bool {
	if p1.CurrentHP <= 0 {
		fmt.Println("\n[WASTED]")
		p1.CurrentHP = p1.MaxHP / 2
		fmt.Println("Ressuscité avec 50% PV.")
		return true
	}
	return false
}

// TACHE 7 : Marchand
func merchantMenu() {
	fmt.Println("\n--- MARCHAND ---")
	fmt.Println("1. Potion de vie (Gratuit)")
	fmt.Println("2. Potion de Poison (Gratuit)")
	fmt.Println("3. Livre de Sort : Boule de Feu (Gratuit)")
	fmt.Println("0. Quitter")

	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1: addInventory("Potion")
	case 2: addInventory("Potion de Poison")
	case 3: addInventory("Livre de Sort : Boule de Feu")
	}
}