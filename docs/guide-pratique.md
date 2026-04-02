# Guide Pratique - PROJET RED

## 🎮 Guide de Jeu Complet

### 1. Démarrage du Jeu

Vous avez le choix entre :
```bash
go run src/main.go
```
ou
```bash
./bin/projet-red.exe
```

### 2. Création du Personnage

Lors du lancement, le jeu vous demande :
1. **Créer votre personnage** (Choix 1)
   - Entrez un nom (lettres uniquement)
   - Choisissez votre classe :
     - 🧝 Elfe (80 PV)
     - 👨 Humain (100 PV)
     - ⚒️ Nain (120 PV)

2. **Personnage par défaut** (Choix 2)
   - Personnage rapide : Alan, Elfe
   - Démarre avec 3 potions de vie

### 3. Menu Principal

Après la création, vous accédez au menu avec 6 options :

#### 1️⃣ Afficher les informations
Montre :
- 📊 Nom, Classe, Niveau
- ❤️ PV actuels / max
- 💙 Mana actuel / max
- 💰 Or disponible
- ⭐ Expérience
- 🎒 Inventaire (slots utilisés / max)

#### 2️⃣ Accéder à l'inventaire
Options :
- Voir tous les items
- **Potions** : Utilisables (restaurent PV/Mana)
- **Livres** : Apprennent des sorts
- **Équipements** : À équiper
- **Matériaux** : Utilisables au forgeron

**Éléments utilisables en inventaire** :
- 💚 Potion de vie → +50 PV
- 💙 Potion de Mana → +30 Mana
- 📖 Livre de Sort → Apprend un nouveau sort
- 🛡️ Équipements → Les équiper

#### 3️⃣ Marchand 🏪
**Articles disponibles** :

| Item | Coût | Type |
|------|------|------|
| Potion de vie | 3 or | Consommable |
| Potion de poison | 6 or | Consommable |
| Livre de Sort : Boule de Feu | 25 or | Apprentissage |
| Potion de Mana | 10 or | Consommable |
| Fourrure de Loup | 4 or | Matériau |
| Peau de Troll | 7 or | Matériau |
| Cuir de Sanglier | 3 or | Matériau |
| Plume de Corbeau | 1 or | Matériau |
| Augmentation d'inventaire | 30 or | Amélioration |

**Stratégie d'achat** :
- Commencez par les potions (3 or chacune)
- Economisez pour les matériaux
- Plantifiez 3 augmentations d'inventaire max

#### 4️⃣ Forgeron ⚒️
**Équipements fabriables** :

| Équipement | Matériaux | Coût | Bonus |
|------------|-----------|------|-------|
| Chapeau | 1 Plume + 1 Cuir | 5 or | +10 PV max |
| Tunique | 2 Fourrure + 1 Peau | 5 or | +25 PV max |
| Bottes | 1 Fourrure + 1 Cuir | 5 or | +15 PV max |

**Stratégie de fabrication** :
1. Achetez les matériaux au marchand
2. Rendez-vous au forgeron
3. Sélectionnez l'équipement
4. Les matériaux sont automatiquement consommés
5. Équipez l'item depuis l'inventaire

**Exemple de progression** :
```
Début: 100 or
→ Acheter 1 Plume (1 or), 1 Cuir (3 or) = 4 pages
→ Acheter 1 Plume (1 or) de plus
→ Total matériaux: 2 Plume + 1 Cuir
→ Fabriquer Chapeau (5 or) → +10 PV
→ Équiper depuis inventaire
```

#### 5️⃣ Entrainement ⚔️
Combat contre le **Gobelin d'entrainement** (40 PV, 5 ATK)

**Actions possibles chaque tour** :
1. **Attaquer** : Attaque basique (5 dégâts)
2. **Inventaire** : Utiliser une potion
3. **Sorts** : Lancer un sort

**Sorts disponibles** :
- 👊 Coup de poing (8 dégâts, 5 mana) - De base
- 🔥 Boule de feu (18 dégâts, 15 mana) - À apprendre

**Pattern du Gobelin** :
- Normal : 5 dégâts
- Tous les 3 tours : 10 dégâts (200%)
- Les tours sont affichés pour vous aider

**Stratégie de combat** :
```
Tour 1: Attaque basique (-5 PV au gobelin)
Tour 2: Utiliser potion si nécessaire
Tour 3: Attaque basique (Gobelin attaque fort +10 dégâts!)
...
Répéter jusqu'à victoire
```

**Récompenses** :
- ⭐ 50 points d'expérience
- 🎊 Montée de niveau si exp complète

#### 6️⃣ Qui sont-ils? 🎭
Révèle les artistes cachés :
- **Partie 2** : Leonardo da Vinci (économie)
- **Partie 3** : Sun Tzu (combat)

### 4. Progression Recommandée

#### Niveau 1-5 (Débutant)
```
✅ Accomplir combats d'entraînement
✅ Gagner de l'expérience régulièrement
✅ Economiser de l'or
✅ Augmenter les PV avec équipements simples
```

#### Étape 1 : Premier équipement
```
1. Accumulez 4 potions (-12 or)
2. Achetez matériaux Chapeau (4 or + 5 or forgeage)
3. Fabriquez + Équipez Chapeau (+10 PV)
4. Nouvel total PV : Previous + 10
```

#### Étape 2 : Augmentation inventaire
```
1. Sauvegardez 30 or
2. Achetez "Augmentation d'inventaire"
3. Utilisez depuis l'inventaire
4. Nouveau max : 20 slots
```

#### Étape 3 : Maîtrise des sorts
```
1. Accumulez 25 or
2. Achetez "Livre de Sort : Boule de Feu"
3. Utilisez depuis l'inventaire
4. Boule de feu déverrouillée (18 dégâts, 15 mana)
```

### 5. Système d'Expérience

```
Gain par combat: 50 XP
Besoin pour niveau 1→2: 100 XP
Besoin pour niveau 2→3: 110 XP (augmente de 10%)
...

Bonus à chaque niveau:
- +10 PV max
- +10 Mana max
```

### 6. Conseils Tactiques

#### Combat intelligent :
- 🛡️ Utilisez les potions quand vous êtes bas (< 25 PV)
- 🔥 Boule de feu est plus forte mais consomme plus de mana
- 👊 Coup de poing est économe en mana
- 📊 Vérifiez les PV du gobelin pour prévoir fin du combat

#### Économie :
- 💰 Les matériaux sont importants - achetez-les régulièrement
- 🎒 Augmentez inventaire rapidement (limite = problème!)
- 🛑 Les équipements donnent gros bonus (+50 PV possible!)

#### Progression rapide :
1. Combattez régulièrement (+ expérience + or)
2. Vendez/utilisez items inutiles
3. Concentrez-vous d'abord sur équipement tête
4. Puis augmentez mana pour plus de sort

### 7. Dépannage

**Problème** : "Inventaire plein !"
- **Solution** : Achetez augmentation d'inventaire ou utilisez des items

**Problème** : "Pas assez d'or !"
- **Solution** : Combattez plus pour gagner de l'expérience (et récupérez des récompenses futures)

**Problème** : "Pas assez de matériaux !"
- **Solution** : Retournez au marchand, économisez, puis réessayez

**Problème** : "Pas assez de mana !"
- **Solution** : Utilisez potion de mana (10 or) ou augmentez mana avec niveau

### 8. Objectifs de Gameplay

- [ ] Atteindre niveau 5
- [ ] Apprendre des deux sorts
- [ ] Équiper les 3 armures (+50 PV bonus!)
- [ ] Augmenter inventaire 3 fois
- [ ] Accumuler 500 or
- [ ] Gagner 10 combats
- [ ] Découvrir les artistes cachés

---

**Amusez-vous bien ! 🎮✨**
