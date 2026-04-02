# PROJET RED - Jeu en Ligne de Commande (CLI)

## Présentation

**Projet RED** est un mini-jeu en ligne de commande écrit en **Go** qui combine plusieurs systèmes de jeu complets :
- 🎮 **Création et gestion de personnages**
- 💰 **Système d'économie et de fabrication d'équipement**
- ⚔️ **Combat tour par tour**

## Contenu du Jeu

### 1️⃣ Création du Personnage & Fonctionnalités de Base
- Créer un personnage avec une classe (Humain, Elfe, Nain)
- Afficher les informations du personnage
- Gérer l'inventaire (limité à 10 slots de base)
- Utiliser des potions (vie, poison, mana)
- Apprendre des sorts (Coup de poing, Boule de feu)
- Menu interactif pour naviguer

### 2️⃣ Système d'Économie & Fabrication
- **Marchand** : Achetez des consommables et matériaux
  - Potions (vie, poison, mana)
  - Matériaux (Fourrure de Loup, Peau de Troll, Cuir de Sanglier, Plume de Corbeau)
  - Augmentations d'inventaire
- **Forgeron** : Fabriquez des équipements
  - Chapeau de l'aventurier (+10 PV max)
  - Tunique de l'aventurier (+25 PV max)
  - Bottes de l'aventurier (+15 PV max)
- Système de monnaie (pièces d'or)

### 3️⃣ Combat Tour par Tour
- **Combat d'entraînement** contre un Gobelin
- Système de tour avec initiative
- Actions disponibles :
  - Attaquer (5 dégâts)
  - Utiliser l'inventaire
  - Lancer des sorts
- Pattern d'IA du Gobelin
- Système d'expérience et montée en niveau

## Installation

### Prérequis
- **Go 1.16+** installé sur votre système

### Étapes
1. Clonez le dépôt :
```bash
git clone https://github.com/votre-nom/projet-red_votre-nom.git
cd projet-red_votre-nom
```

2. Compilez le projet :
```bash
go build -o bin/projet-red.exe src/main.go
```

## Lancement

### Depuis la source
```bash
go run src/main.go
```

### Depuis l'exécutable
```bash
./bin/projet-red.exe
```

## Commandes du Menu Principal

| Option | Description |
|--------|-------------|
| 1 | Afficher les informations du personnage |
| 2 | Accéder à l'inventaire |
| 3 | Aller au marchand |
| 4 | Aller au forgeron |
| 5 | Démarrer un combat d'entraînement |
| 6 | Voir les artistes cachés |
| 0 | Quitter le jeu |

## Mécaniques de Jeu

### Personnage
- **PV** : Points de vie (augmentables via équipement)
- **Or** : Monnaie pour acheter des items
- **Mana** : Ressource pour lancer des sorts
- **Expérience** : Accumulez pour monter de niveau
- **Inventaire** : Limite de 10 slots (augmentable)

### Équipement
- Remplace les équipements anciens
- Ajoute des bonus aux PV max
- Complémentaire avec l'inventaire

### Sorts
- **Coup de poing** : 8 dégâts, 5 mana
- **Boule de feu** : 18 dégâts, 15 mana

### Combat
- Tour par tour
- Initiative détermine qui attaque en premier
- Dégâts varient selon l'action choisie
- Fin du combat si l'un des combattants atteint 0 PV

## Missions Bonus

| Mission | Description |
|---------|-------------|
| 1 | Système d'initiative au combat |
| 2 | Système d'expérience et montée en niveau |
| 3 | Utilisation de sorts en combat |
| 4 | Système de mana |
| 5 | Améliorations supplémentaires du jeu |
| 6 | Qui sont les artistes cachés ? |

## Structure du Projet

```
projet-red_NOM-DU-PROJET/
├── README.md           # Cette documentation
├── src/
│   └── main.go         # Code source principal
├── bin/                # Dossier contenant l'exécutable
│   └── projet-red.exe
└── docs/               # Documentation du projet
    └── gestion-projet.md
```

## Améliorations Potentielles

- [ ] Système de mondes avec plusieurs monstres
- [ ] Système de quêtes
- [ ] Sauvegarde/Chargement de personnages
- [ ] Interface graphique
- [ ] Multijoueur
- [ ] Sons et musique

## Auteurs

**Créé par** : Alan PHILIPIERT & Équipe
**Jeu en Go** : Développé pendant l'Ymmersion

## Licence

Projekt RED © 2024. Tous droits réservés.

---

**Bon jeu! 🎮**
