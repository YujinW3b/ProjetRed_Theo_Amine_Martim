package src

import "fmt"

// THEO - T20 goblinPattern, T21 characterTurn, T22 trainingFight, M1, M2.
// Le compteur de tours vit ici et est PASSE a goblinPattern.

func goblinDamage(attack int, turn int) int {
	if turn%3 == 0 { // tout les 3 tour latack fais x2 donc on regarde divisible par 3
		return attack * 2
	}
	return attack // sinon sa attaque normal
}

func askCombatAction() string {
	for { // je redemande tant que c'est pas 1 ou 2
		fmt.Println()
		fmt.Println("   --- Ton tour ---")
		fmt.Println("   1. Attaquer")
		fmt.Println("   2. Besace")
		fmt.Print("   Ton choix : ")

		choix := lireChoix()

		switch choix {
		case "1":
			return "attack"
		case "2":
			return "inventory"
		default:
			fmt.Println("   Parle plus clairement, recrue.")
		}
	}
}

func goblinPattern(goblin Monster, turn int) {
	if turn%3 == 0 { // le critere T20 veut un message pour attaque renforcee tt les 3 tour 
		fmt.Println("   Le gobelin prend son elan... attaque renforcee !")
	}

	damage := goblinDamage(goblin.Attack, turn) // init damage avec attack et nombre tour
	joueur.Pv = joueur.Pv - damage // pv joueur apres attaque = pv joeuur - damage

	if joueur.Pv < 0 { // jamais de PV negatifs a l'ecran on met a 0 si inferieur a 0
		joueur.Pv = 0
	}

	fmt.Println("  ", goblin.Name, "inflige à", joueur.Name, damage, "de dégâts") 
	fmt.Println("  ", joueur.Name, "PV :", joueur.Pv, "/", joueur.Pvmax)
}