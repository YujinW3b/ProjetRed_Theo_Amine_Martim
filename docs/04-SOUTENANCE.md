# AUBEFER — Préparation de la soutenance

> Le critère « Savoir présenter un projet » pèse **8 points sur 40** — autant que tout le système d'économie et de craft réunis. C'est le critère le moins cher à réussir et le plus souvent bâclé.

---

## 1. Répartition de la parole

Chacun présente **ce qu'il a écrit**. C'est la règle, et c'est ce qui rend l'oral facile : personne n'a à expliquer le code d'un autre.

| Bloc | Qui parle | Durée |
|---|---|---|
| Contexte, thème, pourquoi le camp d'entraînement | Théo | 1 min 30 |
| Démo — création du personnage, fiche de recrue, besace, habillage ASCII | **Martim** | 2 min 30 |
| Démo — Régisseur, Forge, achats, fabrication | **Amine** | 2 min 30 |
| Démo — équipement, combat, schéma du gobelin, mort et résurrection | **Théo** | 2 min 30 |
| Missions bonus + les graffitis (Mission 6) | **Martim** | 1 min |
| Organisation de l'équipe, Trello, Git, ce qu'on referait autrement | Théo | 1 min 30 |
| Questions du jury | À trois | — |

**Chacun parle.** Un membre silencieux, c'est un jury qui se demande ce qu'il a fait.

> Martim ouvre et ferme la démonstration : c'est lui qui a fait l'habillage, donc c'est lui qui montre l'écran au moment où le jury le découvre, et lui qui livre la réponse de la Mission 6 à la fin. Les deux moments les plus mémorables de l'oral.

## 2. Le scénario de démo — à répéter deux fois avant le jour J

L'ordre est choisi pour que chaque étape **prépare** la suivante et qu'aucune manipulation ne soit perdue.

1. **Lancement** → la bannière ASCII s'affiche
2. **Création** : taper `Th30!` → refusé par le Sergent · taper `tHEO` → devient `Théo` · choisir **Nain** → 60/120 PV
   > *« Les points de vie de départ, c'est volontairement la moitié : la recrue doit d'abord apprendre à se soigner. »*
3. **Fiche de recrue** → montrer la barre de vie, l'or, la besace 0/10
4. **Régisseur** → prendre la **potion de vie gratuite** · essayer de la reprendre → payante
5. **Besace** → boire la potion → 110/120 · en boire une deuxième → **refusée**, PV déjà pleins
   > *« On plafonne : les PV actuels ne dépassent jamais le maximum, c'est dans le sujet. »*
6. **Régisseur** → acheter 2 Fourrures de Loup + 1 Peau de Troll
7. **Forge** → essayer le **Chapeau** → refusé, matériaux manquants, **or intact** (le montrer)
   > *« Le contrôle se fait avant toute consommation : un échec ne coûte rien au joueur. »*
8. **Forge** → fabriquer la **Tunique** → matériaux consommés, −5 or
9. **Besace** → équiper la Tunique → **PV max 120 → 145**, l'objet quitte la besace
10. **Terrain d'exercice** → combat :
    - tour 1 : attaquer → 5 dégâts au gobelin, le gobelin rend 5
    - tour 3 : **le gobelin frappe à 10** → *« tous les 3 tours, 200 % de son attaque »*
    - ouvrir la besace en combat, boire une potion → *« et ça consomme le tour, sinon le combat n'a plus d'enjeu »*
    - finir le gobelin → récapitulatif de fin de combat
11. **Poison** → acheter et boire une potion de poison → 3 ticks visibles à l'écran
12. Si les PV tombent à 0 → **mort puis résurrection à 50 %** en direct
13. **Graffitis de la palissade** → ABBA et Steven Spielberg
14. **Quitter**

**Durée cible : 7 minutes.** Chronométrez-la.

## 3. Les phrases qui font gagner des points

À placer naturellement pendant la démo — ce sont les justifications que le jury attend :

- « On a gardé **tous** les noms imposés par le sujet dans le code ; le thème vit dans les textes affichés. On voulait les points de thème sans perdre les points de conformité. »
- « `addInventory` et `removeInventory` sont les **seuls** points d'entrée de l'inventaire. C'est pour ça que la limite d'emplacements ne peut pas être contournée, d'où qu'on ajoute un objet. »
- « À la Forge, on contrôle **avant** de consommer : matériaux, puis or, puis place. Un échec ne coûte jamais rien au joueur. »
- « Quand on retire un équipement, les PV max baissent — donc on replafonne les PV actuels. Sinon on se retrouve à 110 sur 85. »
- « Ouvrir l'inventaire en combat consomme le tour. C'est une décision de game design, pas un oubli : sinon on peut boire dix potions sans que le gobelin bouge. »
- « Chaque message d'erreur passe par le Sergent. Il n'y a aucun message technique dans le jeu. »

## 4. Questions probables du jury et réponses préparées

| Question | Réponse |
|---|---|
| « Pourquoi ce type pour l'inventaire ? » | Une collection ordonnée qui accepte les doublons : on peut avoir deux Fourrures de Loup, ce qu'un ensemble de valeurs uniques interdirait |
| « Que se passe-t-il si la besace est pleine au moment d'un achat ? » | L'achat est refusé **avant** le débit. On peut le montrer tout de suite |
| « Et si on équipe une pièce alors que la besace est pleine ? » | Refusé, sinon l'ancienne pièce serait détruite en revenant dans la besace |
| « Pourquoi le gobelin frappe plus fort au tour 3 ? » | C'est le schéma imposé par le sujet : 200 % tous les 3 tours. On affiche le numéro de tour pour que le joueur puisse l'anticiper — c'est tout l'intérêt d'un terrain d'exercice |
| « Comment vous êtes-vous répartis le travail ? » | Un domaine = un fichier = un responsable. Le noyau a été fait à trois en pair programming parce que tout en dépend. Le reste en parallèle, avec Trello et une PR relue par carte |
| « Quel a été votre plus gros blocage ? » | Les PV max variables avec l'équipement : trois endroits recalculaient la valeur. On a tout centralisé dans une seule fonction |
| « Qu'est-ce que vous referiez autrement ? » | On figerait la structure `Character` encore plus tôt : on a ajouté `skill`, l'or et l'équipement au fil des tâches, et chaque ajout a touché plusieurs fichiers |
| « Les deux artistes cachés ? » | ABBA pour la partie 2 — six titres de chansons. Steven Spielberg pour la partie 3 — *Fighter Squad*, *Duel*, *A.I.*, *Ready Player One* |

## 5. Check technique avant de passer

- [ ] Le jeu compile et se lance **sur la machine de démo**, testée la veille
- [ ] Terminal agrandi, police lisible de loin, fond sombre
- [ ] Les couleurs ANSI s'affichent bien sur ce terminal-là
- [ ] Le dépôt Git est à jour et le lien déposé sur Moodle
- [ ] Le board Trello est à jour : rien qui traîne en « En cours »
- [ ] Une sauvegarde du projet sur une clé USB **et** en ligne
- [ ] Le scénario de démo a été joué deux fois en entier, chronomètre en main

## 6. Ce qu'il ne faut pas faire

- Lire ses notes. Le scénario est mémorisé, pas récité.
- Découvrir un bug en direct sans réagir. Si ça arrive : **le nommer, expliquer la cause probable, continuer**. Un jury pardonne un bug expliqué, jamais un bug ignoré.
- Faire défiler le code pendant dix minutes. Le jury veut voir le **jeu** ; le code ne s'ouvre que si on pose une question dessus.
- Répondre « je ne sais pas, c'est lui qui l'a fait ». Chacun doit pouvoir expliquer, au moins dans les grandes lignes, ce que font les deux autres.
