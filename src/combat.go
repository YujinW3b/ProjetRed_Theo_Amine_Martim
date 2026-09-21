package main

// THEO - T20 goblinPattern, T21 characterTurn, T22 trainingFight, M1, M2.
// Le compteur de tours vit ici et est PASSE a goblinPattern.

func goblinDamage(attack int, turn int) int {
	if turn%3 == 0 { // tours 3, 6, 9... le gobelin cogne double
		return attack * 2
	}
	return attack
}
