# Gestion de projet — AUBEFER

Projet RED, B1 Informatique, Ynov Campus Aix. Rendu le jeudi 24 septembre 2026.
Équipe de trois : Théo Sugier, Amine, Martim.

## Qui fait quoi

| Membre | Périmètre |
|---|---|
| Théo | Menu, création de personnage, combat (T6, T11, T17, T20, T21, T22) |
| Amine | Économie : marchand, forge, prix, potions, sorts (T5, T7, T9, T10, T12, T14, T15, T18) |
| Martim | Structures et affichage, or, équipement, monstre (T1 à T4, T8, T13, T16, T19) |

Pas de chef de projet séparé. Les décisions de game design se prennent à trois et
sont écrites : une décision non écrite n'existe pas.

## Planning

| Jour | Objectif de fin de journée |
|---|---|
| Ven 18 | Dépôt, board, thème validé, ça compile chez les trois |
| Lun 21 | Structures, menu, inventaire, création de personnage |
| Mar 22 | Marchand, forge, or, prix, tours de combat |
| Mer 23 | Combat complet — les 22 tâches obligatoires sont livrées |
| Jeu 24 | Bonus si le socle est vert, relecture, démo, dépôt Moodle |

Jalon non négociable : les 22 tâches obligatoires sont finies mercredi soir. Aucun
bonus tant que le socle n'est pas complet, parce qu'un bonus ne rattrape jamais une
tâche obligatoire manquante dans la grille.

## Dépendances

Certaines tâches en bloquent plusieurs autres. Elles sont prioritaires.

```
T7  addInventory / removeInventory  ->  T12, T14, T15, T17, T18
T16 Equipment                       ->  T17
T19 Monster                         ->  T20, T21, T22
T22 trainingFight                   ->  M1, M2
```

Amine livre T7 en premier. Martim livre T16 et T19 avant le reste.

## Méthode Git

- `main` doit toujours compiler. Personne ne pousse dessus directement.
- Une branche par tâche : `feat/t11-creation`, `fix/t12-limite-inventaire`.
- Une pull request par tâche, relue par un autre membre avant fusion.
- Le relecteur vérifie les critères de validation de la tâche avant d'approuver.
- On lance `go run ./src` avant chaque push. Un build cassé sur `main` bloque les trois.

Une carte passe en Terminé quand le code compile, que les critères de validation sont
cochés, que les cas limites ont été testés à la main, que la PR est relue et fusionnée,
et qu'aucun message technique ne reste affiché à l'écran.

## Outils

Trello pour le suivi des tâches, GitHub pour le code et les relectures, Discord pour
la communication quotidienne et le partage d'écran.

## Risques

| Risque | Parade |
|---|---|
| Deux personnes modifient la structure `Character` en même temps | La structure est figée dès le début, toute modification passe par Théo et est annoncée |
| Conflits Git à répétition | Un fichier par domaine, deux personnes ne touchent presque jamais le même |
| Quelqu'un est bloqué sans le dire | Point du matin obligatoire, règle des 30 minutes : bloqué plus longtemps, on demande |
| On se perd dans les bonus et le socle n'est pas fini | Jalon du mercredi soir, aucun bonus avant |
| La démo plante devant le jury | Deux répétitions complètes le jeudi matin |

## Avancement

À mettre à jour à chaque fin de journée.

| Partie | Tâches | Fait |
|---|---|---|
| Partie 1 | T1 à T12 | 6 / 12 |
| Partie 2 | T13 à T18 | 0 / 6 |
| Partie 3 | T19 à T22 | 0 / 4 |
| Bonus | M1 à M6 | 0 / 6 |