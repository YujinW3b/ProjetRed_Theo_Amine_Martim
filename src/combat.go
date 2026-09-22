package src

import "fmt"

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
			fmt.Println("   Parle plus clairement, recrue.")
		}
	}
}

func castSpell(goblin *Monster) {
	if len(joueur.Skill) == 0 {
		fmt.Println("   Sergent : « T'as pas appris le moindre sort, recrue. »")
		return
	}

	fmt.Println()
	fmt.Println("   --- Sorts connus ---")
	for i, spell := range joueur.Skill {
		fmt.Printf("   %d. %s\n", i+1, spell)
	}
	fmt.Print("   Ton choix : ")

	choix := lireChoix()
	var index int
	_, err := fmt.Sscanf(choix, "%d", &index)
	if err != nil || index < 1 || index > len(joueur.Skill) {
		fmt.Println("   Sergent : « Parle plus clairement, recrue. »")
		return
	}

	sort := joueur.Skill[index-1]
	switch sort {
	case spellCoupDePoing:
		goblin.Pv -= 8
		if goblin.Pv < 0 {
			goblin.Pv = 0
		}
		fmt.Println("  ", joueur.Name, "utilise Coup de poing et inflige 8 dégâts à", goblin.Name)
		fmt.Println("  ", goblin.Name, "PV :", goblin.Pv, "/", goblin.Pvmax)

	case spellBouleDeFeu:
		goblin.Pv -= 18
		if goblin.Pv < 0 {
			goblin.Pv = 0
		}
		fmt.Println("  ", joueur.Name, "utilise Boule de Feu et inflige 18 dégâts à", goblin.Name)
		fmt.Println("  ", goblin.Name, "PV :", goblin.Pv, "/", goblin.Pvmax)

	case spellSouffleDuSergent:
		joueur.Pv += 20
		if joueur.Pv > joueur.Pvmax {
			joueur.Pv = joueur.Pvmax
		}
		fmt.Println("   Sergent : « DEBOUT, RECRUE ! »")
		fmt.Println("  ", joueur.Name, "utilise Souffle du Sergent et récupère 20 PV")
		fmt.Println("  ", joueur.Name, "PV :", joueur.Pv, "/", joueur.Pvmax)

	default:
		fmt.Println("   Sergent : « Ce sort ne fait rien pour l'instant, recrue. »")
	}
}

func goblinPattern(goblin Monster, turn int) {
	if turn%3 == 0 {
		fmt.Println("   Le gobelin prend son elan... attaque renforcee !")
	}

	damage := goblinDamage(goblin.Attack, turn)
	joueur.Pv = joueur.Pv - damage

	if joueur.Pv < 0 {
		joueur.Pv = 0
	}

	fmt.Println("  ", goblin.Name, "inflige à", joueur.Name, damage, "de dégâts")
	fmt.Println("  ", joueur.Name, "PV :", joueur.Pv, "/", joueur.Pvmax)

	isDead(&joueur)
}

func characterTurn(goblin *Monster) {
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
	case "inventory":
		takePot(&joueur)
	case "spell":
		castSpell(goblin)
	}
}
