package main

import "fmt"

func hasEnoughGold(player *Character, price int) bool {
	return player.Gold >= price
}

func removeGold(player *Character, amount int) {
	player.Gold -= amount
}

func addGold(player *Character, amount int) {
	player.Gold += amount
}

// Achat générique
func buyItem(player *Character, item string, price int) {

	if !hasEnoughGold(player, price) {
		fmt.Println("Pas assez d'argent")
		return
	}

	if isInventoryFull(player) {
		fmt.Println("Inventaire plein")
		return
	}

	removeGold(player, price)
	addInventory(player, item)

	fmt.Println(item, "acheté !")
}

func upgradeInventorySlot(player *Character) {

	if player.UpgradeCount >= 3 {
		fmt.Println("Limite atteinte")
		return
	}

	if player.Gold < 30 {
		fmt.Println("Pas assez d'argent")
		return
	}

	player.Gold -= 30
	player.InventoryMax += 10
	player.UpgradeCount++

	fmt.Println("Inventaire amélioré !")
}
