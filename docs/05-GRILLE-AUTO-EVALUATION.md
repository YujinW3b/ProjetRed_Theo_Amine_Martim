# AUBEFER — Auto-évaluation face à la grille

> À repasser en entier le vendredi après-midi, à trois, jeu ouvert. Un critère n'est coché que si **quelqu'un l'a vu fonctionner à l'écran**.

La grille du sujet comporte 6 critères pour **40 points**, ramenés sur 20.

---

## 1. Concevoir une structure logique de jeu — **6 points**

| Indicateur de la grille | Ce qu'on livre | Où c'est vérifiable | ✔ |
|---|---|---|---|
| Création de personnage adaptable | `characterCreation` : nom validé et reformaté, 3 lignages avec PV distincts, PV de départ à 50 % | T11 | ☐ |
| Organisation cohérente des menus | Feu de camp + sous-menus, `switch`, « Retour » partout, saisie invalide gérée | T6 | ☐ |
| Interfaces d'information du personnage | `displayInfo` encadré : PV `actuels/max` + barre, or, sorts, équipement, besace `x/capacité` | T3 | ☐ |

**Point d'attention** : « adaptable » signifie que l'on peut ajouter un quatrième lignage sans réécrire le jeu. Vérifier que les valeurs de PV par lignage sont dans une table, pas dans une cascade de conditions.

## 2. Implémenter des systèmes d'économie et de crafting — **8 points**

| Indicateur | Ce qu'on livre | Où | ✔ |
|---|---|---|---|
| Interface du marchand et système de vente | Tente du Régisseur : 8 articles, potion offerte à la 1re visite puis payante, prix exacts, or débité | T7, T14, T18 | ☐ |
| Interface du forgeron et fabrication d'équipement | La Forge : 3 recettes exactes, 5 or, matériaux consommés, 3 messages d'erreur distincts | T15 | ☐ |
| Équipement et amélioration du personnage | Structure `Equipment`, équiper/remplacer, bonus +10/+25/+15, PV max recalculés et replafonnés | T16, T17 | ☐ |

**Le test qui fait la différence** : achat ou fabrication refusé → **l'or et les matériaux sont intacts**. À démontrer explicitement, c'est ce qui sépare « partiellement acquis » d'« acquis ».

## 3. Développer des mécaniques de gameplay — **8 points**

| Indicateur | Ce qu'on livre | Où | ✔ |
|---|---|---|---|
| Système de combat au tour par tour fonctionnel | `trainingFight` : tours numérotés, `characterTurn` puis `goblinPattern`, fin à 0 PV des deux côtés | T22 | ☐ |
| Gestion correcte du déroulement des combats | Schéma du gobelin (200 % tous les 3 tours), inventaire en combat qui consomme le tour, `isDead` + résurrection à 50 %, gobelin neuf à chaque combat | T8, T20, T21 | ☐ |
| Retours utilisateurs pertinents | Chaque action annoncée : attaquant, attaqué, dégâts, PV restants. Barres de vie. Récap de fin de combat. Aucun message technique | T20, T21, M5 | ☐ |

**C'est le critère le plus souvent noté « en cours d'acquisition ».** Ce qui le fait basculer : le déroulement doit être **lisible**. Un jury qui ne comprend pas ce qui vient de se passer à l'écran considère que le combat n'est pas maîtrisé, même s'il fonctionne.

## 4. Adapter une base technique à un univers thématique original — **5 points**

| Indicateur | Ce qu'on livre | ✔ |
|---|---|---|
| Respect des tâches tout en les intégrant dans un thème cohérent | Chaque option du menu est un **lieu du camp** ; chaque contrainte du sujet est justifiée dans le lore (voir le tableau §3 du GDD) | ☐ |
| Concevoir et mettre en place un thème original | Le camp d'Aubefer, le Sergent narrateur, la nuit unique d'entraînement, la fosse au gobelin | ☐ |

**Ce qui se dit à l'oral** : « la potion gratuite, c'est la dotation d'entrée ; la résurrection à 50 %, c'est le Sergent qui te relève ; le gobelin qu'on affronte en boucle, c'est un gobelin capturé gardé pour l'exercice. Aucune mécanique imposée n'a eu besoin d'être tordue. »

## 5. Approfondir un projet par l'initiative personnelle — **5 points**

| Indicateur | Ce qu'on livre | ✔ |
|---|---|---|
| Fonctionnalités non exigées, cohérentes avec le thème | Potion offerte une seule fois (mémorisée), option Fuir, récap de fin de combat, sort maison « Souffle du Sergent », initiative liée au lignage | ☐ |
| Réalisation des tâches bonus (missions) | M1 Initiative · M2 Expérience · M3 Combat magique · M4 Mana · M5 Améliorations · M6 Qui sont-ils ? | ☐ |
| Amélioration de l'expérience utilisateur et des interfaces (ASCII art) | Bannière ASCII, écrans encadrés de largeur fixe, barres de vie, couleurs ANSI, nettoyage d'écran | ☐ |

**L'ASCII art est nommément cité dans la grille.** Ce n'est pas optionnel : c'est un indicateur d'évaluation.

## 6. Savoir présenter un projet — **8 points**

| Indicateur | Ce qu'on livre | ✔ |
|---|---|---|
| Explication structurée et claire du projet | Plan de soutenance écrit, rôles répartis, les trois membres parlent | ☐ |
| Maîtrise de la démonstration CLI | Scénario de démo en 14 étapes, répété deux fois, chronométré à 7 min | ☐ |
| Réponses pertinentes aux questions du jury | 8 questions probables préparées avec leurs réponses | ☐ |

Détail dans `04-SOUTENANCE.md`. **8 points sur 40 se jouent en 10 minutes d'oral** : c'est le meilleur rapport effort/points du projet.

---

## Auto-notation blanche — à faire vendredi après-midi

Chacun note le projet **seul**, sans se concerter, puis on compare. Les écarts désignent exactement ce qu'il reste à faire.

| Critère | Points | Théo | Amine | Martim |
|---|---|---|---|---|
| Structure logique de jeu | /6 | | | |
| Économie et crafting | /8 | | | |
| Mécaniques de gameplay | /8 | | | |
| Univers thématique | /5 | | | |
| Initiative personnelle | /5 | | | |
| Présentation | /8 | | | |
| **Total** | **/40** | | | |
| **Note sur 20** | | | | |

## Rendu — à vérifier avant de déposer sur Moodle

- [ ] Le dépôt s'appelle **`projet-red_AUBEFER`**
- [ ] Il contient un dossier **`src/`** avec le code source
- [ ] Il contient un dossier **`docs/`** avec le document de gestion de projet
- [ ] Il contient un **`README.md`** complet : présentation courte + installation + lancement
- [ ] Le lien du dépôt est déposé sur Moodle **avant la date des encadrants**
- [ ] Les trois membres ont des commits à leur nom dans l'historique
