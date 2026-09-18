# AUBEFER — Architecture technique

> Comment les fichiers sont découpés, et **pourquoi**. À lire avant d'écrire la première ligne.

## 1. Arborescence du dépôt

```
projet-red_AUBEFER/
├── README.md                  ← présentation + installation + lancement (exigé par le sujet)
├── .gitignore
├── go.mod
├── docs/
│   ├── 00-GESTION-DE-PROJET.md
│   ├── 01-GAME-DESIGN.md
│   ├── 02-SPECIFICATIONS.md
│   ├── 03-ARCHITECTURE.md      ← ce fichier
│   ├── 04-SOUTENANCE.md
│   └── 05-GRILLE-AUTO-EVALUATION.md
└── src/
    ├── main.go                 ← 🟢 point d'entrée : lance la création puis la boucle de menu
    ├── menu.go                 ← 🟢 feu de camp, sous-menus, lecture et validation des saisies
    ├── creation.go             ← 🟢 characterCreation (validation et reformatage du nom)
    ├── equipment.go            ← 🟢 structure Equipment, équiper, déséquiper, recalcul des PV max
    ├── combat.go               ← 🟢 trainingFight, characterTurn, goblinPattern, initiative, XP
    ├── inventory.go            ← 🟡 addInventory, removeInventory, limite, upgradeInventorySlot
    ├── items.go                ← 🟡 takePot, poisonPot, spellBook, catalogue des objets et prix
    ├── merchant.go             ← 🟡 tente du Régisseur : menu, achat, débit de l'or
    ├── blacksmith.go           ← 🟡 la Forge : recettes, contrôles, fabrication
    ├── character.go            ← 🟠 structure Character, initCharacter, displayInfo, isDead
    ├── monster.go              ← 🟠 structure Monster, initGoblin
    └── display.go              ← 🟠 accessInventory, bannière ASCII, encadrés, barres de vie, couleurs
```

## 2. La règle qui évite 90 % des conflits Git

**Un domaine = un fichier = un responsable.**

| Fichier | Responsable |
|---|---|
| `main.go`, `menu.go`, `creation.go`, `equipment.go`, `combat.go` | 🟢 Théo |
| `inventory.go`, `items.go`, `merchant.go`, `blacksmith.go` | 🟡 Amine |
| `character.go`, `monster.go`, `display.go` | 🟠 Martim |

Personne ne touche au fichier d'un autre sans le prévenir sur Discord.

**Une seule exception, et elle est structurante** : la structure `Character` vit dans `character.go` (Martim), mais Amine y ajoute l'or (T13 est à Martim, mais T10 `skill` et M4 `mana` sont à Amine) et Théo l'équipement (T17). **Toute modification de `Character` est annoncée avant d'être poussée**, et se fait par ajout d'un champ, jamais par réécriture. C'est le seul point du projet où trois personnes touchent au même fichier, donc le seul où on peut vraiment se marcher dessus.

Comme trois personnes travaillent en même temps sur le même jeu, c'est le découpage des fichiers — pas la bonne volonté — qui empêche les conflits.

## 3. Les trois fonctions qui structurent tout le reste

### `addInventory` / `removeInventory` — 🟡 Amine

Le sujet les suggère en « astuce » à la tâche 7. **On les traite comme une obligation.** Ce sont les **seuls** points d'entrée pour modifier l'inventaire, de bout en bout du projet.

Tout ce qui y est centralisé :

- le contrôle de la limite d'emplacements (T12) ;
- la capacité courante, qui change avec `upgradeInventorySlot` (T18) ;
- le message d'erreur du Sergent quand la besace déborde.

Si un jour quelqu'un ajoute un objet sans passer par `addInventory`, la limite d'inventaire devient contournable **et le bug est invisible** — il ne se verra qu'en soutenance. C'est le point que le relecteur doit chercher dans chaque PR.

### `isDead` — 🟠 Martim

Appelée **après chaque source de dégâts** : chaque tick de poison, chaque attaque du gobelin. Elle lit les PV max **au moment de l'appel**, parce que l'équipement les fait varier.

### Le recalcul des PV max — 🟢 Théo

Les PV max ne sont pas une valeur fixe : ils valent *base du lignage + bonus des équipements portés + bonus de niveau*. Une seule fonction fait ce calcul, et elle est appelée à chaque fois qu'un équipement change ou qu'un niveau est gagné. Elle replafonne les PV actuels si le maximum a baissé.

**Ne recopiez jamais ce calcul ailleurs.** Trois copies = trois occasions d'oublier d'en mettre une à jour.

## 4. Gestion des saisies utilisateur

Tout le jeu tourne autour d'une seule chose : lire ce que l'utilisateur tape sans planter. Une **fonction unique** de lecture est utilisée partout, et elle :

1. lit la ligne entière (pas juste un caractère) ;
2. retire les espaces et les retours à la ligne autour ;
3. vérifie que la saisie correspond à un choix proposé ;
4. si non, affiche un message du Sergent et **redemande**, sans quitter le menu.

Cas à tester dans chaque menu, systématiquement : **saisie vide**, **lettre au lieu d'un chiffre**, **nombre hors bornes**, **espaces avant/après**.

> C'est ce qui départage un projet « en cours d'acquisition » d'un projet « acquis » sur le critère *retours utilisateurs pertinents*. Un jury qui tape n'importe quoi et voit le jeu répondre calmement a immédiatement une bonne impression.

## 5. Conventions de code

- Noms de fonctions **exactement** ceux du sujet (`initCharacter`, `takePot`, `goblinPattern`…) — le correcteur doit les retrouver
- Variables et commentaires en français, cohérents avec les textes du jeu
- Une fonction = une responsabilité ; si elle dépasse une trentaine de lignes, elle en fait probablement deux
- Aucun nombre magique dans le code : les prix, les dégâts, les bonus et les capacités sont des constantes nommées, regroupées en haut de leur fichier
- **Aucun texte affiché à l'utilisateur écrit directement au milieu de la logique** : les messages du Sergent sont regroupés, pour rester relisables et sans fautes

## 6. Ordre de développement recommandé

```
VEN 18   T1 ─► T2 ─► T3   (🟠 à trois)        T6 (🟢)
              │
DIM 20        └── rattrapage seulement
              │
LUN 21   T4 (🟠) · T16 (🟠) 🔴 · T19 (🟠) 🔴   T7 (🟡) 🔴 · T5 (🟡)   T11 (🟢)
              │
MAR 22   T8 (🟠) · T13 (🟠)   T9 · T10 · T12 (🟡)   T17 (🟢)
              │                                   ◄── Partie 1 complète
MER 23   M5 habillage (🟠)   T14 · T15 · T18 (🟡)   T20 ─► T21 ─► T22 (🟢)
              │                                   ◄── LES 22 TÂCHES
JEU 24   M6 (🟠)   M3 · M4 (🟡)   M1 · M2 (🟢)   puis rendu à trois
```
