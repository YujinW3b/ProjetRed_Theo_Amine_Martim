# Gestion de projet — AUBEFER

Projet RED, B1 Informatique, Ynov Campus Aix.
Soutenance le vendredi 25 septembre 2026, de 14h00 à 16h00.
Équipe de trois : Théo SUGIER, Amine KOUAIDI, Martim GALHARDO.

## Qui fait quoi

Pas de chef de projet séparé : à trois, le lead technique tient aussi le board. Les
décisions de game design se prennent à trois et sont écrites — une décision non
écrite n'existe pas.

| Membre | Périmètre | Fichiers |
|---|---|---|
| Théo | Menu du camp, création de personnage, combat, habillage, dépôt | `menu.go`, `creation.go`, `combat.go`, `game.go`, `ascii.go` |
| Amine | Besace, or, Régisseur, Forge, sorts et mana, équipement | `inventory.go`, `items.go`, `merchant.go`, `blacksmith.go`, `equipment.go` |
| Martim | Structures `Character` et `Equipment`, fiche de recrue, gobelin | `character.go`, `monster.go` |

Répartition des commits au 24 septembre : 93 sur `main`, dont 49 Théo, 26 Martim,
20 Amine.

## Planning

| Jour | Objectif de fin de journée |
|---|---|
| Ven 18 | Dépôt, board, thème validé, ça compile chez les trois |
| Lun 21 | Structures, menu, besace, création de personnage |
| Mar 22 | Régisseur, Forge, or, prix, tours de combat |
| Mer 23 | Combat complet — les 22 tâches obligatoires sont écrites |
| Jeu 24 | Bonus, relecture, README, support de soutenance, dépôt Moodle |
| Ven 25 | Répétition de la démo le matin, soutenance à 14h00 |

Jalon non négociable : les 22 tâches obligatoires sont finies mercredi soir. Aucun
bonus tant que le socle n'est pas complet — un bonus ne rattrape jamais une tâche
obligatoire manquante dans la grille.

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
- Commits au format `type(portée) : description`, en minuscules et sans accent.
- On lance `go run main.go` avant chaque push. Un build cassé sur `main` bloque les trois.

Une carte passe en Terminé quand le code compile, que les critères de validation sont
cochés, que les cas limites ont été testés à la main, que la PR est relue et fusionnée,
et qu'aucun message technique ne reste affiché à l'écran.

## Outils

Trello pour le suivi des tâches, GitHub pour le code et les relectures, Discord pour
la communication quotidienne et le partage d'écran.

## Risques et ce qui s'est réellement passé

| Risque prévu | Parade | Est-ce arrivé ? |
|---|---|---|
| Deux personnes modifient `Character` en même temps | Structure figée tôt, toute modification annoncée | Oui, deux fois. Résolu en fusionnant les deux versions |
| Conflits Git à répétition | Un fichier par domaine | Oui au début, quasiment disparus après le découpage |
| Quelqu'un est bloqué sans le dire | Point du matin, règle des 30 minutes | Non |
| On se perd dans les bonus et le socle n'est pas fini | Jalon du mercredi soir | Non, le socle est passé en premier |
| La démo plante devant le jury | Deux répétitions complètes le vendredi matin | À vérifier le jour J |

Trois incidents non prévus au départ :

- `go.mod` portait la version de Go de la machine qui a créé le module. Les deux
  autres postes refusaient de compiler. Corrigé en fixant `go 1.21`.
- Des marqueurs de conflit ont été poussés sur `main` deux fois. `main` ne compilait
  plus pour personne. D'où la règle : on relance le jeu avant chaque push.
- Deux façons de lire le clavier coexistaient et produisaient des saisies fantômes
  dans le menu. Règle adoptée : un seul lecteur d'entrée dans tout le jeu.

## Avancement au 24 septembre

| Partie | Tâches | Écrites |
|---|---|---|
| Partie 1 | T1 à T12 | 12 / 12 |
| Partie 2 | T13 à T18 | 6 / 6 |
| Partie 3 | T19 à T22 | 4 / 4 |
| Bonus | M1 à M6 | 3 / 6 |

Réserves assumées :

- `accessForgeron` (T15) et `equiperItem` (T17) sont écrites mais ne sont reliées à
  aucune entrée du menu. Le code existe, un joueur ne peut pas y accéder.
- M1 : l'initiative est calculée par lignage et stockée, mais l'ordre des tours ne
  s'en sert pas encore.
- M2 : `initGoblin` renseigne `XPDonnee`, mais `Character` n'a pas de champ
  d'expérience, donc rien ne consomme cette valeur.

Bonus livrés : M3 combat magique, M4 mana, M5 habillage ASCII et couleurs.
