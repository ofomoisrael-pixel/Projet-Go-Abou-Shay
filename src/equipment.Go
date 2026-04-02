package main

// Structure équipement
type Equipment struct {
	Head string
	Body string
	Feet string
}

// Équiper un objet
func equipItem(player *Character, item string) {

	switch item {

	case "Chapeau":
		if player.Head != "" {
			addInventory(player, player.Head)
			removeBonus(player, player.Head)
		}
		player.Head = item

	case "Tunique":
		if player.Body != "" {
			addInventory(player, player.Body)
			removeBonus(player, player.Body)
		}
		player.Body = item

	case "Bottes":
		if player.Feet != "" {
			addInventory(player, player.Feet)
			removeBonus(player, player.Feet)
		}
		player.Feet = item
	}

	removeInventory(player, item)
	applyBonus(player, item)
}

// Bonus stats
func applyBonus(player *Character, item string) {

	switch item {
	case "Chapeau":
		player.MaxHP += 10
	case "Tunique":
		player.MaxHP += 25
	case "Bottes":
		player.MaxHP += 15
	}
}

func removeBonus(player *Character, item string) {

	switch item {
	case "Chapeau":
		player.MaxHP -= 10
	case "Tunique":
		player.MaxHP -= 25
	case "Bottes":
		player.MaxHP -= 15
	}
}
