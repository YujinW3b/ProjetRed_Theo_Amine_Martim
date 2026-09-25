# AUBEFER — Ordre de passage

**Vendredi 25/09, 14h00** · créneau de 5 à 10 min, puis questions
Support : 16 slides. Navigation aux flèches, `O` pour le sommaire.

Deux versions selon le temps qu'on sent :

- **Version courte — 8 min 40.** On saute les slides 4, 5 et 7 (marquées « sautable »). C'est celle qu'on vise.
- **Version complète — 10 min 25.** On ne la fait que si le jury a l'air large. C'est le plafond, pas une cible.

---

## L'ordre

| # | Slide | Qui parle | Temps | |
|---|---|---|---|---|
| 1 | Couverture | **Théo** | 0:25 | |
| 2 | Le thème | **Théo** | 0:45 | |
| 3 | Le parcours du joueur | **Martim** | 0:30 | |
| 4 | Le camp (terminal) | **Martim** | 0:30 | *sautable* |
| 5 | Un tour de combat | **Martim** | 0:35 | *sautable* |
| 6 | Ce qu'on a livré | **Amine** | 0:45 | |
| 7 | L'architecture | **Théo** | 0:40 | *sautable* |
| 8 | **DÉMONSTRATION** | **les trois** | **2:30** | |
| 9 | L'organisation | **Théo** | 0:50 | |
| 10 | Qui a fait quoi | **les trois** | 0:25 | une phrase chacun sur sa colonne |
| 11 | Les difficultés | **Amine** | 0:55 | |
| 12 | Objectifs + ce qu'on a appris | **Martim** | 0:45 | |
| 13 | Ce qu'on ferait autrement | **Martim** | 0:35 | |
| 14 | Merci + questions | **Théo** | 0:15 | |
| 15-16 | Annexes de code | **Théo** | — | seulement si on les demande |

Temps de parole en version courte : Théo ~3 min 30, Martim ~2 min 40, Amine ~2 min 30, démo comprise. Personne n'est spectateur.

---

## La démo — slide 8

C'est 25 points sur 100. On se passe le clavier, on ne le garde pas.

**Théo** — lance `go run main.go`. Écran titre, entrer au camp.
Crée la recrue : taper un nom **en minuscules** (montre la majuscule automatique), puis essayer un nom **avec un chiffre** (montre le refus), puis choisir un lignage.
Montre la fiche `[1]` et la besace `[2]`.

**Amine** — prend le clavier. La tente du Régisseur `[3]` : acheter une potion, montrer l'or qui baisse. Commenter un prix.

**Martim** — prend le clavier. Le terrain d'exercice `[4]` : une attaque, une potion en plein combat, un sort, et finir le gobelin. Montrer l'expérience gagnée à la fin.

**Théo** — la palissade `[5]` en passant, puis `[0]` pour quitter proprement.

### Règles

- On ne clique jamais sur une option qu'on n'a pas répétée.
- **Si ça plante : on ne débugue pas en direct.** « On a une capture, je reprends » → slide suivante. Les slides 4 et 5 rejouent le jeu, elles servent exactement à ça.
- On parle pendant qu'on tape. Sinon il y a des silences.
- Terminal maximisé, police agrandie, jeu déjà compilé **avant** d'entrer dans la salle.

---

## Les questions — qui répond

| Question probable | Qui | La réponse en une phrase |
|---|---|---|
| Pourquoi Go ? | Théo | Imposé par les exemples du sujet. |
| Pourquoi `src/` et un `main.go` à la racine ? | Théo | Le jeu se lance avec `go run main.go`, le code reste rangé ; `main.go` appelle juste `src.Start()`. |
| Pourquoi `Start` avec une majuscule ? | Théo | En Go un identifiant capitalisé est exporté, donc visible hors du package. |
| C'est quoi un pointeur, pourquoi ici ? | Théo | `characterTurn` reçoit `*Monster` : sans l'adresse on taperait sur une copie et les dégâts seraient perdus. |
| Pourquoi une slice et pas une map pour l'inventaire ? | Amine | L'ordre de parcours d'une map est volontairement aléatoire en Go, l'affichage changerait à chaque tour. |
| Objet absent de la besace ? | Amine | `findItemIndex` renvoie `-1`, l'appelant teste `-1` avant d'agir. |
| Comment le gobelin choisit son action ? | Martim | `goblinPattern` dépend du numéro de tour : prévisible, donc testable. |
| Pourquoi `goblinPattern` renvoie un booléen ? | Martim | `isDead` remet le joueur debout, donc tester `Pv <= 0` après coup ne détecterait jamais la mort. |
| Comment vous avez géré les conflits Git ? | Théo | Un fichier par domaine, une branche par tâche, relecture avant fusion. |
| Qui a fait quoi ? | chacun | Chacun répond pour son domaine. **Personne ne bluffe.** |

Si une question dépasse : « je ne sais pas, je regarderais dans la documentation de Go » vaut mieux qu'une invention. Le jury le voit tout de suite.

---

## Avant d'entrer

- [ ] Dépôt accessible au jury (il est privé)
- [ ] Support déposé sur Moodle en PDF
- [ ] Démo répétée deux fois en entier, chronomètre en main
- [ ] Terminal maximisé, police agrandie, jeu compilé
- [ ] Support ouvert dans un onglet, terminal dans l'autre
- [ ] D'accord à trois sur qui revendique l'habillage ASCII
