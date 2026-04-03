package main

import "fmt"

func startCombat() {
	monster := createMonster()

	fmt.Println("\n⚔️ COMBAT COMMENCE ⚔️")
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
			fmt.Printf("Tu infliges %d dégâts !\n", damage)

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
		fmt.Println("💀 Tu as perdu...")
	} else {
		fmt.Println("🏆 Victoire !")
		gainXP(50)
		dropLoot()
	}
}
func dropLoot() {
	loot := "Potion"
	fmt.Println("🎁 Tu as reçu :", loot)
	addInventory(&p1, loot)
}

func startBossCombat() {
	monster := createBoss()
	fmt.Println("🔥 BOSS FINAL :", monster.Name)
	// même logique que startCombat
}
