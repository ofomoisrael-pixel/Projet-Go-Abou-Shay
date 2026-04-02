package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== BIENVENUE SUR PROJET RED ===")
	characterCreation()

	for {
		fmt.Println("\n--- MENU PRINCIPAL ---")
		fmt.Println("1. Informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Aller voir le Marchand")
		fmt.Println("4. Passer au combat")
		fmt.Println("5. Quitter ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			displayInfo()
		case 2:
			accessInventory()
		case 3:
			merchantMenu(&p1)
		case 4:
			startCombat()
		case 5:
			fmt.Println("Au revoir !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
