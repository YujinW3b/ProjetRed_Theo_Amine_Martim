# AUBEFER

Mini-jeu en ligne de commande écrit en Go. Projet RED, B1 Informatique, Ynov Campus Aix.

Vous êtes une recrue au camp d'entraînement d'Aubefer. Vous choisissez un nom et un
lignage, vous équipez votre personnage, vous achetez chez le Régisseur, vous forgez,
et vous vous battez contre le gobelin d'entraînement. Tout se joue au clavier, dans
le terminal.

## Lancer le jeu

Il faut Go 1.21 ou plus.

```
git clone https://github.com/YujinW3b/ProjetRed_Theo_Amine_Martim.git
cd ProjetRed_Theo_Amine_Martim
go run ./src
```

## Le camp

Le menu principal est le camp. Chaque option est un lieu.

```
[1] Ta fiche de recrue        [4] La Forge
[2] Ta besace                 [5] Le terrain d'exercice
[3] La tente du Régisseur     [6] Les graffitis de la palissade
                              [0] Quitter
```

On tape le chiffre, on valide. Toute autre saisie fait râler le Sergent, rien ne plante.

## Les lignages

| Lignage | PV max | PV au départ |
|---|---|---|
| Humain | 100 | 50 |
| Elfe | 80 | 40 |
| Nain | 120 | 60 |

Une recrue commence niveau 1, avec 100 pièces d'or, le sort Coup de poing et une
besace de 10 emplacements. Les PV de départ valent la moitié du maximum : le Sergent
ne donne jamais une recrue à pleine forme.

## Organisation du code

```
src/        tout le code Go, en package main
docs/       le document de gestion de projet
go.mod      module aubefer
```

Un fichier par domaine, pour éviter que deux personnes travaillent au même endroit :
`character.go`, `menu.go`, `creation.go`, `inventory.go`, `merchant.go`,
`blacksmith.go`, `equipment.go`, `combat.go`, `monster.go`, `items.go`, `display.go`.

## L'équipe

- Théo Sugier — noyau, menu, création de personnage, combat
- Amine — or, marchand, forge, potions et sorts
- Martim — structures, affichage, équipement, monstre

Suivi des tâches sur Trello, une carte par tâche du sujet. Une branche et une pull
request par tâche, relue par un autre membre avant fusion.