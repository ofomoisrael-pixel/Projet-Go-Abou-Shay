# Changelog - PROJET RED

Toutes les modifications notables du Projet RED seront documentées dans ce fichier.

Le format est basé sur [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [1.0.0] - 2026-04-02

### ✨ Ajouté

#### Partie 1 - Création du Personnage & Bases
- [x] Création de la structure Character avec tous les attributs
- [x] Fonction initCharacter pour initialiser un personnage
- [x] Fonction displayInfo pour afficher les informations
- [x] Système d'inventaire avec gestion des slots
- [x] Fonction accessInventory pour consulter l'inventaire
- [x] Potions de vie (restaurent 50 PV)
- [x] Menu principal avec 7 options
- [x] Marchand avec 9 articles à acheter
- [x] Potion de poison (10 dégâts par seconde pendant 3s)
- [x] Système de sorts (Coup de poing de base)
- [x] Fonction spellBook pour apprendre des sorts
- [x] Création personnalisée de personnage
  - [x] Validation du nom (lettres uniquement)
  - [x] Normalisation du nom (1ère majuscule, reste minuscules)
  - [x] Choix de classe (Humain/Elfe/Nain)
  - [x] PV max selon la classe
- [x] Limite d'inventaire (10 slots par défaut)

#### Partie 2 - Économie & Fabrication
- [x] Système de monnaie (100 or de départ)
- [x] Prix pour tous les articles du marchand
- [x] Vérification des fonds avant achat
- [x] Interface du forgeron
- [x] Fabrication de 3 équipements
- [x] Système de recettes avec matériaux
- [x] Structure Equipment pour les équipements
- [x] Équipement d'items avec bonus PV
- [x] Remplacement automatique d'équipements
- [x] Ajout de bonus PV :
  - [x] Chapeau : +10 PV
  - [x] Tunique : +25 PV
  - [x] Bottes : +15 PV
- [x] Fonction upgradeInventorySlot
  - [x] +10 slots par utilisation
  - [x] Limite de 3 augmentations

#### Partie 3 - Combat
- [x] Structure Monster pour les créatures
- [x] Fonction initGoblin pour créer un gobelin
- [x] Pattern d'IA du gobelin
  - [x] Dégâts 100% normal
  - [x] Dégâts 200% tous les 3 tours
- [x] Système de combat tour par tour
- [x] Fonction characterTurn pour le tour du joueur
- [x] 3 actions en combat :
  - [x] Attaquer (5 dégâts)
  - [x] Inventaire (utiliser items)
  - [x] Sorts (lancer un sort)
- [x] Fonction trainingFight pour un combat complet
- [x] Système de mort et résurrection (isDead)
  - [x] Résurrection avec 50% PV max

#### Missions Bonus
- [x] MISSION 1 : Initiative
  - [x] Initiative au personnage et monstre
  - [x] Détermination du joueur qui commence
- [x] MISSION 2 : Expérience
  - [x] Gain d'XP après combat (50 points)
  - [x] Système de montée en niveau
  - [x] Augmentation des ressources (+10 PV, +10 Mana)
  - [x] Augmentation progressive du XP requis
- [x] MISSION 3 : Combat Magique
  - [x] Utilisation de sorts en combat
  - [x] Coup de poing (8 dégâts)
  - [x] Boule de feu (18 dégâts apprendre)
  - [x] Livre de sort marchand
- [x] MISSION 4 : Ressource Mana
  - [x] Système de mana complète
  - [x] Coûts en mana pour les sorts
  - [x] Potion de Mana (restaure 30)
  - [x] Vérification mana avant sort
- [x] MISSION 5 : Amélioration du jeu
  - [x] Messages détaillés avec emojis
  - [x] Gestion d'erreurs complète
  - [x] Feedback utilisateur constant
  - [x] Interface utilisateur fluide
- [x] MISSION 6 : Qui sont-ils?
  - [x] Leonardo da Vinci (Économie)
  - [x] Sun Tzu (Combat)

### 📊 Contenu du Jeu

#### Items Marchand (9 articles)
- Potion de vie (3 or)
- Potion de poison (6 or)
- Livre de Sort : Boule de Feu (25 or)
- Potion de Mana (10 or)
- Fourrure de Loup (4 or)
- Peau de Troll (7 or)
- Cuir de Sanglier (3 or)
- Plume de Corbeau (1 or)
- Augmentation d'inventaire (30 or)

#### Équipements Forgeron (3 articles)
- Chapeau de l'aventurier (1 Plume + 1 Cuir, 5 or, +10 PV)
- Tunique de l'aventurier (2 Fourrure + 1 Peau, 5 or, +25 PV)
- Bottes de l'aventurier (1 Fourrure + 1 Cuir, 5 or, +15 PV)

#### Sorts (2 types)
- Coup de poing (8 dégâts, 5 mana) - De base
- Boule de feu (18 dégâts, 15 mana) - À apprendre

#### Classes (3 types)
- Humain (100 PV)
- Elfe (80 PV)
- Nain (120 PV)

### 📚 Documentation
- [x] README.md complet
- [x] Guide pratique d'utilisation
- [x] Documentation technique
- [x] Gestion de projet (ce fichier)
- [x] Fichiers de test (test.sh, test.bat)
- [x] Makefile pour faciliter la compilation

### 🛠️ Infrastructure
- [x] Code Go compilable
- [x] Exécutable généré (bin/projet-red.exe)
- [x] go.mod automatique
- [x] Structure de projet selon consignes
  - [x] Dossier src/ avec main.go
  - [x] Dossier docs/ avec documentations
  - [x] README.md à la racine

### 📋 Caractéristiques Techniques
- **Langage** : Go 1.16+
- **Architecture** : Programmation procédurale
- **Pattern** : Switch case + boucles infinies
- **Gestion d'état** : Pointeurs pour modification globale
- **I/O** : bufio pour entrée utilisateur
- **Délais** : time.Sleep pour animations

### 🎮 Mécanique du Jeu
- ✅ Économie fonctionnelle
- ✅ Progression claire (niveaux, équipements)
- ✅ Combat équilibré
- ✅ Limite d'inventaire réaliste
- ✅ Système de recettes
- ✅ Initiative au combat
- ✅ Expérience et montée en niveau

---

## [0.1.0] - Archivé

Version préalable aux implémentations complètes.

---

## 📌 Notes de Version

### Statut
- Production Ready ✅
- Toutes les constantes sont à titre indicatif
- Équilibre du jeu : Choix de design libre

### Connu Fonctionnement
- ✅ Création de personnage
- ✅ Gestion d'inventaire
- ✅ Système d'économie
- ✅ Fabrication d'équipement
- ✅ Combat tour par tour
- ✅ Expérience et montée de niveau

### Par Défaut du Jeu
- Aucun bug connu
- Toutes les fonctionnalités testées
- Code valide et compilable

### Limitations Volontaires
- Un seul monstre (Gobelin) - à expendre
- Un monde simple - peut être complexifié
- Pas de persistance - peut ajouter JSON
- Pas d'interface graphique - version CLI

---

## 🔄 Versionning

Format : `[MAJEURE].[MINEURE].[CORRECTIF]`

- **MAJEURE** : Changements incompatibles
- **MINEURE** : Features ajoutées
- **CORRECTIF** : Bug fixes

---

## 👨‍💻 Auteurs

- Créé par : Alan PHILIPIERT
- Équipe Ymmersion
- Projet RED © 2026

---

## 📄 Licence

Propriété du campus. À usage éducatif uniquement.

---

**Dernière mise à jour** : 2 Avril 2026
