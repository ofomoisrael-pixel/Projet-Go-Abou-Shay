package main

import (
	"fmt"
	"time"
)

// TACHE 4 : Accès Inventaire
func accessInventory() {
	fmt.Println("\n--- INVENTAIRE ---")
	if len(p1.Inventory) == 0 {
		fmt.Println("Vide.")
	}
	for i, item := range p1.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("0. Retour")

	var choice int
	fmt.Scan(&choice)
	if choice > 0 && choice <= len(p1.Inventory) {
		useItem(p1.Inventory[choice-1], choice-1)
	}
}

func useItem(item string, index int) {
	switch item {
	case "Potion":
		takePot(index)
	case "Potion de Poison":
		poisonPot(index)
	case "Livre de Sort : Boule de Feu":
		spellBook(index)
	default:
		fmt.Println("Cet objet n'est pas utilisable.")
	}
}

// TACHE 5 : Potion
func takePot(index int) {
	if p1.CurrentHP >= p1.MaxHP {
		fmt.Println("PV déjà au max.")
		return
	}
	p1.CurrentHP += 50
	if p1.CurrentHP > p1.MaxHP { p1.CurrentHP = p1.MaxHP }
	removeInventory(index)
	fmt.Printf("Soigné ! PV : %d/%d\n", p1.CurrentHP, p1.MaxHP)
}

// TACHE 9 : Poison
func poisonPot(index int) {
	removeInventory(index)
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		p1.CurrentHP -= 10
		fmt.Printf("Dégâts poison... PV : %d/%d\n", p1.CurrentHP, p1.MaxHP)
		if isDead() { break }
	}
}

// TACHE 10 : Apprendre un sort
func spellBook(index int) {
	for _, s := range p1.Skills {
		if s == "Boule de Feu" {
			fmt.Println("Déjà connu !")
			return
		}
	}
	p1.Skills = append(p1.Skills, "Boule de Feu")
	removeInventory(index)
	fmt.Println("Sort appris : Boule de Feu !")
}

// Outils (TACHE 12)
func addInventory(item string) {
	if len(p1.Inventory) < p1.InventoryLimit {
		p1.Inventory = append(p1.Inventory, item)
		fmt.Printf("Ajouté : %s\n", item)
	} else {
		fmt.Println("Inventaire plein !")
	}
}

func removeInventory(index int) {
	p1.Inventory = append(p1.Inventory[:index], p1.Inventory[index+1:]...)
}