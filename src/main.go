package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

// ==================== TACHE 1 ====================
// Structure Character avec les attributs requis
type Character struct {
	Name             string
	Class            string
	Level            int
	MaxHP            int
	CurrentHP        int
	Inventory        []string
	Gold             int
	Skills           []string
	Equipment        Equipment
	MaxInventory     int
	CurrentInventory int
	Mana             int
	MaxMana          int
	Experience       int
	MaxExperience    int
	Initiative       int
}

// Structure Equipment - TACHE 16
type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

// Structure Monster - TACHE 19
type Monster struct {
	Name       string
	MaxHP      int
	CurrentHP  int
	Attack     int
	Initiative int
}

// ==================== TACHE 1 ====================
// Fonction pour initialiser un personnage
func initCharacter(name string, class string, maxHP int) Character {
	return Character{
		Name:             name,
		Class:            class,
		Level:            1,
		MaxHP:            maxHP,
		CurrentHP:        maxHP / 2,
		Inventory:        []string{},
		Gold:             100,
		Skills:           []string{"Coup de poing"},
		Equipment:        Equipment{},
		MaxInventory:     10,
		CurrentInventory: 0,
		Mana:             50,
		MaxMana:          50,
		Experience:       0,
		MaxExperience:    100,
		Initiative:       10,
	}
}

// ==================== TACHE 2 ====================
// Fonction pour afficher les informations du personnage
func displayInfo(c Character) {
	fmt.Println("\n========== INFORMATIONS DU PERSONNAGE ==========")
	fmt.Printf("Nom: %s\n", c.Name)
	fmt.Printf("Classe: %s\n", c.Class)
	fmt.Printf("Niveau: %d\n", c.Level)
	fmt.Printf("PV: %d / %d\n", c.CurrentHP, c.MaxHP)
	fmt.Printf("Mana: %d / %d\n", c.Mana, c.MaxMana)
	fmt.Printf("Or: %d pieces\n", c.Gold)
	fmt.Printf("Experience: %d / %d\n", c.Experience, c.MaxExperience)
	fmt.Printf("Inventaire: %d / %d\n", c.CurrentInventory, c.MaxInventory)
	fmt.Println("==============================================\n")
}

// ==================== TACHE 4 & TACHE 12 ====================
// Fonction pour ajouter un item à l'inventaire
func addInventory(c *Character, item string) bool {
	if c.CurrentInventory >= c.MaxInventory {
		fmt.Printf("❌ Inventaire plein! Maximum: %d items\n", c.MaxInventory)
		return false
	}
	c.Inventory = append(c.Inventory, item)
	c.CurrentInventory++
	fmt.Printf("✅ %s ajouté à l'inventaire\n", item)
	return true
}

// Fonction pour retirer un item de l'inventaire
func removeInventory(c *Character, index int) bool {
	if index < 0 || index >= len(c.Inventory) {
		fmt.Println("❌ Item invalide")
		return false
	}
	item := c.Inventory[index]
	c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	c.CurrentInventory--
	fmt.Printf("✅ %s retiré de l'inventaire\n", item)
	return true
}

// ==================== TACHE 4 ====================
// Fonction pour utiliser une potion de vie
func takePot(c *Character) {
	for i, item := range c.Inventory {
		if item == "Potion de vie" {
			removeInventory(c, i)
			oldHP := c.CurrentHP
			c.CurrentHP = min(c.CurrentHP+50, c.MaxHP)
			fmt.Printf("💚 Vous utilisez une potion de vie! +%d PV\n", c.CurrentHP-oldHP)
			fmt.Printf("PV: %d / %d\n", c.CurrentHP, c.MaxHP)
			return
		}
	}
	fmt.Println("❌ Vous n'avez pas de potion de vie!")
}

// Fonction pour utiliser une potion de poison
func usePoisonPot(c *Character) {
	for i, item := range c.Inventory {
		if item == "Potion de poison" {
			removeInventory(c, i)
			fmt.Println("☠️  Vous activez une potion de poison!")
			fmt.Println("Vous subissez 10 dégâts par seconde pendant 3s...")

			for j := 1; j <= 3; j++ {
				c.CurrentHP -= 10
				if c.CurrentHP < 0 {
					c.CurrentHP = 0
				}
				fmt.Printf("Tour %d - PV: %d / %d\n", j, c.CurrentHP, c.MaxHP)
				time.Sleep(1 * time.Second)
			}
			return
		}
	}
	fmt.Println("❌ Vous n'avez pas de potion de poison!")
}

// ==================== TACHE 3 ====================
// Fonction pour accéder à l'inventaire
func accessInventory(c *Character) {
	if len(c.Inventory) == 0 {
		fmt.Println("❌ Inventaire vide!")
		return
	}

	for {
		fmt.Println("\n========== INVENTAIRE ==========")
		for i, item := range c.Inventory {
			fmt.Printf("%d. %s\n", i+1, item)
		}
		fmt.Println("0. Retour")
		fmt.Print("Choisissez une option: ")

		choice := getInput()

		if choice == "0" {
			break
		}

		index := parseInput(choice)
		if index < 1 || index > len(c.Inventory) {
			fmt.Println("❌ Choix invalide!")
			continue
		}

		item := c.Inventory[index-1]

		// Gestion des items utilisables
		switch item {
		case "Potion de vie":
			takePot(c)
		case "Potion de poison":
			usePoisonPot(c)
		case "Livre de Sort : Boule de Feu":
			if spellBook(c, "Boule de feu") {
				removeInventory(c, index-1)
			}
		case "Livre de Sort : Augmentation de Mana":
			c.MaxMana += 20
			fmt.Printf("✅ Mana maximum augmenté! Nouveau max: %d\n", c.MaxMana)
			removeInventory(c, index-1)
		case "Augmentation d'inventaire":
			if upgradeInventorySlot(c) {
				removeInventory(c, index-1)
			}
		case "Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier":
			equipItem(c, item)
		default:
			fmt.Printf("❌ Le %s ne peut pas être utilisé directement\n", item)
		}
	}
}

// ==================== TACHE 5-6 ====================
// Fonction pour afficher le menu principal
func displayMainMenu() {
	fmt.Println("\n========== MENU PRINCIPAL ==========")
	fmt.Println("1. Afficher les informations")
	fmt.Println("2. Accéder à l'inventaire")
	fmt.Println("3. Marchand")
	fmt.Println("4. Forgeron")
	fmt.Println("5. Entrainement")
	fmt.Println("6. Qui sont-ils?")
	fmt.Println("0. Quitter")
	fmt.Print("Choisissez une option: ")
}

// ==================== TACHE 7 ====================
// Fonction du Marchand
func merchant(c *Character) {
	for {
		fmt.Println("\n========== MARCHAND ==========")
		fmt.Println("--- Consommables ---")
		fmt.Println("1. Potion de vie (3 or)")
		fmt.Println("2. Potion de poison (6 or)")
		fmt.Println("3. Livre de Sort : Boule de Feu (25 or)")
		fmt.Println("4. Potion de Mana (10 or)")
		fmt.Println("--- Matériaux ---")
		fmt.Println("5. Fourrure de Loup (4 or)")
		fmt.Println("6. Peau de Troll (7 or)")
		fmt.Println("7. Cuir de Sanglier (3 or)")
		fmt.Println("8. Plume de Corbeau (1 or)")
		fmt.Println("--- Autres ---")
		fmt.Println("9. Augmentation d'inventaire (30 or)")
		fmt.Printf("\nOr disponible: %d\n", c.Gold)
		fmt.Println("0. Retour")
		fmt.Print("Choisissez un article: ")

		choice := getInput()
		if choice == "0" {
			break
		}

		handleMerchantChoice(c, choice)
	}
}

// Gestion des achats au marchand
func handleMerchantChoice(c *Character, choice string) {
	var itemName string
	var cost int

	switch choice {
	case "1":
		itemName = "Potion de vie"
		cost = 3
	case "2":
		itemName = "Potion de poison"
		cost = 6
	case "3":
		itemName = "Livre de Sort : Boule de Feu"
		cost = 25
	case "4":
		itemName = "Potion de Mana"
		cost = 10
	case "5":
		itemName = "Fourrure de Loup"
		cost = 4
	case "6":
		itemName = "Peau de Troll"
		cost = 7
	case "7":
		itemName = "Cuir de Sanglier"
		cost = 3
	case "8":
		itemName = "Plume de Corbeau"
		cost = 1
	case "9":
		itemName = "Augmentation d'inventaire"
		cost = 30
	default:
		fmt.Println("❌ Choix invalide")
		return
	}

	// Vérifier l'or
	if c.Gold < cost {
		fmt.Printf("❌ Vous n'avez pas assez d'or! (Besoin: %d, Vous avez: %d)\n", cost, c.Gold)
		return
	}

	// Ajouter l'item
	if addInventory(c, itemName) {
		c.Gold -= cost
		fmt.Printf("💰 Vous avez dépensé %d or. Or restant: %d\n", cost, c.Gold)
	}
}

// ==================== TACHE 10 ====================
// Fonction pour normaliser le nom (Majuscule + minuscules)
func normalizeName(name string) string {
	if len(name) == 0 {
		return ""
	}

	// Vérifier que le nom contient uniquement des lettres
	for _, r := range name {
		if !unicode.IsLetter(r) && r != ' ' {
			return ""
		}
	}

	name = strings.ToLower(name)
	runes := []rune(name)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// ==================== TACHE 11 ====================
// Fonction de création de personnage personnalisé
func characterCreation() Character {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n========== CREATION DE PERSONNAGE ==========")

	// Récupérer le nom
	var name string
	for {
		fmt.Print("Entrez votre nom (uniquement des lettres): ")
		inputName, _ := reader.ReadString('\n')
		inputName = strings.TrimSpace(inputName)
		normalized := normalizeName(inputName)

		if normalized == "" {
			fmt.Println("❌ Nom invalide! Utilisez uniquement des lettres.")
			continue
		}
		name = normalized
		break
	}

	// Récupérer la classe
	var class string
	for {
		fmt.Println("\nChoisissez votre classe:")
		fmt.Println("1. Humain (100 PV)")
		fmt.Println("2. Elfe (80 PV)")
		fmt.Println("3. Nain (120 PV)")
		fmt.Print("Classe: ")
		choice := getInput()

		switch choice {
		case "1":
			class = "Humain"
			break
		case "2":
			class = "Elfe"
			break
		case "3":
			class = "Nain"
			break
		default:
			fmt.Println("❌ Choix invalide!")
			continue
		}
		break
	}

	var maxHP int
	switch class {
	case "Humain":
		maxHP = 100
	case "Elfe":
		maxHP = 80
	case "Nain":
		maxHP = 120
	default:
		maxHP = 100
	}

	return initCharacter(name, class, maxHP)
}

// ==================== TACHE 18 ====================
// Fonction pour augmenter les slots d'inventaire
var inventoryUpgradeCount = 0

func upgradeInventorySlot(c *Character) bool {
	if inventoryUpgradeCount >= 3 {
		fmt.Println("❌ Vous ne pouvez augmenter votre inventaire que 3 fois!")
		return false
	}

	for i, item := range c.Inventory {
		if item == "Augmentation d'inventaire" {
			inventoryUpgradeCount++
			c.MaxInventory += 10
			fmt.Printf("✅ Inventaire augmenté! Nouveau maximum: %d\n", c.MaxInventory)
			removeInventory(c, i)
			return true
		}
	}
	return false
}

// ==================== TACHE 8 ====================
// Fonction pour vérifier si le joueur est mort
func isDead(c *Character) bool {
	if c.CurrentHP <= 0 {
		fmt.Println("💀 WASTED! Vous êtes mort!")
		c.CurrentHP = c.MaxHP / 2
		fmt.Printf("🔄 Vous êtes ressuscité avec %d PV\n", c.CurrentHP)
		return true
	}
	return false
}

// ==================== TACHE 10 ====================
// Fonction pour apprendre un sort
func spellBook(c *Character, spell string) bool {
	// Vérifier si le sort existe déjà
	for _, s := range c.Skills {
		if s == spell {
			fmt.Printf("❌ Vous connaissez déjà %s!\n", spell)
			return false
		}
	}

	c.Skills = append(c.Skills, spell)
	fmt.Printf("✅ Vous avez appris: %s\n", spell)
	return true
}

// ==================== TACHE 15 ====================
// Fonction du Forgeron
func blacksmith(c *Character) {
	for {
		fmt.Println("\n========== FORGERON ==========")
		fmt.Println("1. Chapeau de l'aventurier (1 PL + 1 CS, 5 or)")
		fmt.Println("2. Tunique de l'aventurier (2 FL + 1 PT, 5 or)")
		fmt.Println("3. Bottes de l'aventurier (1 FL + 1 CS, 5 or)")
		fmt.Printf("Or: %d\n", c.Gold)
		fmt.Println("0. Retour")
		fmt.Print("Choisissez un équipement: ")

		choice := getInput()
		if choice == "0" {
			break
		}

		handleBlacksmithChoice(c, choice)
	}
}

// Gestion de la fabrication
func handleBlacksmithChoice(c *Character, choice string) {
	var itemName string
	var requiredItems map[string]int
	var cost = 5

	switch choice {
	case "1":
		itemName = "Chapeau de l'aventurier"
		requiredItems = map[string]int{
			"Plume de Corbeau": 1,
			"Cuir de Sanglier": 1,
		}
	case "2":
		itemName = "Tunique de l'aventurier"
		requiredItems = map[string]int{
			"Fourrure de Loup": 2,
			"Peau de Troll":    1,
		}
	case "3":
		itemName = "Bottes de l'aventurier"
		requiredItems = map[string]int{
			"Fourrure de Loup": 1,
			"Cuir de Sanglier": 1,
		}
	default:
		fmt.Println("❌ Choix invalide!")
		return
	}

	// Vérifier l'or
	if c.Gold < cost {
		fmt.Printf("❌ Pas assez d'or! (Besoin: %d, Vous avez: %d)\n", cost, c.Gold)
		return
	}

	// Vérifier les matériaux
	for material, count := range requiredItems {
		inventoryCount := countItem(c, material)
		if inventoryCount < count {
			fmt.Printf("❌ Il vous manque %d %s(s)! (Avez: %d)\n", count-inventoryCount, material, inventoryCount)
			return
		}
	}

	// Retirer les matériaux
	for material, count := range requiredItems {
		for i := 0; i < count; i++ {
			for j, item := range c.Inventory {
				if item == material {
					removeInventory(c, j)
					break
				}
			}
		}
	}

	// Effectuer la fabrication
	c.Gold -= cost
	if addInventory(c, itemName) {
		fmt.Printf("✅ Vous avez fabriqué: %s!\n", itemName)
	}
}

// Fonction pour compter les items dans l'inventaire
func countItem(c *Character, itemName string) int {
	count := 0
	for _, item := range c.Inventory {
		if item == itemName {
			count++
		}
	}
	return count
}

// ==================== TACHE 16 ====================
// Fonction pour équiper un item
func equipItem(c *Character, itemName string) {
	switch itemName {
	case "Chapeau de l'aventurier":
		// Retirer l'ancien équipement
		if c.Equipment.Head != "" {
			addInventory(c, c.Equipment.Head)
			c.MaxHP -= getEquipmentBonus(c.Equipment.Head)
			fmt.Printf("Vous avez retiré: %s\n", c.Equipment.Head)
		}
		// Équiper le nouveau
		c.Equipment.Head = itemName
		c.MaxHP += getEquipmentBonus(itemName)
		fmt.Printf("✅ Vous avez équipé: %s (+10 PV max)\n", itemName)

	case "Tunique de l'aventurier":
		if c.Equipment.Torso != "" {
			addInventory(c, c.Equipment.Torso)
			c.MaxHP -= getEquipmentBonus(c.Equipment.Torso)
			fmt.Printf("Vous avez retiré: %s\n", c.Equipment.Torso)
		}
		c.Equipment.Torso = itemName
		c.MaxHP += getEquipmentBonus(itemName)
		fmt.Printf("✅ Vous avez équipé: %s (+25 PV max)\n", itemName)

	case "Bottes de l'aventurier":
		if c.Equipment.Feet != "" {
			addInventory(c, c.Equipment.Feet)
			c.MaxHP -= getEquipmentBonus(c.Equipment.Feet)
			fmt.Printf("Vous avez retiré: %s\n", c.Equipment.Feet)
		}
		c.Equipment.Feet = itemName
		c.MaxHP += getEquipmentBonus(itemName)
		fmt.Printf("✅ Vous avez équipé: %s (+15 PV max)\n", itemName)
	}

	// Supprimer de l'inventaire
	for i, item := range c.Inventory {
		if item == itemName {
			removeInventory(c, i)
			break
		}
	}
}

// Fonction pour obtenir le bonus d'équipement
func getEquipmentBonus(itemName string) int {
	switch itemName {
	case "Chapeau de l'aventurier":
		return 10
	case "Tunique de l'aventurier":
		return 25
	case "Bottes de l'aventurier":
		return 15
	default:
		return 0
	}
}

// ==================== TACHE 19 ====================
// Fonction pour initialiser un Gobelin d'entrainement
func initGoblin() Monster {
	return Monster{
		Name:       "Gobelin d'entrainement",
		MaxHP:      40,
		CurrentHP:  40,
		Attack:     5,
		Initiative: 5,
	}
}

// ==================== TACHE 20 ====================
// Pattern d'attaque du Gobelin
func goblinPattern(goblin *Monster, character *Character, turn int) {
	damage := goblin.Attack

	// Tous les 3 tours, inflige 200% de dégâts
	if turn%3 == 0 && turn != 0 {
		damage = goblin.Attack * 2
	}

	character.CurrentHP -= damage
	fmt.Printf("⚔️  %s inflige à %s %d dégâts\n", goblin.Name, character.Name, damage)
	fmt.Printf("PV de %s: %d / %d\n", character.Name, max(0, character.CurrentHP), character.MaxHP)
}

// ==================== TACHE 21 ====================
// Fonction pour le tour du joueur
func characterTurn(c *Character, goblin *Monster, reader *bufio.Reader) bool {
	fmt.Println("\n========== VOTRE TOUR ==========")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Println("3. Sorts")
	fmt.Print("Choisissez une action: ")

	choice := getInput()

	switch choice {
	case "1":
		basicAttack(c, goblin)
	case "2":
		useCombatInventory(c)
	case "3":
		useCombatSpell(c, goblin)
	default:
		fmt.Println("❌ Action invalide! Vous perdez votre tour.")
	}

	return true
}

// Fonction d'attaque basique
func basicAttack(c *Character, goblin *Monster) {
	damage := 5
	goblin.CurrentHP -= damage
	fmt.Printf("💥 %s utilise Attaque basique! %d dégâts\n", c.Name, damage)
	fmt.Printf("PV de %s: %d / %d\n", goblin.Name, max(0, goblin.CurrentHP), goblin.MaxHP)
}

// Fonction pour utiliser un item en combat
func useCombatInventory(c *Character) {
	if len(c.Inventory) == 0 {
		fmt.Println("❌ Inventaire vide!")
		return
	}

	fmt.Println("\n--- Inventaire ---")
	for i, item := range c.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("0. Retour")
	fmt.Print("Choisissez un item: ")

	choice := getInput()
	if choice == "0" {
		fmt.Println("Vous n'utilisez rien.")
		return
	}

	index := parseInput(choice)
	if index < 1 || index > len(c.Inventory) {
		fmt.Println("❌ Item invalide!")
		return
	}

	item := c.Inventory[index-1]
	if item == "Potion de vie" {
		oldHP := c.CurrentHP
		c.CurrentHP = min(c.CurrentHP+50, c.MaxHP)
		removeInventory(c, index-1)
		fmt.Printf("💚 Vous utilisez %s! +%d PV\n", item, c.CurrentHP-oldHP)
	} else if item == "Potion de Mana" {
		c.Mana = min(c.Mana+30, c.MaxMana)
		removeInventory(c, index-1)
		fmt.Printf("💙 Vous utilisez Potion de Mana! Mana: %d / %d\n", c.Mana, c.MaxMana)
	} else {
		fmt.Printf("❌ %s n'est pas utilisable en combat!\n", item)
	}
}

// Fonction pour utiliser un sort en combat
func useCombatSpell(c *Character, goblin *Monster) {
	if len(c.Skills) == 0 {
		fmt.Println("❌ Vous n'avez pas de sorts!")
		return
	}

	fmt.Println("\n--- Sorts ---")
	for i, spell := range c.Skills {
		var manaCost int
		var damage int

		switch spell {
		case "Coup de poing":
			manaCost = 5
			damage = 8
		case "Boule de feu":
			manaCost = 15
			damage = 18
		default:
			manaCost = 0
			damage = 0
		}

		fmt.Printf("%d. %s (%d mana) - %d dégâts\n", i+1, spell, manaCost, damage)
	}
	fmt.Println("0. Retour")
	fmt.Print("Choisissez un sort: ")

	choice := getInput()
	if choice == "0" {
		return
	}

	index := parseInput(choice)
	if index < 1 || index > len(c.Skills) {
		fmt.Println("❌ Sort invalide!")
		return
	}

	spell := c.Skills[index-1]
	castSpell(c, goblin, spell)
}

// Fonction pour lancer un sort
func castSpell(c *Character, goblin *Monster, spell string) {
	var manaCost int
	var damage int

	switch spell {
	case "Coup de poing":
		manaCost = 5
		damage = 8
	case "Boule de feu":
		manaCost = 15
		damage = 18
	default:
		fmt.Println("❌ Sort inconnu!")
		return
	}

	if c.Mana < manaCost {
		fmt.Printf("❌ Pas assez de mana! (Besoin: %d, Avez: %d)\n", manaCost, c.Mana)
		return
	}

	c.Mana -= manaCost
	goblin.CurrentHP -= damage
	fmt.Printf("✨ %s lance %s! %d dégâts\n", c.Name, spell, damage)
	fmt.Printf("PV de %s: %d / %d\n", goblin.Name, max(0, goblin.CurrentHP), goblin.MaxHP)
	fmt.Printf("Mana restant: %d / %d\n", c.Mana, c.MaxMana)
}

// ==================== TACHE 22 ====================
// Fonction pour le combat d'entrainement
func trainingFight(c *Character) {
	goblin := initGoblin()
	turn := 1

	fmt.Println("\n========== COMBAT D'ENTRAINEMENT ==========")
	fmt.Printf("Affrontement: %s vs %s\n", c.Name, goblin.Name)
	fmt.Println("=========================================\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		// Déterminer qui commence (initiative)
		var playerFirst bool

		// MISSION 1: Ajouter initiative
		if turn == 1 {
			if c.Initiative >= goblin.Initiative {
				playerFirst = true
				fmt.Println("✅ Vous commencez!")
			} else {
				playerFirst = false
				fmt.Println("❌ Le gobelin commence!")
			}
		}

		fmt.Printf("\n===== TOUR %d =====\n", turn)
		fmt.Printf("%s PV: %d / %d\n", c.Name, max(0, c.CurrentHP), c.MaxHP)
		fmt.Printf("%s PV: %d / %d\n\n", goblin.Name, max(0, goblin.CurrentHP), goblin.MaxHP)

		// Première attaque
		if playerFirst {
			characterTurn(c, &goblin, reader)
			if goblin.CurrentHP <= 0 {
				fmt.Printf("\n🎉 Victoire! %s est vaincu!\n", goblin.Name)
				// MISSION 2: Expérience
				gainExperience(c, 50)
				return
			}

			goblinPattern(&goblin, c, turn)
			if c.CurrentHP <= 0 {
				isDead(c)
				return
			}
		} else {
			goblinPattern(&goblin, c, turn)
			if c.CurrentHP <= 0 {
				isDead(c)
				return
			}

			characterTurn(c, &goblin, reader)
			if goblin.CurrentHP <= 0 {
				fmt.Printf("\n🎉 Victoire! %s est vaincu!\n", goblin.Name)
				gainExperience(c, 50)
				return
			}
		}

		turn++
	}
}

// ==================== MISSION 2 ====================
// Fonction pour gagner de l'expérience (mission 2)
func gainExperience(c *Character, exp int) {
	c.Experience += exp
	fmt.Printf("⭐ Vous gagnez %d points d'expérience\n", exp)

	// Vérifier le levelup
	for c.Experience >= c.MaxExperience {
		c.Experience -= c.MaxExperience
		c.Level++
		c.MaxExperience = int(float64(c.MaxExperience) * 1.1)
		c.MaxHP += 10
		c.MaxMana += 10
		fmt.Printf("🎊 LEVELUP! Vous êtes niveau %d!\n", c.Level)
		fmt.Printf("PV Max +10, Mana Max +10\n")
	}
}

// ==================== MISSION 6 ====================
// Fonction pour afficher qui sont les artistes
func whoAreThey() {
	fmt.Println("\n========== QUI SONT-ILS? ==========")
	fmt.Println("Partie 2 - ECONOMIA:")
	fmt.Println("✨ Leonardo da Vinci ✨")
	fmt.Println("L'artiste du commerce et de la création")
	fmt.Println()
	fmt.Println("Partie 3 - COMBAT:")
	fmt.Println("⚔️  Sun Tzu ⚔️")
	fmt.Println("L'artiste du combat et de la stratégie")
	fmt.Println("====================================\n")
}

// ==================== FONCTIONS UTILITAIRES ====================
func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func parseInput(s string) int {
	var choice int
	fmt.Sscanf(s, "%d", &choice)
	return choice
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ==================== MAIN ====================
func main() {
	// TACHE 11: Création personnalisée du personnage
	var player Character

	fmt.Println("Bienvenue dans le PROJET RED!")
	fmt.Println("Créez votre personnage ou utilisez le personnage par défaut?")
	fmt.Println("1. Créer mon personnage")
	fmt.Println("2. Personnage par défaut (Alan, Elfe)")
	fmt.Print("Choisissez: ")

	choice := getInput()

	if choice == "1" {
		player = characterCreation()
	} else {
		player = initCharacter("Alan", "Elfe", 80)
		// Ajouter des potions de départ
		addInventory(&player, "Potion de vie")
		addInventory(&player, "Potion de vie")
		addInventory(&player, "Potion de vie")
	}

	// Ajouter de l'initiative au personnage (MISSION 1)
	player.Initiative = 10

	// Boucle principale
	for {
		displayMainMenu()
		choice := getInput()

		switch choice {
		case "1":
			displayInfo(player)
		case "2":
			accessInventory(&player)
		case "3":
			merchant(&player)
		case "4":
			blacksmith(&player)
		case "5":
			trainingFight(&player)
		case "6":
			whoAreThey()
		case "0":
			fmt.Println("Merci d'avoir joué! Au revoir!")
			return
		default:
			fmt.Println("❌ Choix invalide!")
		}
	}
}
