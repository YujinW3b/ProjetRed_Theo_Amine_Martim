package src

import "fmt"

func accessMarchand(c *Character) {
	for {
		fmt.Println("\n--- LE MARCHAND ---")
		fmt.Println("1. Potion de vie (Gratuit)")
		fmt.Println("2. Retour")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if !c.PotionGratuiteRecue {
				addInventory(c, "Potion de vie")
				c.PotionGratuiteRecue = true
				fmt.Println("Vous avez reçu : Potion de vie")
			} else {
				fmt.Println("Vous avez déjà reçu votre potion gratuite !")
			}
		case 2:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
