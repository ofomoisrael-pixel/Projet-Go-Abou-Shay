package main

import (
	"fmt"
)

// ⚔️ Combat classique
func startCombat() {
	monster := createMonster()

	fmt.Println("\n⚔️ COMBAT ⚔️")
	fmt.Println("Un", monster.Name, "apparaît !")

	for monster.HP > 0 && p1.CurrentHP > 0 {

		fmt.Println("\n--- TON TOUR ---")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Utiliser un sort")

		var choice int
		fmt.Scan(&choice)

		switch choice {

		case 1:
			damage := 15
			monster.HP -= damage
			fmt.Println("Tu infliges", damage, "dégâts !")

		case 2:
			useSkill(&monster)

		default:
			fmt.Println("Action invalide")
		}

		// Tour du monstre
		if monster.HP > 0 {
			monsterAttack(monster)
		}
	}

	// Fin du combat
	if p1.CurrentHP <= 0 {
		fmt.Println("💀 Tu as perdu...")
	} else {
		fmt.Println("🏆 Victoire !")
		gainXP(50)
		dropLoot()
		giveGold(20)
	}
}

/////////////////////////////////////////////////////

// 🔥 Combat Boss
func startBossCombat() {
	monster := createBoss()

	fmt.Println("\n🔥 BOSS FINAL 🔥")
	fmt.Println("Le", monster.Name, "apparaît !")

	for monster.HP > 0 && p1.CurrentHP > 0 {

		fmt.Println("\n--- TON TOUR ---")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Utiliser un sort")

		var choice int
		fmt.Scan(&choice)

		switch choice {

		case 1:
			damage := 20
			monster.HP -= damage
			fmt.Println("Tu infliges", damage, "dégâts !")

		case 2:
			useSkill(&monster)

		default:
			fmt.Println("Action invalide")
		}

		if monster.HP > 0 {
			monsterAttack(monster)
		}
	}

	if p1.CurrentHP <= 0 {
		fmt.Println("💀 Tu as perdu contre le boss...")
	} else {
		fmt.Println("🏆 Tu as vaincu le boss !!!")
		gainXP(100)
		dropLoot()
		giveGold(100)

	}
}

// 🎁 Loot après combat
func dropLoot() {
	loot := "Potion"
	fmt.Println("🎁 Tu as reçu :", loot)
	addInventory(&p1, loot)
}
