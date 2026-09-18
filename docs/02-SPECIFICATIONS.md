# AUBEFER — Spécifications fonctionnelles

> Une entrée par tâche du sujet. Pour chacune : ce que le sujet demande, comment on l'habille dans le camp, et **les critères de validation** — c'est-à-dire ce qu'on doit pouvoir montrer au jury pour dire « c'est acquis ».
>
> Aucune ligne de code ici : c'est un document de spécification, le code vit dans `src/`.

**Légende des responsables** — 🟢 Théo (menus, création, équipement, combat) · 🟡 Amine (inventaire, objets, marchand, forge) · 🟠 Martim (structures, affichage, habillage ASCII)

---

# PARTIE 1 — Création du personnage & fonctionnalités de base

## T1 — Structure `Character` 🟠

**Demandé** : une structure `Character` avec nom, classe, niveau, PV max, PV actuels, inventaire. Les types sont à déduire.

**Dans le camp** : la fiche de recrue que le Sergent remplit à l'entrée.

**À décider en équipe (à faire avant d'écrire la structure)** :
- le nom et la classe sont du texte ;
- le niveau et les points de vie sont des entiers — et **jamais négatifs** à l'affichage ;
- l'inventaire est une **collection ordonnée de texte** : on doit pouvoir avoir deux fois « Potion de vie » dedans, donc ce n'est pas un ensemble de valeurs uniques.

**Critères de validation**
- [ ] La structure existe dans un fichier dédié, pas dans le `main`
- [ ] Les six attributs du sujet sont présents, ni plus ni moins à ce stade
- [ ] Le choix de chaque type est justifiable à l'oral en une phrase

---

## T2 — `initCharacter` 🟠

**Demandé** : une fonction `initCharacter` qui initialise un personnage. Dans le `main`, créer `c1` : nom = ton nom, classe = Elfe, niveau 1, PV max 100, PV actuels 40, inventaire = 3 potions.

**Dans le camp** : le Sergent inscrit la recrue sur le registre.

> ⚠️ Cette initialisation « en dur » est **provisoire** : la tâche 11 la remplacera par une vraie création interactive. Prévoyez-le dès maintenant pour ne pas avoir à tout réécrire — `initCharacter` doit recevoir ses valeurs **en paramètres**, jamais les décider elle-même.

**Critères de validation**
- [ ] `initCharacter` renvoie un `Character` complet
- [ ] `c1` est créé avec exactement les valeurs du sujet
- [ ] Les trois potions sont bien trois entrées distinctes dans l'inventaire

---

## T3 — `displayInfo` 🟠

**Demandé** : afficher les informations du personnage.

**Dans le camp** : « Ta fiche de recrue », option 1 du feu de camp.

**Rendu attendu** : un encadré, pas une suite de `println`. Nom, lignage, niveau, PV actuels **sur** PV max avec barre de vie, or, sorts connus, équipement porté, remplissage de la besace (`4 / 10`).

**Critères de validation**
- [ ] Affiche les six attributs de la structure
- [ ] Les PV sont affichés au format `actuels / max`
- [ ] L'écran est encadré et de largeur fixe
- [ ] La fonction ne modifie rien : elle lit et affiche, c'est tout

---

## T4 — `accessInventory` 🟠

**Demandé** : afficher tous les objets de l'inventaire, **utilisables par la suite**.

**Dans le camp** : « Ta besace », option 2.

**Point important** : « utilisables par la suite » veut dire que cette fonction ne doit pas seulement lister — elle doit **numéroter** les objets pour qu'on puisse en choisir un. C'est la base des tâches 5, 9, 10 et 17.

**Critères de validation**
- [ ] Chaque objet est numéroté
- [ ] Une besace vide affiche un message du Sergent, pas une liste vide
- [ ] Un choix « Retour » ramène au feu de camp
- [ ] Le compteur `x / capacité` est visible

---

## T5 — `takePot` 🟡

**Demandé** : utiliser une potion depuis la besace. Elle est **consommée** (retirée de l'inventaire), rend **50 PV**, puis on affiche PV actuels / PV max. Les PV actuels ne peuvent **jamais** dépasser les PV max.

**Dans le camp** : « Tu bois. Le goût est infect. »

**Cas limites à traiter explicitement** :
- pas de potion dans la besace → message du Sergent, aucun PV rendu
- PV déjà au maximum → on peut refuser la potion plutôt que de la gâcher (choix d'équipe : **on refuse**, et on le dit)
- soin qui dépasserait le maximum → on plafonne, on ne dépasse pas

**Critères de validation**
- [ ] La potion disparaît de l'inventaire après usage
- [ ] +50 PV appliqués
- [ ] Plafonnement aux PV max démontré en live (par exemple 90/100 + potion → 100/100)
- [ ] Les trois cas limites ci-dessus produisent un message différent

---

## T6 — Le menu 🟢

**Demandé** : un menu avec « Afficher les informations », « Accéder à l'inventaire », « Quitter », construit avec un **switch case**, lisant l'entrée utilisateur, relié aux fonctions précédentes, avec un choix « Retour » pour naviguer.

**Dans le camp** : le feu de camp. C'est le cœur du jeu — tout y revient.

**Architecture à respecter** : une **boucle infinie** qui affiche le menu, lit un choix, exécute, puis recommence. Seul « Quitter » sort de la boucle. Les sous-menus (besace, marchand, forge) sont des boucles identiques avec leur propre « Retour ».

**Critères de validation**
- [ ] Un `switch` est bien utilisé (exigence explicite du sujet)
- [ ] Une saisie invalide ne fait pas planter le jeu et ne fait pas sortir de la boucle
- [ ] Une saisie vide, une lettre à la place d'un chiffre, un nombre hors bornes : trois cas testés
- [ ] Chaque sous-menu a son « Retour »
- [ ] « Quitter » termine proprement le programme

---

## T7 — Le Marchand 🟡

**Demandé** : ajouter un choix « Marchand » au menu. Le marchand vend 1 Potion de vie **gratuitement**. L'objet choisi est ajouté à l'inventaire et son nom est affiché après l'achat. Astuce du sujet : créer `addInventory` et `removeInventory`.

**Dans le camp** : la tente du Régisseur.

> **Faites vraiment `addInventory` / `removeInventory`.** Tout le reste du projet (marchand, forge, potions, équipement, sorts) passe par ces deux fonctions. Si elles n'existent pas, la limite d'inventaire de la tâche 12 devra être recopiée à huit endroits — et c'est là qu'on perd des points de structure.

**Critères de validation**
- [ ] « Marchand » est dans le menu principal
- [ ] `addInventory` et `removeInventory` existent et sont les **seuls** points d'entrée pour modifier l'inventaire
- [ ] Le nom de l'objet acheté est affiché après l'achat
- [ ] Le sous-menu du marchand a un « Retour »

---

## T8 — `isDead` 🟠

**Demandé** : vérifier si le joueur est à 0 PV ; si oui, il meurt puis **ressuscite à 50 % de ses PV max**.

**Dans le camp** : « Debout, recrue. On ne meurt pas à l'exercice. »

**Attention au calcul** : 50 % des PV **maximum**, pas des PV actuels. Et les PV max changent avec l'équipement (tâche 17) — donc `isDead` doit lire la valeur au moment de l'appel, pas une constante.

**Critères de validation**
- [ ] Le test se déclenche à **0 PV ou moins** (le poison et le gobelin peuvent faire descendre en dessous de 0)
- [ ] Résurrection à exactement la moitié des PV max courants
- [ ] Un message de mort puis un message de résurrection sont affichés
- [ ] Appelée après chaque source de dégâts : poison, attaque du gobelin

---

## T9 — `poisonPot` 🟡

**Demandé** : 10 dégâts par seconde pendant 3 secondes, avec affichage des PV actuels / max **à chaque seconde**. Utiliser la bibliothèque `time` pour temporiser. Ajouter la Potion de poison au marchand.

**Dans le camp** : une fiole verte que le Régisseur vend sans commentaire.

**Points de vigilance** :
- l'affichage doit se faire **entre** chaque seconde, pas tout d'un coup à la fin ;
- 3 × 10 = 30 dégâts peuvent tuer une recrue : `isDead` doit être appelé **après** le dernier tick, et le poison ne doit pas continuer après une résurrection.

**Critères de validation**
- [ ] Trois affichages espacés d'une seconde, visibles à l'œil nu pendant la démo
- [ ] 30 dégâts au total
- [ ] Enchaînement poison → mort → résurrection démontré
- [ ] La potion de poison est achetable chez le Régisseur

---

## T10 — `spellBook` (Wingardium leviosa) 🟡

**Demandé** : ajouter un attribut `skill` (liste de sorts) à `Character`, modifier `initCharacter`, donner le sort de base « Coup de poing ». Créer `spellBook` qui ajoute « Boule de Feu » — **un même sort ne peut être appris qu'une fois**. Ajouter « Livre de Sort : Boule de Feu » chez le marchand ; à l'usage depuis la besace, il appelle `spellBook`.

**Dans le camp** : un grimoire corné que le Régisseur a récupéré sur un mort.

**Question à trancher en équipe** : si le joueur utilise un deuxième Livre de Sort alors qu'il connaît déjà Boule de Feu, le livre est-il consommé ? **Décision : non.** Le Sergent dit « Tu sais déjà faire ça », le livre reste dans la besace, l'or n'est pas perdu deux fois. C'est plus juste pour le joueur, et c'est une décision à savoir défendre à l'oral.

**Critères de validation**
- [ ] `skill` est dans la structure et contient « Coup de poing » dès la création
- [ ] L'achat puis l'utilisation du livre ajoute « Boule de Feu »
- [ ] Le double apprentissage est bloqué avec un message clair
- [ ] Les sorts connus apparaissent dans `displayInfo`

---

## T11 — `characterCreation` 🟢

**Demandé** : le joueur crée lui-même son personnage. Nom : **lettres uniquement**, reformaté en Majuscule + minuscules. Classe au choix parmi Humain / Elfe / Nain, avec 100 / 80 / 120 PV max. PV de départ = 50 % des PV max. Niveau 1. Sort : Coup de Poing. **Cette tâche remplace l'initialisation en dur de la tâche 2.**

**Dans le camp** : la scène d'ouverture. Le Sergent demande le nom, puis le lignage.

**Validation du nom — c'est le piège de la tâche** : il faut refuser les chiffres, les espaces, les caractères spéciaux, et la chaîne vide, **puis redemander** au lieu de planter ou d'accepter. Le reformatage doit marcher quelle que soit la saisie : `tHÉO`, `THEO`, `theo` donnent tous `Théo`.

**Critères de validation**
- [ ] `aze123` refusé, `Théo!` refusé, saisie vide refusée, et le jeu redemande à chaque fois
- [ ] `tHEO` → `Theo` démontré en live
- [ ] Les trois lignages donnent bien 100 / 80 / 120 PV max
- [ ] PV de départ = exactement la moitié
- [ ] L'appel en dur de la tâche 2 a été supprimé du `main`

---

## T12 — Limite d'inventaire 🟡

**Demandé** : une fonction empêchant de dépasser **10 objets**. Ce contrôle est fait à chaque ajout.

**Dans le camp** : « Ta besace déborde, recrue. »

**Où le brancher** : dans `addInventory`, et **uniquement là**. Si le contrôle est branché ailleurs, c'est qu'`addInventory` est contourné quelque part — à corriger.

**Critères de validation**
- [ ] Le 11ᵉ objet est refusé, quelle que soit sa provenance (marchand, forge, déséquipement)
- [ ] L'or n'est **pas** débité quand l'ajout est refusé — à tester explicitement
- [ ] Message du Sergent affiché
- [ ] La capacité est une variable, pas le nombre 10 écrit en dur (la tâche 18 va la faire varier)

---

# PARTIE 2 — Économie & fabrication d'équipement

## T13 — L'or 🟠

**Demandé** : un attribut argent dans `Character`, 100 pièces d'or au départ.

**Critères de validation**
- [ ] Attribut présent, initialisé à 100 dans `characterCreation`
- [ ] Visible dans `displayInfo`
- [ ] Ne peut jamais devenir négatif

---

## T14 — Les prix (Two for the Price of One) 🟡

**Demandé** : Potion de vie 3 · Potion de poison 6 · Livre de Sort : Boule de Feu 25. Ajouter au marchand : Fourrure de Loup 4 · Peau de Troll 7 · Cuir de Sanglier 3 · Plume de Corbeau 1. L'objet choisi est ajouté à l'inventaire, le prix est déduit de la bourse.

**Dans le camp** : le Régisseur passe en mode commerce — la dotation gratuite, c'était pour la première potion.

**Ordre des contrôles — à faire dans cet ordre exact** :
1. le joueur a-t-il assez d'or ? sinon → message, on s'arrête
2. la besace a-t-elle de la place ? sinon → message, on s'arrête, **sans débiter**
3. on débite, on ajoute, on confirme

**Critères de validation**
- [ ] Les sept prix sont exacts
- [ ] Achat avec 0 or refusé
- [ ] Achat avec besace pleine refusé **et or intact** (le test qui fait perdre des points s'il est oublié)
- [ ] Le solde restant est affiché après chaque achat

---

## T15 — La Forge (Gimme! Gimme! Gimme!) 🟡

**Demandé** : un choix « Forgeron » au menu principal, ouvrant un sous-menu avec Chapeau / Tunique / Bottes de l'aventurier. Fabriquer coûte **5 pièces d'or** et consomme les matériaux :

| Équipement | Matériaux consommés |
|---|---|
| Chapeau de l'aventurier | 1 Plume de Corbeau + 1 Cuir de Sanglier |
| Tunique de l'aventurier | 2 Fourrures de Loup + 1 Peau de Troll |
| Bottes de l'aventurier | 1 Fourrure de Loup + 1 Cuir de Sanglier |

Messages d'erreur **personnalisés** exigés par le sujet : matériaux manquants, or manquant, plus de place.

**Ordre des contrôles** : matériaux → or → place → puis seulement consommer, débiter, ajouter. **Rien ne doit être retiré si un seul contrôle échoue.** C'est le bug classique : les matériaux disparaissent puis l'ajout échoue faute de place.

**Critères de validation**
- [ ] Les trois recettes sont exactes, y compris les **2** Fourrures de Loup de la Tunique
- [ ] Les matériaux disparaissent de la besace à la fabrication
- [ ] Les trois messages d'erreur sont distincts et incarnés par le Sergent
- [ ] Un échec ne consomme **ni** matériaux **ni** or — démontré en live
- [ ] La recette est affichée à côté de chaque objet dans le menu de la Forge

---

## T16 — Structure `Equipment` (I saw it in the mirror) 🟠

**Demandé** : une structure `Equipment` avec un emplacement tête, un torse, un pieds. Ajouter à `Character` un attribut équipement basé sur cette structure.

**Critères de validation**
- [ ] Structure créée avec exactement trois emplacements
- [ ] `Character` la contient
- [ ] Un emplacement vide est représenté de façon lisible (« — » à l'affichage, pas une chaîne vide)
- [ ] `displayInfo` affiche les trois emplacements

---

## T17 — Équiper (Mamma Mia) 🟢

**Demandé** : rendre les équipements utilisables depuis la besace, les ranger au bon emplacement d'`Equipment`, **les faire disparaître de l'inventaire**. Bonus de PV max : Chapeau +10, Tunique +25, Bottes +15. Si l'emplacement est déjà occupé, le nouvel équipement remplace l'ancien **qui retourne dans la besace**.

**Le point délicat — les PV max qui bougent** : quand on retire un équipement, les PV max redescendent. Si les PV actuels étaient au-dessus du nouveau maximum, il faut les **replafonner**. Exemple : 110/110 avec la Tunique, on l'enlève → 85/85, pas 110/85.

**Critères de validation**
- [ ] Équiper retire l'objet de la besace et remplit l'emplacement
- [ ] Les trois bonus sont exacts
- [ ] Le remplacement renvoie l'ancienne pièce dans la besace
- [ ] Remplacement refusé si la besace est pleine (sinon l'ancien équipement est perdu)
- [ ] Les PV actuels sont replafonnés quand les PV max baissent
- [ ] Une pièce équipée n'apparaît plus dans la besace

---

## T18 — `upgradeInventorySlot` (On and on and on) 🟡

**Demandé** : `upgradeInventorySlot` ajoute **+10** emplacements, utilisable **3 fois maximum**. Ajouter « Augmentation d'inventaire » chez le marchand pour **30 pièces d'or**.

**Dans le camp** : la Besace renforcée. Le Régisseur en a quatre en stock et n'en vendra jamais une quatrième à la même recrue.

**À gérer** : un compteur d'améliorations déjà utilisées, stocké dans `Character`. Au-delà de 3 → message, et **l'or n'est pas débité**.

**Critères de validation**
- [ ] 10 → 20 → 30 → 40 emplacements
- [ ] La 4ᵉ tentative est refusée avec message, or intact
- [ ] La nouvelle capacité est visible dans `displayInfo`
- [ ] La limite de la tâche 12 utilise bien la capacité courante

---

# PARTIE 3 — Combat au tour par tour

## T19 — Structure `Monster` + `initGoblin` (La Chose) 🟠

**Demandé** : structure `Monster` avec nom, PV max, PV actuels, points d'attaque. Fonction `initGoblin` : « Gobelin d'entrainement », 40 PV max, PV actuels = PV max, attaque = 5.

**Critères de validation**
- [ ] Structure créée avec les quatre attributs
- [ ] `initGoblin` renvoie un gobelin conforme
- [ ] Un **nouveau** gobelin est créé à chaque combat (sinon le deuxième combat démarre avec un gobelin déjà blessé)

---

## T20 — `goblinPattern` (A.I.) 🟢

**Demandé** : chaque tour, le gobelin inflige 100 % de son attaque ; **tous les 3 tours**, il inflige 200 %. Afficher le nom de l'attaquant, le nom de l'attaqué et les dégâts. Puis afficher PV actuels / PV max de l'attaqué.

**Format imposé** : `Gobelin d'entrainement inflige à Théo 5 de dégâts`

**Le piège du « tous les 3 tours »** : le tour 3 est un tour renforcé, le tour 6 aussi, le tour 9 aussi. Le tour 1 ne l'est pas. Le compteur de tours vient de `trainingFight` (T22), il n'est pas géré par le gobelin lui-même : `goblinPattern` doit **recevoir le numéro de tour**.

**Critères de validation**
- [ ] Tours 1, 2 → 5 dégâts · tour 3 → 10 dégâts · tours 4, 5 → 5 · tour 6 → 10
- [ ] Le message annonce une attaque renforcée les tours multiples de 3
- [ ] `isDead` est appelé après les dégâts
- [ ] Le format d'affichage imposé est respecté au mot près

---

## T21 — `characterTurn` (Ready Player One) 🟢

**Demandé** : le tour du joueur, avec un menu : Attaquer / Inventaire. « Attaquer » utilise « Attaque basique » et inflige **5 dégâts**, en affichant le nom de l'attaque, les dégâts, puis les PV restants de l'adversaire. « Inventaire » affiche les objets ; l'objet choisi est utilisé et son effet s'applique. Dans les deux cas, c'est ensuite au monstre de jouer.

**Règle à assumer** : ouvrir l'inventaire **consomme le tour**. Sinon le joueur peut boire dix potions sans que le gobelin bouge, et le combat n'a plus d'enjeu. À dire au jury, c'est une vraie décision de game design.

**Critères de validation**
- [ ] Le menu de combat s'affiche à chaque tour du joueur
- [ ] L'attaque basique inflige 5 dégâts et l'affichage est complet
- [ ] Utiliser une potion en combat la consomme et applique l'effet
- [ ] Un objet inutilisable en combat (un matériau) est refusé sans perdre le tour
- [ ] Après l'action, le gobelin joue

---

## T22 — `trainingFight` (Fighter Squad / Duel) 🟢

**Demandé** : lancer le combat d'entraînement contre le monstre, au tour par tour. Une variable indique **à quel tour de combat on se situe**. `trainingFight` appelle `characterTurn` puis `goblinPattern` l'une après l'autre. Le numéro de tour est affiché au début de chaque tour. Une option « Entrainement » est ajoutée au menu principal. Si les PV du joueur **ou** du monstre tombent à 0 ou moins, le combat s'arrête et le joueur revient au menu de départ.

**Dans le camp** : le terrain d'exercice, la fosse au gobelin.

**Cas à traiter** : le joueur peut mourir en plein combat — `isDead` le ressuscite à 50 %. **Décision d'équipe : la résurrection met fin au combat** et renvoie au feu de camp, sinon le combat ne se termine jamais. Le Sergent le dit : « Ça suffit pour ce soir. »

**Critères de validation**
- [ ] « Entrainement » est dans le menu principal
- [ ] Le numéro de tour s'affiche en début de chaque tour et s'incrémente
- [ ] Le combat s'arrête à 0 PV ou moins, des deux côtés
- [ ] Retour propre au menu principal dans les deux issues
- [ ] Un deuxième combat démarre avec un gobelin neuf à 40/40
- [ ] Récapitulatif de fin de combat affiché (initiative personnelle)

---

# MISSIONS BONUS

## M1 — Initiative 🟢

Ajouter un attribut initiative au personnage **et** au gobelin ; le tour de jeu commence par celui qui a la plus grande initiative.

- [ ] Attribut présent des deux côtés
- [ ] L'ordre d'appel dans `trainingFight` s'inverse selon l'initiative
- [ ] Égalité gérée (décision d'équipe : le joueur commence)
- [ ] Le lignage influence l'initiative (Elfe le plus rapide) — cohérence avec le thème

## M2 — Expérience 🟢

Points d'expérience propres à chaque monstre, donnés en fin de combat. Attributs expérience actuelle / expérience max ; au maximum atteint, le niveau augmente. **L'excédent est reporté** sur le niveau suivant, et le palier augmente à chaque niveau. Tout gain de statistique au passage de niveau compte en bonus.

- [ ] Le gobelin donne son XP à la victoire
- [ ] L'excédent est reporté, pas perdu — à démontrer avec un gros gain
- [ ] Le palier augmente à chaque niveau (règle simple et annoncée : `palier × 1,5`)
- [ ] Plusieurs niveaux peuvent être gagnés d'un coup
- [ ] Le passage de niveau donne +10 PV max et rend le joueur à plein — et c'est affiché

## M3 — Combat magique 🟡

Utiliser les sorts en combat : Coup de poing **8 dégâts**, Boule de Feu **18 dégâts**. Possibilité d'ajouter soins, bonus ou malus.

- [ ] Un choix « Sorts » dans le menu de combat, listant uniquement les sorts connus
- [ ] Les dégâts sont exacts
- [ ] Boule de Feu absente tant que le livre n'a pas été lu
- [ ] Sort ajouté maison : « Souffle du Sergent » (soigne 20 PV) — initiative personnelle

## M4 — Mana 🟡

Attributs mana / mana max. Les sorts consomment du mana, et sont **impossibles** si le mana est insuffisant. Possibilité d'ajouter une potion de mana chez le marchand.

- [ ] Coût par sort défini et affiché à côté du sort
- [ ] Sort bloqué avec message si mana insuffisant — et le tour n'est pas perdu
- [ ] Mana affiché dans `displayInfo` et dans l'interface de combat
- [ ] Potion de mana vendue par le Régisseur
- [ ] L'attaque basique ne coûte pas de mana (sinon on peut être bloqué sans action possible)

## M5 — Améliorer le jeu 🟠

Enrichir librement le contenu. Ce qu'on livre :

- [ ] Bannière ASCII au lancement
- [ ] Barres de vie et écrans encadrés partout
- [ ] Couleurs ANSI (rouge dégâts, vert soins, jaune or)
- [ ] Nettoyage de l'écran entre deux écrans
- [ ] Récapitulatif de fin de combat
- [ ] Option « Fuir » en combat
- [ ] Zéro message technique : tout passe par le Sergent

## M6 — Qui sont-ils ? 🟠

Deux artistes sont cachés dans les titres des parties 2 et 3 du sujet. Ajouter une option de menu affichant leur nom.

**Réponse — trouvée et vérifiée** :

| Partie | Artiste | Indices dans les titres des tâches |
|---|---|---|
| **Partie 2** | **ABBA** | *Money, Money, Money* (T13) · *Two for the Price of One* (T14) · *Gimme! Gimme! Gimme!* (T15) · *I Saw It in the Mirror* (T16) · *Mamma Mia* (T17) · *On and On and On* (T18) — six chansons du groupe |
| **Partie 3** | **Steven Spielberg** | *Fighter Squad* (son court-métrage amateur de 1961) · *Duel* (1971) · *A.I. Intelligence Artificielle* (2001) · *Ready Player One* (2018) |

**Dans le camp** : « Les graffitis de la palissade ». Des recrues d'autres nuits ont gravé des noms dans le bois. On les lit, on découvre les deux artistes. L'easter egg devient un lieu du camp au lieu d'un `println` collé à la fin — c'est ce qui fait la différence sur le critère de cohérence thématique.

- [ ] Option présente dans le menu principal
- [ ] Les deux noms sont affichés
- [ ] Les titres de tâches qui servent d'indices sont cités (montre qu'on a compris, pas deviné)
