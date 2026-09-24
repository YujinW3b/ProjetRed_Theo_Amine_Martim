package src
//gobelin
func initGoblin() Monster {
	goblin := Monster{}

	goblin.Name = "Gobelin d'entrainement"
	goblin.Pvmax = 40
	goblin.Pv = goblin.Pvmax
	goblin.Attack = 5
	goblin.XPDonnee = 40

	return goblin
}
