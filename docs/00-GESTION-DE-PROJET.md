# AUBEFER — Document de gestion de projet

> Projet RED — Ymmersion — Ynov Campus Aix, B1 Informatique
> Dépôt : `projet-red_AUBEFER` · Board Trello : *Projet RED — AUBEFER*

---

## 1. L'équipe

| Membre | Rôle | Périmètre | Responsabilité transverse |
|---|---|---|---|
| **Théo Sugier** | Lead technique | Menus, création de personnage, équipement, **tout le combat** | Architecture, arbitrage des conflits Git, relecture finale |
| **Amine** | Développeur — Économie | Inventaire, objets, marchand, forge, sorts | Cohérence des règles économiques, équilibrage |
| **Martim** | Développeur — Personnage & interface | Structures de données, affichage, **tout l'habillage ASCII** | Lisibilité du jeu à l'écran, README |

**Comment on a réparti.** Pas à parts égales en nombre de lignes, mais **à charge équivalente**. Théo prend six tâches parmi les plus délicates du sujet (le combat, le recalcul des points de vie maximum, la validation des saisies). Amine et Martim en prennent dix chacun, plus nombreuses mais chacune bornée à une seule chose.

**Le lot de Martim est volontairement composé de tâches à périmètre fermé** : des structures de données, des fonctions d'affichage, des valeurs à poser. Aucune ne demande de tenir plusieurs états en tête en même temps. C'est le bon lot pour quelqu'un qui débute en Go, et ce n'est pas un lot au rabais : **l'habillage ASCII est nommément cité dans la grille d'évaluation**, et c'est la première chose que le jury voit à l'écran pendant toute la démonstration.

Pas de chef de projet séparé : à trois, le lead technique tient aussi le board. Les décisions de game design se prennent **à trois** et sont écrites dans `01-GAME-DESIGN.md` — une décision non écrite n'existe pas.

## 2. Répartition des tâches

### 🟢 Théo — 6 tâches + 2 missions

| Tâche | Pourquoi c'est pour lui |
|---|---|
| **T06** Menu principal (`switch`) + fonction unique de lecture de saisie | C'est l'architecture : tout le jeu passe par cette boucle |
| **T11** `characterCreation` | Validation du nom, reformatage, redemande en boucle — le piège du sujet |
| **T17** Équiper / remplacer + recalcul des PV max | La tâche la plus piégeuse du sujet (voir sa spec) |
| **T20** `goblinPattern` | Le schéma « tous les 3 tours » |
| **T21** `characterTurn` | Interface de combat |
| **T22** `trainingFight` | La boucle de combat qui orchestre tout le reste |
| **M1** Initiative · **M2** Expérience | Se greffent sur le combat, donc sur son code |

### 🟡 Amine — 8 tâches + 2 missions

| Tâche | |
|---|---|
| **T07** Marchand + `addInventory` / `removeInventory` | 🔴 bloquante pour cinq autres cartes |
| **T05** `takePot` · **T09** `poisonPot` · **T10** `spellBook` | Les objets consommables |
| **T12** Limite d'inventaire · **T18** `upgradeInventorySlot` | La capacité de la besace |
| **T14** Les prix · **T15** La Forge | L'économie et le craft |
| **M3** Combat magique · **M4** Mana | Se greffent sur ses sorts |

### 🟠 Martim — 8 tâches + 2 missions

| Tâche | Pourquoi c'est borné |
|---|---|
| **T01** Structure `Character` | Poser six champs — fait à trois le premier jour |
| **T02** `initCharacter` | Remplir une structure avec des paramètres |
| **T03** `displayInfo` | Afficher, ne rien modifier |
| **T04** `accessInventory` | Lister et numéroter |
| **T08** `isDead` | Une condition, une division par deux |
| **T13** L'or | Ajouter un champ, l'initialiser à 100 |
| **T16** Structure `Equipment` | Trois champs |
| **T19** Structure `Monster` + `initGoblin` | Quatre champs et quatre valeurs fixes |
| **M5** Habillage ASCII complet | Bannière, encadrés, barres de vie, couleurs — visuel, sans logique, **et c'est dans la grille** |
| **M6** Les graffitis de la palissade | Afficher deux noms |
| **README.md** final | |

### Dépendances à surveiller

```
T07 (addInventory/removeInventory) ──► T12, T14, T15, T17, T18
T16 (Equipment) ─────────────────────► T17
T19 (Monster) ───────────────────────► T20, T21, T22
T10 (spellBook) ─────────────────────► M3 ──► M4
T22 (trainingFight) ─────────────────► M1, M2
```

Conséquence : **Amine livre T07 en priorité absolue**, et **Martim livre T16 et T19 dès lundi**. Ces trois cartes sont marquées 🔴 sur le board.

## 3. Planning — jusqu'au JEUDI 24

⚠️ **Cinq jours de travail : vendredi 18, puis lundi 21 → jeudi 24.** Le week-end n'est pas dans le plan (voir plus bas).

### VENDREDI 18 — Fondations + noyau, à trois

Personne ne part de son côté aujourd'hui.

| Qui | Quoi |
|---|---|
| Théo | Dépôt Git, arborescence, `go.mod` — ça compile chez les trois (30 min) |
| À trois | Thème validé et figé, GDD lu par les trois (20 min) |
| **Martim au clavier**, les deux autres relisent | **T01** `Character` · **T02** `initCharacter` · **T03** `displayInfo` |
| Théo, en parallèle | **T06** menu principal + la fonction unique de lecture de saisie |

> Martim tient le clavier sur les trois premières tâches **exprès** : ce sont les fondations, il faut qu'il les ait écrites lui-même pour suivre le reste de la semaine, et c'est le seul moment où les deux autres sont là pour relire en direct.

**Fin de vendredi** : la structure `Character` est figée, le menu tourne.

### SAMEDI 19 – DIMANCHE 20 — rattrapage seulement

**Aucune tâche n'est planifiée le week-end.** Un plan qui compte sur le week-end est un plan qui a déjà glissé.

Qui a du retard le rattrape. Qui n'en a pas peut avancer **M5** (bannière ASCII, encadrés) : c'est du visuel pur, ça ne casse rien chez les autres.

### LUNDI 21 — les cartes bloquantes

| 🟢 Théo | 🟡 Amine | 🟠 Martim |
|---|---|---|
| **T11** `characterCreation` | 🔴 **T07** Marchand + `addInventory`/`removeInventory` · **T05** `takePot` | **T04** `accessInventory` · 🔴 **T16** `Equipment` · 🔴 **T19** `Monster` + `initGoblin` |

**Fin de lundi** : on crée son personnage, le marchand fonctionne, les structures que les autres attendent existent. Si T07, T16 et T19 ne sont pas finies ce soir-là, on prend une heure de plus — on ne reporte pas.

### MARDI 22 — on ferme la Partie 1

| 🟢 Théo | 🟡 Amine | 🟠 Martim |
|---|---|---|
| **T17** équiper / remplacer + recalcul des PV max | **T09** `poisonPot` · **T10** `spellBook` · **T12** limite d'inventaire | **T08** `isDead` · **T13** l'or |

**Fin de mardi** : Partie 1 complète. **C'est le point de contrôle de la semaine** (voir plus bas).

### MERCREDI 23 — on ferme les 22 tâches

| 🟢 Théo | 🟡 Amine | 🟠 Martim |
|---|---|---|
| **T20** `goblinPattern` · **T21** `characterTurn` · **T22** `trainingFight` | **T14** les prix · **T15** la Forge · **T18** besace renforcée | **M5** habillage : bannière, encadrés, barres de vie, couleurs, nettoyage d'écran |

🎯 **Fin de mercredi : LES 22 TÂCHES OBLIGATOIRES SONT LIVRÉES.** Jalon non négociable.

### JEUDI 24 — bonus le matin, rendu l'après-midi

**Matin**

| 🟢 Théo | 🟡 Amine | 🟠 Martim |
|---|---|---|
| **M1** initiative · **M2** expérience | **M3** combat magique · **M4** mana | **M6** les graffitis · finitions d'habillage |

**Après-midi, à trois — on ne code plus, on rend**

| Durée | Quoi |
|---|---|
| 1 h | Relecture croisée du code |
| 30 min | Chasse aux messages techniques restants et aux fautes d'orthographe |
| 45 min | `README.md` (Martim) + document de gestion de projet de l'école (Théo) |
| 30 min | Grille d'auto-évaluation : chacun note seul, puis on compare |
| 45 min | **Deux répétitions complètes de la démo, chronomètre en main** |
| 15 min | Vérification du rendu et dépôt du lien sur Moodle |

### Si on prend du retard

On coupe **dans cet ordre**, et jamais ailleurs :

1. **M4** mana
2. **M2** expérience
3. **M1** initiative
4. **M3** combat magique

**On ne coupe jamais dans les 22 tâches, ni dans M5 (l'ASCII art est un indicateur de la grille), ni dans la répétition de la démo (8 points sur 40).** Un bonus ne rattrape jamais une tâche obligatoire manquante.

### Le point de contrôle qui sauve la semaine

**Mardi 22 au soir, on fait le compte.** Si la Partie 1 n'est pas complète, on annule tous les bonus **à ce moment-là**, et le jeudi matin sert à finir les 22 tâches au lieu des missions.

Décider ça le mardi coûte une discussion de dix minutes. Le découvrir le jeudi après-midi coûte la note.

## 4. Méthode de travail

### Rituels

| Quand | Quoi | Durée |
|---|---|---|
| Chaque matin | Debout devant le board : ce que j'ai fini, ce que je prends, ce qui me bloque | 10 min |
| Chaque fin de journée | Chacun pousse sa branche, même inachevée | 5 min |
| Chaque soir | On relance le jeu ensemble et on rejoue le parcours complet | 15 min |

### Git

- Branche `main` : **toujours fonctionnelle**. Personne ne pousse directement dessus.
- Une branche par tâche : `feat/t07-marchand`, `feat/t22-training-fight`, `fix/t12-limite-inventaire`
- Commits en français, à l'impératif, une idée par commit : `ajoute la limite d'inventaire dans addInventory`
- Une Pull Request par tâche, **relue par un autre membre** avant fusion
- Le relecteur vérifie la checklist « Critères de validation » de la spec avant d'approuver

### Définition de « terminé »

Une carte passe en **Terminé** quand, et seulement quand :

1. le code compile sans avertissement ;
2. **tous** les critères de validation de la spec sont cochés ;
3. les cas limites listés dans la spec ont été testés à la main ;
4. un autre membre a relu et fusionné la PR ;
5. aucun message technique ne reste à l'écran.

## 5. Outils

| Outil | Usage |
|---|---|
| **Trello** | Suivi des tâches — une carte par tâche du sujet |
| **Git / GitHub** | Versionnement, branches, PR, relecture |
| **Go** | Langage du projet (imposé par les exemples du sujet) |
| **Discord** | Communication quotidienne, partage d'écran pour le pair programming |

## 6. Risques identifiés et parades

| Risque | Probabilité | Impact | Parade |
|---|---|---|---|
| Deux personnes modifient `Character` en même temps | Élevée | Élevé | La structure est figée dès le vendredi. Toute modification ultérieure passe par Théo et est annoncée sur Discord |
| Conflits Git à répétition | Élevée | Moyen | Un fichier par domaine (voir `03-ARCHITECTURE.md`) : deux personnes ne touchent presque jamais le même fichier |
| Quelqu'un est bloqué et ne le dit pas | Moyenne | Élevé | Le point du matin est obligatoire. Règle : **30 minutes bloqué = on demande** |
| L'inventaire est modifié sans passer par `addInventory` | Moyenne | Élevé | Relecture de PR : le relecteur cherche explicitement ce cas |
| On se perd dans les bonus et le socle n'est pas fini | Moyenne | Très élevé | Jalon mercredi 23 : aucun bonus avant que les 22 tâches soient vertes. Point de contrôle le mardi soir |
| Démo qui plante devant le jury | Faible | Très élevé | Deux répétitions complètes le jeudi après-midi, scénario de démo écrit dans `04-SOUTENANCE.md` |

## 7. Suivi d'avancement

| Partie | Tâches | Fait | Reste |
|---|---|---|---|
| Partie 1 | T1 → T12 | 0 / 12 | 12 |
| Partie 2 | T13 → T18 | 0 / 6 | 6 |
| Partie 3 | T19 → T22 | 0 / 4 | 4 |
| Missions bonus | M1 → M6 | 0 / 6 | 6 |
| **Total** | | **0 / 28** | **28** |

> Tableau à mettre à jour chaque soir. C'est la première chose que regarde un encadrant qui ouvre le dossier `docs/`.
