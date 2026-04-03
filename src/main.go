package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(`
██████╗ ███████╗██████╗ 
██╔══██╗██╔════╝██╔══██╗
██████╔╝█████╗  ██║  ██║
██╔══██╗██╔══╝  ██║  ██║
██║  ██║███████╗██████╔╝
╚═╝  ╚═╝╚══════╝╚═════╝ 
`)
	characterCreation()

	for {
		fmt.Println("\n--- MENU PRINCIPAL ---")
		fmt.Println("1. Informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Aller voir le Marchand")
		fmt.Println("4. Passer au combat")
		fmt.Println("5. Boss")
		fmt.Println("6. Quitter")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			displayInfo()
		case 2:
			accessInventory()
		case 3:
			merchantMenu(&p1)
			printMerchantLogo()
		case 4:
			startCombat()
			printCombatLogo()
		case 5:
			startBossCombat()
			printBossLogo()
		case 6:
			fmt.Println("Au revoir !")
			os.Exit(0)

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
