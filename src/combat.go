package src

import "fmt"

const (
	coutManaCoupDePoing = 5
	coutManaBouleDeFeu  = 10
)

func goblinDamage(attack int, turn int) int {
	if turn%3 == 0 {
		return attack * 2
	}
	return attack
}

func askCombatAction() string {
	for {
		fmt.Println()
		fmt.Println("   --- Ton tour ---")
		fmt.Println("   1. Attaquer")
		fmt.Println("   2. Besace")
		fmt.Println("   3. Sorts")
		fmt.Print("   Ton choix : ")

		choix := lireChoix()

		switch choix {
		case "1":
			return "attack"
		case "2":
			return "inventory"
		case "3":
			return "spell"
		default:
			sergent("   Parle plus clairement, recrue.")
		}
	}
}

func coutManaAffiche(sort string) string {
	switch sort {
	case spellCoupDePoing:
		return fmt.Sprintf(" (%d mana)", coutManaCoupDePoing)
	case spellBouleDeFeu:
		return fmt.Sprintf(" (%d mana)", coutManaBouleDeFeu)
	default:
		return ""
	}
}

// castSpell tente de lancer un sort. Retourne true si un sort a
// effectivement été lancé (mana consommé / effet appliqué),
// false si l'action a échoué ou a été annulée : dans ce cas
// le tour du joueur n'est PAS consommé.
func castSpell(goblin *Monster) bool {
	if len(joueur.Skill) == 0 {
		fmt.Println("   Sergent : « T'as pas appris le moindre sort, recrue. »")
		return false
	}

	fmt.Println()
	fmt.Println("   --- Sorts connus ---")
	for i, spell := range joueur.Skill {
		fmt.Printf("   %d. %s%s\n", i+1, spell, coutManaAffiche(spell))
	}
	fmt.Print("   Ton choix : ")

	choix := lireChoix()
	var index int
	_, err := fmt.Sscanf(choix, "%d", &index)
	if err != nil || index < 1 || index > len(joueur.Skill) {
		fmt.Println("   Sergent : « Parle plus clairement, recrue. »")
		return false
	}

	sort := joueur.Skill[index-1]
	switch sort {
	case spellCoupDePoing:
		if joueur.Mana < coutManaCoupDePoing {
			fmt.Println("   Sergent : « Pas assez de mana, recrue ! »")
			return false
		}
		joueur.Mana -= coutManaCoupDePoing
		goblin.Pv -= 8
		if goblin.Pv < 0 {
			goblin.Pv = 0
		}
		fmt.Println("  ", joueur.Name, "utilise Coup de poing et inflige 8 dégâts à", goblin.Name)
		fmt.Println("  ", goblin.Name, "PV :", goblin.Pv, "/", goblin.Pvmax)
		return true

	case spellBouleDeFeu:
		if joueur.Mana < coutManaBouleDeFeu {
			fmt.Println("   Sergent : « Pas assez de mana, recrue ! »")
			return false
		}
		joueur.Mana -= coutManaBouleDeFeu
		goblin.Pv -= 18
		if goblin.Pv < 0 {
			goblin.Pv = 0
		}
		fmt.Println("  ", joueur.Name, "utilise Boule de Feu et inflige 18 dégâts à", goblin.Name)
		fmt.Println("  ", goblin.Name, "PV :", goblin.Pv, "/", goblin.Pvmax)
		return true

	case spellSouffleDuSergent:
		joueur.Pv += 20
		if joueur.Pv > joueur.Pvmax {
			joueur.Pv = joueur.Pvmax
		}
		fmt.Println("   Sergent : « DEBOUT, RECRUE ! »")
		fmt.Println("  ", joueur.Name, "utilise Souffle du Sergent et récupère 20 PV")
		fmt.Println("  ", joueur.Name, "PV :", joueur.Pv, "/", joueur.Pvmax)
		return true

	default:
		fmt.Println("   Sergent : « Ce sort ne fait rien pour l'instant, recrue. »")
		return false
	}
}

func goblinPattern(goblin Monster, turn int) bool {
	if turn%3 == 0 {
		typewriter(cRouge + "   Le gobelin prend son elan... attaque renforcee !" + cReset)
	}

	damage := goblinDamage(goblin.Attack, turn)
	joueur.Pv = joueur.Pv - damage

	if joueur.Pv < 0 {
		joueur.Pv = 0
	}

	fmt.Println("  ", goblin.Name, "inflige à", joueur.Name, damage, "de dégâts")
	fmt.Println("  ", joueur.Name, "PV :", joueur.Pv, "/", joueur.Pvmax)

	return joueur.isDead() // true si le joueur est tombe a 0 (il a ete ressuscite)
}

// characterTurn simule le tour du joueur. Retourne true si le tour
// a réellement été consommé (une action a eu lieu), false si le
// joueur doit reprendre la main immédiatement (sort annulé/bloqué).
func characterTurn(goblin *Monster) bool {
	action := askCombatAction()

	switch action {
	case "attack":
		goblin.Pv = goblin.Pv - 5
		if goblin.Pv < 0 {
			goblin.Pv = 0
		}
		fmt.Println()
		fmt.Println(cJaune + "     o==[]::::::::::::>" + cReset)
		fmt.Println("  ", joueur.Name, "utilise Attaque basique et inflige 5 dégâts à", goblin.Name)
		fmt.Println("  ", goblin.Name, "PV :", goblin.Pv, "/", goblin.Pvmax)
		return true
	case "inventory":
		takePot(&joueur)
		return true
	case "spell":
		return castSpell(goblin)
	}
	return false
}

func trainingFight() {
	goblin := initGoblin() // un gobelin neuf a chaque combat
	goblin.Initiative = 9 // M1 : plus lent que l'Elfe et l'Humain, plus rapide que le Nain
	turn := 1              // init tour a 1 pas 0

	fmt.Println()
	sergent("   Le Sergent ouvre la fosse. Le gobelin d'entrainement t'attend.")
	afficherArt(artGobelin, cVert) // le gobelin en ascii, une seule fois au debut

	for {
		fmt.Println()
		fmt.Println(cGris+"   ===== TOUR", turn, "====="+cReset) // nombre tour ecris

		// le rappel des deux barres de vie et du mana a chaque tour
		fmt.Println("   " + joueur.Name + " " + barreDeVie(joueur.Pv, joueur.Pvmax))
		fmt.Println("   " + goblin.Name + " " + barreDeVie(goblin.Pv, goblin.Pvmax))
		fmt.Printf("   %s Mana : %d/%d\n", joueur.Name, joueur.Mana, joueur.ManaMax)

		actionReussie := characterTurn(&goblin)

		if !actionReussie {
			// Sort annulé ou bloqué (mana insuffisant) : le joueur
			// reprend la main immédiatement, le gobelin ne joue pas.
			continue
		}

		if goblin.Pv <= 0 { // le gobelin est tombe donc victoire on sort
			fmt.Println()
			if goblin.Pv <= 0 { // le gobelin est tombe donc victoire on sort
			joueur.gagnerExperience(goblin)
			fmt.Println()
			typewriter(cJaune + "   Le gobelin s'effondre. Bien joue, recrue." + cReset)
			afficherArt(artGobelinKO, cGris)
			afficherArt(artVictoire, cJaune)
			pause()
			return
			}
			typewriter(cJaune + "   Le gobelin s'effondre. Bien joue, recrue." + cReset)
			afficherArt(artGobelinKO, cGris)
			afficherArt(artVictoire, cJaune)
			pause()
			return
		}

		mort := goblinPattern(goblin, turn)

		if mort { // le joueur est tombe donc le Sergent arrete l'exercice
			fmt.Println()
			typewriter(cRouge + "   Ca suffit pour ce soir, recrue." + cReset)
			afficherArt(artKO, cRouge)
			pause()
			return
		}

		turn = turn + 1 // tour suivant
	}
}
