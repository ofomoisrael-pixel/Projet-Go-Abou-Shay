# Gestion de Projet - PROJET RED

## 📋 Vue d'ensemble

Ce document détaille la réalisation du **Projet RED**, un mini-jeu en ligne de commande (CLI) en Go implémentant un système complet de gestion de personnage, d'économie et de combat.

## ✅ Tâches Complétées

### Partie 1️⃣ : Création du Personnage & Fonctionnalités de Base

#### ✔️ TACHE 1 : Création du personnage
- **Statut** : ✅ Complétée
- **Description** : Création de la structure `Character` avec tous les attributs requis
- **Attributs implémentés** :
  - `Name` (string) - Nom du personnage
  - `Class` (string) - Classe (Humain/Elfe/Nain)
  - `Level` (int) - Niveau actuel
  - `MaxHP` (int) - Points de vie maximum
  - `CurrentHP` (int) - Points de vie actuels
  - `Inventory` ([]string) - Inventaire du personnage
  - `Gold` (int) - Monnaie
  - `Skills` ([]string) - Liste de sorts
  - `Equipment` (Equipment) - Équipements équipés
  - `MaxInventory` / `CurrentInventory` (int) - Gestion des slots
  - `Mana` / `MaxMana` (int) - Système de mana
  - `Experience` / `MaxExperience` (int) - Expérience (Bonus Mission 2)
  - `Initiative` (int) - Initiative au combat (Bonus Mission 1)

#### ✔️ TACHE 2 : Affichage des informations
- **Statut** : ✅ Complétée
- **Fonction** : `displayInfo()`
- **Fonctionnalité** : Affiche toutes les informations du personnage de manière formatée

#### ✔️ TACHE 3 : Accès à l'inventaire
- **Statut** : ✅ Complétée
- **Fonction** : `accessInventory()`
- **Fonctionnalité** : Menu interactif pour voir et utiliser les items de l'inventaire

#### ✔️ TACHE 4 : Utilisation de potion de vie
- **Statut** : ✅ Complétée
- **Fonction** : `takePot()`
- **Mécanique** :
  - Consomme une potion de l'inventaire
  - Récupère 50 PV
  - Les PV ne peuvent pas dépasser le maximum
  - Affiche les PV restants

#### ✔️ TACHE 5 : Menu principal
- **Statut** : ✅ Complétée
- **Fonction** : `displayMainMenu()`
- **Options du menu** :
  1. Afficher les informations
  2. Accéder à l'inventaire
  3. Marchand
  4. Forgeron
  5. Entrainement
  6. Qui sont-ils?
  0. Quitter

#### ✔️ TACHE 6 : Liaison du menu
- **Statut** : ✅ Complétée
- **Implémentation** :
  - Menu principal dans `main()`
  - Switch case pour gérer les choix
  - Navigation entre les différent modes

#### ✔️ TACHE 7 : Interface du marchand
- **Statut** : ✅ Complétée
- **Fonction** : `merchant()` & `handleMerchantChoice()`
- **Items vendus** :
  - Potion de vie (3 or)
  - Potion de poison (6 or)
  - Livre de Sort : Boule de Feu (25 or)
  - Potion de Mana (10 or) - Bonus
  - Fourrure de Loup (4 or)
  - Peau de Troll (7 or)
  - Cuir de Sanglier (3 or)
  - Plume de Corbeau (1 or)
  - Augmentation d'inventaire (30 or)

#### ✔️ TACHE 8 : Système wasted (mort et résurrection)
- **Statut** : ✅ Complétée
- **Fonction** : `isDead()`
- **Mécanique** :
  - Vérifie si le joueur est mort (≤ 0 PV)
  - Ressuscite avec 50% des PV max
  - Affiche un message dramatique 💀

#### ✔️ TACHE 9 : Potion de poison
- **Statut** : ✅ Complétée
- **Fonction** : `usePoisonPot()`
- **Mécanique** :
  - Inflige 10 points de dégâts par seconde
  - Dure 3 secondes
  - Affiche les PV après chaque seconde
  - Utilise `time.Sleep()` pour le délai

#### ✔️ TACHE 10 : Système de sorts (Wingardium Leviosa)
- **Statut** : ✅ Complétée
- **Fonction** : `spellBook()`
- **Implémentation** :
  - Attribut `Skills` contenant les sorts
  - De base : "Coup de poing"
  - Nouveau sort : "Boule de feu" via un livre
  - Vérification : un sort ne peut être appris qu'une seule fois
  - Item marchand : "Livre de Sort : Boule de Feu"

#### ✔️ TACHE 11 : Amélioration création de personnage
- **Statut** : ✅ Complétée (2 parties)
- **Fonction** : `characterCreation()` & `normalizeName()`
- **Partie 1** : Contrôle du nom
  - Uniquement des lettres
  - Première lettre majuscule, rest minuscules
  - Validation avec regex
- **Partie 2** : Choix de classe
  - Humain : 100 PV
  - Elfe : 80 PV
  - Nain : 120 PV
  - PV actuels = 50% du max
  - Niveau 1 de base
  - Sort de base : "Coup de poing"

#### ✔️ TACHE 12 : Limite d'inventaire
- **Statut** : ✅ Complétée
- **Fonction** : `addInventory()`
- **Mécanique** :
  - Limite de 10 slots par défaut
  - Vérification avant chaque ajout
  - Message d'erreur si plein

### Partie 2️⃣ : Économie et Système de Fabrication

#### ✔️ TACHE 13 : Système d'argent
- **Statut** : ✅ Complétée
- **Attribut** : `Gold` (int)
- **Valeur initiale** : 100 pièces d'or

#### ✔️ TACHE 14 : Marchand avec coûts
- **Statut** : ✅ Complétée
- **Implémentation** :
  - Déduction de l'or à chaque achat
  - Vérification des fonds suffisants
  - Message d'erreur si argent insuffisant
  - Tous les prix implémentés

#### ✔️ TACHE 15 : Interface du forgeron
- **Statut** : ✅ Complétée (2 parties)
- **Partie 1** : Menu du forgeron
  - Fonction : `blacksmith()`
  - Menu interactif avec les 3 équipements
- **Partie 2** : Système de fabrication
  - Fonction : `handleBlacksmithChoice()`
  - Vérification des matériaux requis
  - Vérification de l'or (5 pièces)
  - Suppression des matériaux après fabrication
  - Messages d'erreur personnalisés
  
**Recettes implémentées** :
  - Chapeau : 1 Plume + 1 Cuir = 5 or
  - Tunique : 2 Fourrure + 1 Peau = 5 or
  - Bottes : 1 Fourrure + 1 Cuir = 5 or

#### ✔️ TACHE 16 : Structure d'équipement
- **Statut** : ✅ Complétée (2 parties)
- **Partie 1** : Structure Equipment
  - 3 slots : `Head`, `Torso`, `Feet`
  - Ajout à la structure Character
- **Partie 2** : Équipement des items
  - Fonction : `equipItem()`
  - Bonus aux PV max :
    - Chapeau : +10 PV
    - Tunique : +25 PV
    - Bottes : +15 PV
  - Remplacement automatique de l'ancien équipement
  - Ancien équipement retourné à l'inventaire
  - Suppression de l'inventaire lors de l'équipement

#### ✔️ TACHE 17 / 18 : Augmentation d'inventaire
- **Statut** : ✅ Complétée
- **Fonction** : `upgradeInventorySlot()`
- **Mécanique** :
  - Augmente de +10 slots
  - Maximum 3 utilisations
  - Coût : 30 pièces d'or au marchand
  - Utilisation via l'inventaire

### Partie 3️⃣ : Système de Combat

#### ✔️ TACHE 19 : Structure Monster & Gobelin
- **Statut** : ✅ Complétée (2 parties)
- **Partie 1** : Structure Monster
  - Attributs : `Name`, `MaxHP`, `CurrentHP`, `Attack`, `Initiative`
- **Partie 2** : Gobelin d'entrainement
  - Fonction : `initGoblin()`
  - Attributs :
    - Nom : "Gobelin d'entrainement"
    - MaxHP : 40
    - CurrentHP : 40
    - Attack : 5
    - Initiative : 5

#### ✔️ TACHE 20 : Pattern d'IA du Gobelin
- **Statut** : ✅ Complétée
- **Fonction** : `goblinPattern()`
- **Mécanique** :
  - 100% de ses dégâts chaque tour (5 dégâts)
  - 200% de ses dégâts tous les 3 tours (10 dégâts)
  - Affichage avec le nombre de dégâts infligés
  - Affichage des PV restants

#### ✔️ TACHE 21 : Tour du joueur en combat
- **Statut** : ✅ Complétée (2 parties)
- **Partie 1** : Interface de combat
  - Fonction : `characterTurn()`
  - Menu avec 3 choix :
    1. Attaquer (5 dégâts)
    2. Inventaire (utiliser des items)
    3. Sorts (lancer un sort)
- **Partie 2** : Actions en combat
  - `basicAttack()` : Attaque basique (5 dégâts)
  - `useCombatInventory()` : Utiliser potion en combat
  - `useCombatSpell()` : Lancer un sort

#### ✔️ TACHE 22 : Combat complet
- **Statut** : ✅ Complétée (2 parties)
- **Partie 1** : Système de tour
  - Fonction : `trainingFight()`
  - Système tour par tour
  - Affichage du tour actuel
  - Vérification de la mort après chaque action
- **Partie 2** : Menu principal
  - Ajout du choix "Entrainement"
  - Fin du combat et retour au menu si mort

### 🎯 Missions Bonus

#### ✔️ MISSION 1 : Système d'initiative
- **Statut** : ✅ Complétée
- **Implémentation** :
  - Attribut `Initiative` pour Character et Monster
  - Détermination du joueur qui commence
  - Affichage de qui commence
  - Valeur : 10 pour le joueur, 5 pour le gobelin

#### ✔️ MISSION 2 : Système d'expérience
- **Statut** : ✅ Complétée
- **Fonction** : `gainExperience()`
- **Mécanique** :
  - Gain d'expérience après combat
  - Montée en niveau automatique
  - Augmentation des PV (+10) et Mana (+10) à chaque niveau
  - Augmentation de l'expérience requise (x1.1)
  - Gestion de l'expérience excédentaire

#### ✔️ MISSION 3 : Sorts de combat
- **Statut** : ✅ Complétée
- **Fonction** : `castSpell()`
- **Sorts implémentés** :
  - Coup de poing : 8 dégâts
  - Boule de feu : 18 dégâts
- **Utilisation** : Menu "Sorts" en combat

#### ✔️ MISSION 4 : Système de mana
- **Statut** : ✅ Complétée
- **Implémentation** :
  - Attributs `Mana` et `MaxMana` dans Character
  - Coûts en mana pour les sorts :
    - Coup de poing : 5 mana
    - Boule de feu : 15 mana
  - Vérification du mana avant utilisation
  - Potion de Mana : +30 au marchand (10 or)

#### ✔️ MISSION 5 : Améliorations
- **Statut** : ✅ Implémentées
- **Améliorations** :
  - Messages formatés avec emojis pour meilleure UX
  - Système de mana complète
  - Gestion d'erreurs détaillée
  - Validation des entrées utilisateur
  - Feedback visuel constant

#### ✔️ MISSION 6 : Qui sont-ils?
- **Statut** : ✅ Complétée
- **Fonction** : `whoAreThey()`
- **Réponse cachée** :
  - Partie 2 (Économie) : **Leonardo da Vinci**
  - Partie 3 (Combat) : **Sun Tzu**

## 📊 Résumé Statistique

| Catégorie | Nombre |
|-----------|--------|
| Tâches Principales | 22 ✅ |
| Missions Bonus | 6 ✅ |
| Structures de données | 3 |
| Fonctions principales | 30+ |
| Lignes de code | 900+ |
| Items dans le jeu | 12 |
| Équipements | 3 |
| Sorts | 2 |
| Classes de personnage | 3 |

## 🔧 Fonctionnalités Clés

### Gestion de Personnage
- ✅ Création personnalisée
- ✅ Affichage d'informations
- ✅ Système d'inventaire avec limite
- ✅ Équipement avec bonus
- ✅ Niveaux et expérience
- ✅ Mana et sorts

### Économie
- ✅ Système monétaire (or)
- ✅ Marchand avec 9 items
- ✅ Forgeron avec 3 équipements
- ✅ Système de coûts
- ✅ Gestion des ressources

### Combat
- ✅ Combat tour par tour
- ✅ Système d'initiative
- ✅ 3 types d'actions
- ✅ IA simple mais efficace
- ✅ Expérience et récompenses

## 🎨 Points Forts

1. **Complétude** : Toutes les tâches implémentées 100%
2. **UX** : Interface claire avec emojis et couleurs conceptuelles
3. **Validation** : Gestion d'erreurs pour tous les cas
4. **Scalabilité** : Code organisé et facile à étendr
5. **Performance** : Compilation rapide, exécution fluide

## 📝 Notes de Développement

- Code écrit en Go avec bonnes pratiques
- Utilisation de `bufio` pour l'input utilisateur
- `time.Sleep()` pour les délais
- Structures imbriquées pour l'organisation des données
- Gestion des pointeurs pour la modification d'état

## 🚀 Points d'Extension Future

Si le projet était à continuer, on pourrait ajouter :

1. **Système de quêtes** : Missions avec récompenses
2. **Plusieurs monstres** : Monde avec différents ennemis
3. **Système de guildes** : Coopération entre joueurs
4. **Sauvegarde/Chargement** : Persistance des personnages
5. **Interface graphique** : Version GUI pour meilleure UX
6. **Multijoueur** : PvP ou coopératif
7. **Événements** : Donjons, boss spéciaux
8. **Loot** : Système de drops aléatoires

---

**Projet terminé avec succès ! 🎉**

*Dernière mise à jour : Avril 2026*
