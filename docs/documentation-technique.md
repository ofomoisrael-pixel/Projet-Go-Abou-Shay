# Documentation Technique - PROJET RED

## 📚 Table des matières

1. [Architecture du Code](#architecture-du-code)
2. [Structures de Données](#structures-de-données)
3. [Fonctions Principales](#fonctions-principales)
4. [Flux d'Exécution](#flux-dexécution)
5. [Patterns Utilisés](#patterns-utilisés)
6. [Conventions de Code](#conventions-de-code)
7. [Guide de Débogage](#guide-de-débogage)

---

## Architecture du Code

### Hiérarchie Fonctionnelle

```
main()
├── Création personnage
│   ├── characterCreation()
│   │   ├── normalizeName()
│   └── initCharacter()
├── Boucle principale
│   ├── displayMainMenu()
│   ├── Choix utilisateur
│   ├── displayInfo()
│   ├── accessInventory()
│   │   ├── addInventory() / removeInventory()
│   │   └── Utilisation d'items
│   ├── merchant()
│   │   └── handleMerchantChoice()
│   ├── blacksmith()
│   │   └── handleBlacksmithChoice()
│   ├── trainingFight()
│   │   ├── characterTurn()
│   │   │   ├── basicAttack()
│   │   │   ├── useCombatInventory()
│   │   │   └── useCombatSpell()
│   │   │       └── castSpell()
│   │   └── goblinPattern()
│   ├── whoAreThey()
│   └── Quitter
```

---

## Structures de Données

### `Character`

```go
type Character struct {
    Name            string      // Nom du personnage
    Class           string      // Classe (Humain/Elfe/Nain)
    Level           int         // Niveau actuel
    MaxHP           int         // PV maximum
    CurrentHP       int         // PV actuels
    Inventory       []string    // Liste des items
    Gold            int         // Monnaie
    Skills          []string    // Liste des sorts
    Equipment       Equipment   // Équipements équipés
    MaxInventory    int         // Slots maximum
    CurrentInventory int        // Slots utilisés
    Mana            int         // Mana actuel
    MaxMana         int         // Mana maximum
    Experience      int         // XP actuelle
    MaxExperience   int         // XP pour level up
    Initiative      int         // Initiative au combat
}
```

### `Equipment`

```go
type Equipment struct {
    Head  string   // Équipement de tête
    Torso string   // Équipement torse
    Feet  string   // Équipement pieds
}
```

### `Monster`

```go
type Monster struct {
    Name       string  // Nom du monstre
    MaxHP      int     // PV maximum
    CurrentHP  int     // PV actuels
    Attack     int     // Points d'attaque
    Initiative int     // Initiative
}
```

---

## Fonctions Principales

### 1. Gestion du Personnage

#### `initCharacter(name, class, maxHP string) Character`
- **Description** : Initialise un nouveau personnage
- **Paramètres** :
  - `name` : Nom du personnage
  - `class` : Classe (Humain/Elfe/Nain)
  - `maxHP` : Points de vie maximum
- **Retour** : Nouvelle instance de `Character`
- **Initialisation par défaut** :
  - CurrentHP = MaxHP / 2
  - Gold = 100
  - Skills = ["Coup de poing"]
  - MaxInventory = 10

#### `characterCreation() Character`
- **Description** : Crée un personnage via saisie utilisateur
- **Étapes** :
  1. Récupère le nom (validation lettres uniquement)
  2. Normalise le nom
  3. Demande la classe
  4. Détermine MaxHP selon la classe
  5. Appelle `initCharacter()`

#### `normalizeName(name string) string`
- **Description** : Normalise un nom
- **Règles** :
  - Lettres uniquement (a-z, A-Z)
  - Première lettre majuscule
  - Reste en minuscules
  - Retourne "" si caractères invalides

### 2. Affichage

#### `displayInfo(c Character)`
- Affiche toutes les stats du personnage
- Format tabulaire avec emojis
- Informations affichées :
  - Nom, Classe, Niveau
  - PV / Max PV
  - Mana / Max Mana
  - Or disponible
  - Expérience / Max Expérience
  - Inventaire (slots)

#### `displayMainMenu()`
- Affiche le menu principal
- 7 options numérotées
- Pas de validation (fait à l'appel)

### 3. Gestion d'Inventaire

#### `addInventory(c *Character, item string) bool`
- **Paramètres** :
  - `c` : Pointeur sur le caractère
  - `item` : Nom de l'item à ajouter
- **Retour** : `true` si succès, `false` si plein
- **Logique** :
  ```go
  if CurrentInventory >= MaxInventory {
      return false
  }
  Inventory.append(item)
  CurrentInventory++
  return true
  ```

#### `removeInventory(c *Character, index int) bool`
- Supprime un item à l'index spécifié
- Met à jour `CurrentInventory`
- Retour : `true` si succès

#### `countItem(c *Character, itemName string) int`
- Compte les occurrences d'un item
- Utilisé pour vérifier matériaux

### 4. Système de Potions

#### `takePot(c *Character)`
- Cherche "Potion de vie" dans l'inventaire
- Restaure 50 PV (max = MaxHP)
- Supprime la potion
- Affiche le résultat

#### `usePoisonPot(c *Character)`
- Lance la boucle de poison :
  ```
  Pour i = 1 à 3:
      CurrentHP -= 10
      Afficher PV
      Sleep(1 seconde)
  ```

### 5. Système de Sorts

#### `spellBook(c *Character, spell string) bool`
- Ajoute un sort à `Skills`
- Vérifie que le sort n'existe pas
- Retour : `true` si nouveau, `false` si existe

#### `castSpell(c *Character, goblin *Monster, spell string)`
- Vérifie le mana suffisant
- Réduit le mana
- Inflige les dégâts
- Affiche le résultat

**Coûts et Dégâts** :
- Coup de poing : 5 mana, 8 dégâts
- Boule de feu : 15 mana, 18 dégâts

### 6. Système de Marchand

#### `merchant(c *Character)`
- Boucle infinie jusqu'à choix 0
- Affiche la liste des articles
- Appelle `handleMerchantChoice()` pour chaque sélection

#### `handleMerchantChoice(c *Character, choice string)`
- Switch case sur 9 choix d'items
- Détermine nom, coût de l'item
- Vérifie l'or disponible
- Appelle `addInventory()` si ok
- Déduit le coût

### 7. Système de Forgeron

#### `blacksmith(c *Character)`
- Menu interactif des 3 équipements
- Boucle jusqu'à choix 0

#### `handleBlacksmithChoice(c *Character, choice string)`
- Détermine l'équipement à fabriquer
- Liste les matériaux requis (map)
- Vérifie :
  - Matériaux suffisants (via `countItem()`)
  - Or suffisant (5 par défaut)
- Supprime les matériaux
- Ajoute l'équipement
- Affiche messages détaillés

#### `equipItem(c *Character, itemName string)`
- Détermine l'emplacement (Head/Torso/Feet)
- Récupère ancien équipement
- Applique le bonus PV
- Place l'équipement
- Retire de l'inventaire

#### `getEquipmentBonus(itemName string) int`
- Retourne le bonus PV de l'équipement
- Chapeau : +10, Tunique : +25, Bottes : +15

### 8. Système de Combat

#### `initGoblin() Monster`
- Crée un Gobelin d'entraînement
- Stats fixes : 40 PV, 5 ATK, Initiative 5

#### `trainingFight(c *Character)`
- Boucle principale du combat
- Tour par tour
- Affiche : "TOUR X"
- Détermine initiative au tour 1
- Appelle `characterTurn()` puis `goblinPattern()`
- Vérifie la mort après chaque action
- Applique `gainExperience()` en cas de victoire

#### `characterTurn(c *Character, goblin *Monster, reader *bufio.Reader) bool`
- Menu d'action en combat :
  1. Attaquer → `basicAttack()`
  2. Inventaire → `useCombatInventory()`
  3. Sorts → `useCombatSpell()`

#### `basicAttack(c *Character, goblin *Monster)`
- Inflige 5 dégâts au gobelin
- Affiche : "[NAME] utilise Attaque basique! 5 dégâts"
- Affiche PV restants

#### `useCombatInventory(c *Character)`
- Menu des items utilisables
- Support "Potion de vie" (+50 PV)
- Support "Potion de Mana" (+30 Mana)
- Suppression automatique après utilisation

#### `useCombatSpell(c *Character, goblin *Monster)`
- Menu des sorts
- Affiche mana/dégâts pour chaque sort
- Appelle `castSpell()` si choix valide

#### `goblinPattern(goblin *Monster, character *Character, turn int)`
- Calcule les dégâts :
  ```go
  damage = goblin.Attack       // 5
  if turn % 3 == 0 && turn != 0:
      damage = goblin.Attack * 2  // 10
  ```
- Applique les dégâts
- Affiche : "[GOBLIN] inflige à [CHARACTER] X dégâts"

#### `isDead(c *Character) bool`
- Vérifie CurrentHP <= 0
- Si mort :
  - Affiche message
  - Ressuscite avec 50% PV max
  - Retour `true`

### 9. Système d'Expérience

#### `gainExperience(c *Character, exp int)`
- Ajoute l'expérience
- Boucle level up :
  ```go
  for Experience >= MaxExperience {
      Experience -= MaxExperience
      Level++
      MaxExperience = int(float64(MaxExperience) * 1.1)
      MaxHP += 10
      MaxMana += 10
      Afficher levelup
  }
  ```

### 10. Bonus

#### `upgradeInventorySlot(c *Character) bool`
- Augmente MaxInventory de +10
- Limite 3 utilisations totales
- Vérification du compteur global

#### `whoAreThey()`
- Affiche les artistes cachés :
  - Leonardo da Vinci (Partie 2)
  - Sun Tzu (Partie 3)

---

## Variables Globales

```go
var inventoryUpgradeCount = 0  // Compteur d'upgrades d'inventaire
```

**Raison** : Limite le nombre d'augmentations à 3 par partie (pas par personnage)

---

## Flux d'Exécution Principal

```
1. Démarrage (main)
   ↓
2. Affiche menu création
   ↓
3. Utilisateur choisit :
   A) Création personnalisée → characterCreation()
   B) Personnage par défaut → initCharacter("Alan", "Elfe", 80)
   ↓
4. Boucle de jeu infinie
   ├─ Afficher menu principal
   ├─ Récupérer choix utilisateur
   ├─ Exécuter action correspondante
   └─ Retour étape 4 (sauf si choix 0)
   ↓
5. Quitter
```

---

## Patterns Utilisés

### 1. Pointeurs pour Modification d'État

```go
func addInventory(c *Character, item string) bool {
    // Modification directe de c
    c.Inventory = append(c.Inventory, item)
}
```

Raison : Les slices et maps sont des pointeurs en Go,
mais la structure elle-même ne l'est pas

### 2. Switch Case pour Menus

```go
switch choice {
    case "1", "2", "3": // Faire quelque chose
    default:            // Erreur
}
```

Raison : Clarté et facilité de maintenance

### 3. Boucle Infinie pour Menus

```go
for {
    // Afficher options
    // Récupérer choix
    if choix == "0" { break }
    // Exécuter action
}
```

Raison : La boucle se termine quand l'utilisateur le souhaite

### 4. Maps pour Recettes

```go
requiredItems := map[string]int{
    "Plume de Corbeau": 1,
    "Cuir de Sanglier": 1,
}
```

Raison : Structure de données idéale pour item-quantité

---

## Conventions de Code

### Nommage

- **Fonctions** : camelCase, verbes d'action ("display", "handle")
- **Types** : PascalCase ("Character", "Equipment")
- **Variables** : camelCase ("currentHP", "goldAmount")
- **Constantes implicites** : PascalCase ("Humain", "Gobelin d'entrainement")

### Organisation

1. **Imports**
2. **Types/Structures**
3. **Fonctions (groupées par domaine)**
4. **Fonction main()**

### Commentaires

- Section : `// ==================== TÂCHE XX ====================`
- Fonction : `// Description et usage`
- Code complexe : Explications ligne par ligne

### Gestion d'Erreur

Utilise le modèle simple :
```go
if condition {
    fmt.Println("❌ Message d'erreur")
    return
}
// Continuer
```

---

## Guide de Débogage

### Problèmes Courants

#### 1. "Undefined variable"
- Vérifier la portée de la variable
- Les variables en boucle `for {...}` sont locales

#### 2. Erreur "inventory plein"
- Vérifier `CurrentInventory >= MaxInventory`
- Rappel : inventaire = slice, pas map

#### 3. Combat : dégâts incorrects
- Vérifier la division du tour : `turn % 3`
- Rappel : tour 1 n'est PAS un multiple de 3

#### 4. Mana insuffisant
- Vérifier les coûts :
  - Coup de poing : 5 mana
  - Boule de feu : 15 mana

#### 5. Pointeurs vs Valeurs
```go
// ❌ Mauvais
func modify(c Character) { c.Gold = 100 } // Pas de modification

// ✅ Bon
func modify(c *Character) { c.Gold = 100 } // Modification globale
```

### Ajout de Debug

```go
// Ajouter des logs temporaires
fmt.Printf("DEBUG: CurrentHP=%d, MaxHP=%d\n", c.CurrentHP, c.MaxHP)

// Ou utiliser "fmt.Sprintf" pour debugger sans imprimer
debug := fmt.Sprintf("Value: %d", value)
```

### Points de Rupture Logiques

Le code vérifie :
1. **Validité des entrées** (choix utilisateur)
2. **Suffisance des ressources** (or, mana, mateériaux)
3. **État du personnage** (mort, inventaire plein)
4. **Règles métier** (limite upgrades, sort unique)

---

## Améliorations Futures

### Court Terme
1. Ajouter des sons/couleurs
2. Système de sauvegarde JSON
3. Plus de monstres
4. Skill tree pour personnages

### Moyen Terme
1. Base de données pour scores
2. Système de quêtes
3. Interface graphique
4. Multijoueur basique

### Long Terme
1. Monde persistant
2. Guildes
3. Trading entre joueurs
4. Événements limités
5. Système de PvP

---

## Ressources

- [Go Documentation](https://golang.org/doc/)
- [Go Playground](https://play.golang.org/)
- [Bufio Package](https://golang.org/pkg/bufio/)
- [Time Package](https://golang.org/pkg/time/)

---

**Dernier mise à jour** : Avril 2026
**Auteur** : Équipe PROJET RED
