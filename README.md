```
   █████╗ ██╗   ██╗██████╗ ███████╗███████╗███████╗██████╗
  ██╔══██╗██║   ██║██╔══██╗██╔════╝██╔════╝██╔════╝██╔══██╗
  ███████║██║   ██║██████╔╝█████╗  █████╗  █████╗  ██████╔╝
  ██╔══██║██║   ██║██╔══██╗██╔══╝  ██╔══╝  ██╔══╝  ██╔══██╗
  ██║  ██║╚██████╔╝██████╔╝███████╗██║     ███████╗██║  ██║
  ╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝╚═╝     ╚══════╝╚═╝  ╚═╝
        une nuit au dernier camp avant les Marches Grises
```

# AUBEFER

**Jeu de rôle au tour par tour, jouable entièrement au clavier dans un terminal.**
Projet RED — Ymmersion — Ynov Campus Aix, B1 Informatique.

---

## Le jeu

Aubefer est le dernier camp fortifié avant les Marches Grises. Toute recrue y passe **une seule nuit**.

Au crépuscule, on lui donne un nom, un lignage, une besace et cent pièces d'or. Avant l'aube, elle doit avoir appris à se soigner, à négocier avec le Régisseur, à faire chauffer la Forge, et surtout à tenir debout face au gobelin du terrain d'exercice.

Au matin, le Sergent tranche : la recrue part vers le nord, ou elle rentre chez elle.

### Ce qu'on peut faire

| Lieu du camp | Ce qu'on y fait |
|---|---|
| **Ta fiche de recrue** | Consulter ses statistiques, son or, ses sorts et son équipement |
| **Ta besace** | Utiliser ses objets : potions, livre de sort, pièces d'équipement |
| **La tente du Régisseur** | Acheter potions, matériaux, grimoires et besaces renforcées |
| **La Forge** | Fabriquer chapeau, tunique et bottes de l'aventurier à partir de matériaux |
| **Le terrain d'exercice** | Affronter le gobelin d'entraînement au tour par tour |
| **Les graffitis de la palissade** | Lire ce que les recrues des autres nuits ont gravé dans le bois |

### Les trois lignages

| Lignage | PV max | PV de départ |
|---|---|---|
| Humain | 100 | 50 |
| Elfe | 80 | 40 |
| Nain | 120 | 60 |

---

## Installation

### Prérequis

- **Go 1.21 ou supérieur** — vérifier avec `go version`
- Un terminal supportant les couleurs ANSI (Terminal macOS, GNOME Terminal, Windows Terminal, iTerm2…)
- Git

### Récupérer le projet

```bash
git clone <url-du-depot> projet-red_AUBEFER
cd projet-red_AUBEFER
```

---

## Lancement

Depuis la racine du projet :

```bash
go run ./src
```

Ou en compilant un exécutable :

```bash
go build -o aubefer ./src
./aubefer
```

Sous Windows :

```powershell
go build -o aubefer.exe ./src
.\aubefer.exe
```

### Comment jouer

Tout se joue au clavier. À chaque écran, le jeu propose des choix numérotés : on tape le numéro et on valide avec Entrée. `0` ramène toujours à l'écran précédent, ou quitte le camp depuis le feu de camp.

Une saisie incorrecte ne fait jamais planter le jeu : le Sergent vous le fera remarquer, et vous pourrez recommencer.

---

## Structure du dépôt

```
projet-red_AUBEFER/
├── README.md        ← ce fichier
├── go.mod
├── docs/            ← conception, spécifications, gestion de projet, soutenance
└── src/             ← code source du jeu
```

Le dossier `docs/` contient :

| Fichier | Contenu |
|---|---|
| `00-GESTION-DE-PROJET.md` | Équipe, répartition, planning, méthode Git, risques |
| `01-GAME-DESIGN.md` | Univers, personnages, économie, boucle de jeu |
| `02-SPECIFICATIONS.md` | Les 22 tâches et 6 missions, avec leurs critères de validation |
| `03-ARCHITECTURE.md` | Découpage des fichiers et décisions techniques |
| `04-SOUTENANCE.md` | Scénario de démonstration et préparation de l'oral |
| `05-GRILLE-AUTO-EVALUATION.md` | Vérification face à la grille d'évaluation |

---

## L'équipe

| | Rôle | Périmètre |
|---|---|---|
| **Théo Sugier** | Lead technique | Menus, création de personnage, équipement, combat au tour par tour |
| **Amine** | Développeur | Inventaire, objets, marchand, forge, sorts |
| **Martim** | Développeur | Structures de données, affichage, habillage ASCII du jeu |

Suivi du projet sur Trello — une carte par tâche du sujet, une Pull Request relue par tâche.

---

## État d'avancement

| Partie | Avancement |
|---|---|
| Partie 1 — Personnage & fonctionnalités de base | 0 / 12 |
| Partie 2 — Économie & fabrication | 0 / 6 |
| Partie 3 — Combat au tour par tour | 0 / 4 |
| Missions bonus | 0 / 6 |

---

*Projet réalisé dans le cadre de l'Ymmersion — Ynov Campus Aix.*
