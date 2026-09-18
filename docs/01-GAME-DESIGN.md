# AUBEFER — Document de conception (Game Design Document)

> Projet RED — Ynov Campus Aix, B1 Informatique
> Équipe : Théo Sugier · Amine · Martim
> Version 1.0 — document vivant, à mettre à jour à chaque sprint

---

## 1. Le jeu en une phrase

**AUBEFER** est un jeu de rôle au tour par tour, entièrement jouable au clavier dans un terminal, où l'on incarne une recrue qui doit survivre à sa **nuit d'entraînement** dans le dernier camp avant les Marches Grises.

## 2. Pitch

Aubefer est le dernier camp fortifié avant les Marches Grises. Toute recrue y passe **une seule nuit**.

Au crépuscule, on lui donne un nom, un lignage, une besace et cent pièces d'or. Avant l'aube, elle doit avoir appris à se soigner, à négocier avec le Régisseur, à faire chauffer la Forge, et surtout à tenir debout face au **gobelin du terrain d'exercice** — le seul adversaire du camp, et celui que personne ne bat du premier coup.

Au matin, le Sergent tranche : la recrue part vers le nord, ou elle rentre chez elle.

## 3. Pourquoi ce thème

Le sujet impose une liste d'éléments : trois classes aux points de vie différents, un marchand, un forgeron, quatre matériaux de craft, des potions, un livre de sort, un gobelin d'entraînement. Un **camp d'entraînement** est le seul décor où **tous ces éléments cohabitent naturellement au même endroit** :

| Contrainte du sujet | Justification dans l'univers |
|---|---|
| Un marchand qui donne d'abord une potion **gratuitement** | Le Régisseur équipe les recrues : la première potion est fournie par le camp |
| Un forgeron accessible à tout moment | La Forge du camp tourne toute la nuit |
| Un **gobelin d'entraînement** qu'on peut affronter en boucle | C'est un gobelin capturé, gardé dans la fosse pour l'exercice |
| Le joueur ressuscite à 50 % de ses PV quand il tombe | On ne meurt pas à l'entraînement : le Sergent te relève |
| Tout se passe dans un menu unique | Le camp est petit : un feu, quatre tentes, une fosse |

**Aucune mécanique imposée n'a eu besoin d'être tordue pour entrer dans l'histoire.** C'est exactement ce que demande le critère « Respect des tâches tout en les intégrant dans un thème cohérent » (5 points).

## 4. Le camp : la carte mentale

Le menu principal n'est pas une liste d'options, c'est **le camp vu depuis le feu** :

```
                    L E   C A M P   D ' A U B E F E R

        ,^.                 [1] Ta fiche de recrue
       /   \                [2] Ta besace
      | (o) |    ~~~        [3] La tente du Régisseur   (marchand)
       \   /    (   )       [4] La Forge                (forgeron)
        `-'      ~~~        [5] Le terrain d'exercice   (entraînement)
                            [6] Les graffitis de la palissade
       le feu               [0] Quitter le camp
```

Chaque choix du menu correspond à **un lieu**, et chaque lieu correspond à **une tâche du sujet**. Un joueur qui se promène dans le camp déclenche mécaniquement toutes les fonctionnalités demandées.

## 5. Personnages

### 5.1 La recrue (le joueur)

Le joueur choisit son **nom** (lettres uniquement, reformaté avec une majuscule puis des minuscules) et son **lignage**. Le lignage est ce que le sujet appelle « classe » — on garde strictement les trois valeurs imposées :

| Lignage | PV max | PV de départ | Description de saveur |
|---|---|---|---|
| **Humain** | 100 | 50 | Polyvalent. Ce que le Sergent appelle « la moyenne » |
| **Elfe** | 80 | 40 | Fragile, mais on dit qu'ils apprennent les sorts deux fois plus vite |
| **Nain** | 120 | 60 | Encaisse. Le Régisseur lui vend rarement des potions |

Toutes les recrues commencent **niveau 1**, avec **100 pièces d'or**, le sort **Coup de poing** et une besace de **10 emplacements**.

> Les points de vie de départ valent 50 % des points de vie maximum : le Sergent ne donne jamais une recrue à pleine forme, « pour qu'elle apprenne à se soigner avant d'apprendre à frapper ».

### 5.2 Le Sergent d'Aubefer (narrateur)

Le Sergent n'est pas un personnage jouable : c'est **la voix du jeu**. Tous les messages, et surtout **tous les messages d'erreur**, passent par lui. C'est ce qui transforme une erreur technique en élément de thème :

| Situation | Message du Sergent |
|---|---|
| Besace pleine | « Ta besace déborde, recrue. Vide-la ou achète du cuir. » |
| Or insuffisant | « Tu n'as pas de quoi payer. Le Régisseur ne fait pas crédit. » |
| Matériaux manquants | « La Forge ne travaille pas le vide. Reviens avec tes peaux. » |
| Sort déjà connu | « Tu sais déjà faire ça. Garde ton or. » |
| Saisie invalide | « Parle plus clairement, recrue. » |
| Mort en entraînement | « Debout. On ne meurt pas à l'exercice. » |

**Règle d'équipe : aucun message brut du type "erreur" ou "invalid input" ne doit rester dans le jeu final.** Chaque message sort de la bouche du Sergent.

### 5.3 Le Régisseur (marchand)

Il tient la tente d'intendance. Il vend ce que le camp peut céder :

| Objet | Prix | Rôle |
|---|---|---|
| Potion de vie | **gratuit à la 1re visite**, puis 3 pièces d'or | Rend 50 PV |
| Potion de poison | 6 | Inflige 10 dégâts par seconde pendant 3 secondes |
| Livre de Sort : Boule de Feu | 25 | Apprend le sort Boule de Feu (une seule fois) |
| Fourrure de Loup | 4 | Matériau de craft |
| Peau de Troll | 7 | Matériau de craft |
| Cuir de Sanglier | 3 | Matériau de craft |
| Plume de Corbeau | 1 | Matériau de craft |
| Besace renforcée | 30 | +10 emplacements d'inventaire (3 fois maximum) |

> Le sujet impose que la première potion de vie soit gratuite (tâche 7), puis que les prix s'appliquent (tâche 14). Dans l'histoire, c'est la dotation d'entrée : **une** potion offerte à la recrue, ensuite elle paie comme tout le monde.

### 5.4 La Forge

Elle ne vend rien : elle **transforme**. Trois pièces d'équipement, 5 pièces d'or de main-d'œuvre chacune, plus les matériaux qui sont consommés :

| Équipement | Emplacement | Matériaux | Bonus |
|---|---|---|---|
| Chapeau de l'aventurier | Tête | 1 Plume de Corbeau + 1 Cuir de Sanglier | +10 PV max |
| Tunique de l'aventurier | Torse | 2 Fourrures de Loup + 1 Peau de Troll | +25 PV max |
| Bottes de l'aventurier | Pieds | 1 Fourrure de Loup + 1 Cuir de Sanglier | +15 PV max |

Un équipement fabriqué part dans la besace. Il faut ensuite **l'enfiler** depuis la besace pour qu'il compte. Enfiler une pièce alors qu'une autre occupe déjà l'emplacement renvoie l'ancienne dans la besace.

### 5.5 Le gobelin d'entraînement

Le seul adversaire du jeu, et c'est assumé : on est à l'entraînement.

- PV max : 40 · PV actuels : 40 · Attaque : 5
- **Schéma de combat** : il frappe à 100 % de son attaque, sauf **tous les 3 tours** où il frappe à 200 %. Le joueur qui observe finit par comprendre le rythme — c'est tout l'intérêt du terrain d'exercice.

## 6. Boucle de jeu

```
Création de la recrue
        │
        ▼
   ┌─────────────────── LE FEU DE CAMP (menu principal) ──────────────────┐
   │                                                                      │
   │  Fiche ──► lire ses stats                                            │
   │  Besace ──► utiliser potion / poison / livre de sort / s'équiper     │
   │  Régisseur ──► dépenser de l'or, remplir la besace                   │
   │  Forge ──► transformer matériaux + or en équipement                  │
   │  Terrain d'exercice ──► COMBAT TOUR PAR TOUR                         │
   │        │                                                             │
   │        └─► victoire ou défaite ──► retour au feu de camp             │
   │                                                                      │
   └──────────────────────────────────────────────────────────────────────┘
        │
        ▼
   Quitter le camp
```

La boucle est volontairement courte : **acheter → fabriquer → s'équiper → se battre → mourir → recommencer plus fort**. C'est la boucle de progression qui donne envie de relancer un combat, et c'est elle qu'on montrera en soutenance.

## 7. Le combat, tour par tour

```
┌─ TOUR 3 ─────────────────────────────────────────────┐
│                                                      │
│   Théo            Gobelin d'entraînement             │
│   PV  45 / 110    PV  22 / 40                        │
│   [##########------]  [########--------]             │
│                                                      │
│   [1] Attaquer      [2] Besace      [3] Fuir         │
└──────────────────────────────────────────────────────┘
```

Déroulé d'un tour :

1. Affichage du numéro de tour et des deux barres de vie
2. **Tour du joueur** : Attaquer (attaque basique, 5 dégâts / sorts en mission bonus) ou ouvrir la Besace (l'objet utilisé consomme le tour)
3. **Tour du gobelin** : `goblinPattern` — 100 % de son attaque, ou 200 % tous les 3 tours
4. Contrôle de fin : si l'un des deux tombe à 0 PV ou moins, le combat s'arrête et le joueur revient au feu de camp

Chaque action est **annoncée à l'écran** : qui attaque, qui reçoit, combien de dégâts, et les PV restants sur les PV max.

## 8. Identité visuelle en ASCII

Le jeu est en terminal : c'est le seul support visuel dont on dispose, donc on le soigne. Règles d'équipe :

- **Une bannière ASCII** au lancement (le mot AUBEFER)
- **Un encadré** pour chaque écran (fiche, besace, marchand, forge, combat) — même largeur partout
- **Des barres de vie** en caractères pleins/vides, jamais un simple nombre
- **L'écran est nettoyé** entre deux écrans, pour que le terminal ne défile pas à l'infini
- **Des couleurs ANSI** : rouge pour les dégâts, vert pour les soins, jaune pour l'or, gris pour le Sergent
- **Aucune faute d'orthographe** dans les textes affichés : c'est ce que le jury lit pendant toute la démo

## 9. Ce qui fait gagner les points « initiative personnelle »

Fonctionnalités non exigées que l'on ajoute volontairement :

1. **La première potion offerte une seule fois** (mémorisée), puis facturée — cohérence marchande
2. **Les graffitis de la palissade** : la réponse à la Mission 6, intégrée comme un vrai lieu du camp
3. **Barres de vie et écrans encadrés** au lieu de lignes de texte brutes
4. **Messages d'erreur incarnés** par le Sergent, jamais de message technique
5. **Récapitulatif de fin de combat** : tours joués, dégâts infligés, dégâts reçus
6. **Fuite** possible en combat (le gobelin garde un tour d'avance : fuir coûte une attaque gratuite)

## 10. Lexique imposé — ce qu'on ne renomme PAS

Pour que le correcteur retrouve immédiatement le sujet, on conserve **strictement** les noms du brief : `Character`, `Equipment`, `Monster`, `initCharacter`, `displayInfo`, `accessInventory`, `takePot`, `isDead`, `poisonPot`, `spellBook`, `characterCreation`, `upgradeInventorySlot`, `initGoblin`, `goblinPattern`, `characterTurn`, `trainingFight`, ainsi que « Humain / Elfe / Nain », « Potion de vie », « Potion de poison », « Livre de Sort : Boule de Feu », « Fourrure de Loup », « Peau de Troll », « Cuir de Sanglier », « Plume de Corbeau », « Chapeau / Tunique / Bottes de l'aventurier », « Gobelin d'entrainement », « Coup de poing », « Boule de Feu ».

Le thème vit dans **les textes affichés**, pas dans le renommage du code. C'est volontaire : on gagne les points de thème sans perdre les points de conformité.
