package src

import (
	"fmt"
	"time"
)

func isDead(c *Character) bool {
	if c.Pv <= 0 {
		fmt.Printf("%s est mort !\n", c.Name)
		c.Pv = c.Pvmax / 2
		fmt.Println("Sergent : « Debout, recrue. On ne meurt pas à l'exercice. »")
		fmt.Printf("%s ressuscite avec %d/%d PV\n", c.Name, c.Pv, c.Pvmax)
		return true
	}
	return false
}

func poisonPot(c *Character) {
	fmt.Printf("%s boit la Potion de poison...\n", c.Name)

	const (
		degatsParSeconde = 10
		dureeSecondes    = 3
	)

	for tick := 1; tick <= dureeSecondes; tick++ {
		time.Sleep(1 * time.Second)

		c.Pv -= degatsParSeconde
		if c.Pv < 0 {
			c.Pv = 0
		}

		fmt.Printf("[Poison %d/3s] %s : %d/%d PV\n", tick, c.Name, c.Pv, c.Pvmax)

		if c.Pv == 0 {
			isDead(c)
			return
		}
	}
	isDead(c)
}
