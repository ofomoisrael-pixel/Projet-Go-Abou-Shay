package main

import "fmt"

func startCombat() {
	monster := createGoblin()
	turn := 1 

	fmt.Println("\n⚔️ LE COMBAT COMMENCE ⚔️")
	fmt.Printf("Un %s apparaît !\n", monster.Name)

	for monster.HP > 0 && p1.CurrentHP > 0 {
		// Affichage du début de tour
		fmt.Printf("\n========== TOUR %d ==========\n", turn)

		// 1. TOUR DU JOUEUR
		fmt.Println("--- TON TOUR ---")
		fmt.Println("1. Attaquer\n2. Accéder à l'inventaire")
		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			damage := 15
			monster.HP -= damage
			fmt.Printf("Tu infliges %d dégâts au %s !\n", damage, monster.Name)
		} else {
			accessInventory()
		}

		// Vérification si le monstre meurt avant d'attaquer
		if monster.HP <= 0 { 
			fmt.Println("🏆 Victoire ! Le monstre est vaincu.")
			break 
		}

		// 2. TOUR DU MONSTRE
		fmt.Printf("\n--- TOUR DU %s ---\n", monster.Name)
		damageReceived := monster.Attack
		
		// Pattern Tâche 20 : Coup critique tous les 3 tours
		if turn%3 == 0 { 
			damageReceived *= 2
			fmt.Println("⚠️ ATTENTION : Le monstre enrage et double ses dégâts !")
		}

		p1.CurrentHP -= damageReceived
		fmt.Printf("Le %s t'inflige %d dégâts !\n", monster.Name, damageReceived)
		
		// AFFICHAGE DES PV RESTANTS APPRÈS LE TOUR
		fmt.Printf("❤️ PV RESTANTS : %d / %d\n", p1.CurrentHP, p1.MaxHP)
		
		// Vérification de la mort du joueur
		if isDead() { 
			fmt.Println("Le combat s'arrête car tu as succombé.")
			break 
		}
		
		turn++ 
	}
}