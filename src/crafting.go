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
		fmt.Println("Not enough ressources")
		return
	}

	if player.Gold < 5 {
		fmt.Println("Not enough Money")
		return
	}

	removeMaterials(player, recipe)
	player.Gold -= 5

	addInventory(player, item)

	fmt.Println(item, "made !")
}

// Menu forgeron
func forgeMenu(player *Character) {

	choice := -1

	for choice != 0 {
		fmt.Println("=== BLACKSMITH ===")
		fmt.Println("1. hat")
		fmt.Println("2. coat")
		fmt.Println("3. Boots")
		fmt.Println("0. exit")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			craftItem(player, "hat")
		case 2:
			craftItem(player, "coat")
		case 3:
			craftItem(player, "boots")
		}
	}
}
