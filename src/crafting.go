package main

import "fmt"

func hasMaterials(player *Character, recipe map[string]int) bool {

	for item, qty := range recipe {
		if countItem(player, item) < qty {
			return false
		}
	}
	return true
}

func removeMaterials(player *Character, recipe map[string]int) {

	for item, qty := range recipe {
		for i := 0; i < qty; i++ {
			removeInventory(player, item)
		}
	}
}

func craftItem(player *Character, item string) {

	recipe := recipes[item]

	if !hasMaterials(player, recipe) {
		fmt.Println("Pas assez de ressources")
		return
	}

	if player.Gold < 5 {
		fmt.Println("Pas assez d'argent")
		return
	}

	removeMaterials(player, recipe)
	player.Gold -= 5

	addInventory(player, item)

	fmt.Println(item, "fabriqué !")
}

// Menu forgeron
func forgeMenu(player *Character) {

	choice := -1

	for choice != 0 {
		fmt.Println("=== FORGERON ===")
		fmt.Println("1. Chapeau")
		fmt.Println("2. Tunique")
		fmt.Println("3. Bottes")
		fmt.Println("0. Retour")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			craftItem(player, "Chapeau")
		case 2:
			craftItem(player, "Tunique")
		case 3:
			craftItem(player, "Bottes")
		}
	}
}

func gainXP(amount int) {
	p1.XP += amount

	if p1.XP >= 100 {
		p1.Level++
		p1.XP = 0
		p1.MaxHP += 20
		p1.CurrentHP = p1.MaxHP

		fmt.Println("🎉 LEVEL UP ! Niveau :", p1.Level)
	}
}
