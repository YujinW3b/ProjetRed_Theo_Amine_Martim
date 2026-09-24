package src

import "fmt"

func gagnerXP(c *Character, xpGagnee int) {
	fmt.Printf("Sergent : « %s remporte le combat et gagne %d XP ! »\n", c.Name, xpGagnee)
	c.Xp += xpGagnee

	for c.Xp >= c.XpMax {
		c.Xp -= c.XpMax
		c.Level++
		c.PvmaxBase += 10
		nouveauPalier := int(float64(c.XpMax) * 1.5)
		c.XpMax = nouveauPalier

		recalculerPvmax(c)
		c.Pv = c.Pvmax

		fmt.Printf("Sergent : « NIVEAU SUPÉRIEUR ! Tu passes niveau %d, recrue ! »\n", c.Level)
		fmt.Printf("Sergent : « Bonus : +10 PV max (total : %d), soin complet ! »\n", c.Pvmax)
	}

	fmt.Printf("XP : %d/%d\n", c.Xp, c.XpMax)
}
