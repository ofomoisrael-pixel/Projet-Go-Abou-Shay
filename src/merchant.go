package main

import "fmt"

func merchantMenu(player *Character) {

	choice := -1

	for choice != 0 {
		fmt.Println("=== MARCHAND ===")
		fmt.Println("Gold :", player.Gold)
		fmt.Println("1. Potion de vie (3)")
		fmt.Println("2. Potion poison (6)")
		fmt.Println("3. Livre de sort (25)")
		fmt.Println("4. Fourrure (4)")
		fmt.Println("5. Peau de Troll (7)")
		fmt.Println("6. Cuir (3)")
		fmt.Println("7. Plume (1)")
		fmt.Println("8. Upgrade inventaire (30)")
		fmt.Println("0. Retour")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			buyItem(player, "Potion", 3)
			p1.Gold = p1.Gold - 3
		case 2:
			buyItem(player, "Poison", 6)
			p1.Gold = p1.Gold - 6
		case 3:
			buyItem(player, "Livre", 25)
			p1.Gold = p1.Gold - 25
		case 4:
			buyItem(player, "Fourrure", 4)
			p1.Gold = p1.Gold - 4
		case 5:
			buyItem(player, "Peau", 7)
			p1.Gold = p1.Gold - 7
		case 6:
			buyItem(player, "Cuir", 3)
			p1.Gold = p1.Gold - 3
		case 7:
			buyItem(player, "Plume", 1)
			p1.Gold = p1.Gold - 1
		case 8:
			upgradeInventorySlot(player)
		}
	}
}
