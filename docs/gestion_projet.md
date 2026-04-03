Document de Gestion de Projet - RPG RED

    Étudiant : Aboubakar Dewa

    Collaborateur : Shayton

    Formation : Bachelor 1 - Ytrack Ynov

    Technologie : Golang

1. Objectifs du Projet

Réalisation d'un jeu de rôle (RPG) en ligne de commande (CLI) avec gestion de personnage, inventaire, sorts et combats au tour par tour.


2. État d'avancement des Tâches (WBS)
Tâche	Description	État
T1 & T2	Structure Character et initialisation (PV, Classe, Nom)	✅ Terminé
T3 & T6	Affichage des statistiques et Menu Principal	✅ Terminé
T4 & T5	Inventaire et Potion de vie (+50 PV)	✅ Terminé
T8 & T11	Gestion de la mort (Wasted) et formatage du nom	✅ Terminé
T9 & T10	Potion de poison (dégâts/sec) et Livre de Sort (Boule de feu)	✅ Terminé
T19 à T22	Combat contre le Gobelin avec affichage des tours et PV	✅ Terminé


3. Organisation du code source (/src)

Le code a été découpé en plusieurs fichiers pour une meilleure clarté :

    main.go : Boucle principale et navigation.

    character.go : Création du héros (Humain, Elfe, Nain, Rayann, Hollow).

    combat.go : Logique de duel, affichage des tours et PV restants.

    inventory.go : Gestion des objets et effets de statut (poison).

    actions.go : Fonctions système (Stats, Mort, Marchand).


4. Notes Techniques & Bonus

    Formatage : Utilisation de strings.Title pour assurer la propreté du nom du joueur.

    Combats : Implémentation d'un pattern de dégâts (critique tous les 3 tours) pour dynamiser le gameplay.
