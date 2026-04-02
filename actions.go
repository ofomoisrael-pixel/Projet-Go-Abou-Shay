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
