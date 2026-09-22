package src
//gobelin
func initGoblin() Monster {
	goblin := Monster{}

	goblin.Name = "Gobelin d'entrainement"
	goblin.Pvmax = 40
	goblin.Pv = goblin.Pvmax
	goblin.Attack = 5

	return goblin
}

// MARTIM - T19 : structure Monster + initGoblin.
// Un NOUVEAU gobelin est cree a chaque combat.
