package main

import "fmt"

func accessInventory(c *Character, item string) {
	fmt.Println("=== Affichage de l'inventaire ===")
	for itemIndex, itemName := range c.Inventaire {
		fmt.Printf("\t %d - %s\n", (itemIndex + 1), itemName)
	}
	if len(c.Inventaire) == 0 {
		fmt.Println("Inventaire vide ...")
	}
	fmt.Printf("%d/%d\n", len(c.Inventaire), 10)
	for true {
		fmt.Println("0 - retour au feu de champs.")
		fmt.Println("Votre choix ?")
		var chose int
		fmt.Scan(&chose)
		switch chose {
		case 0:
			break
		default:
			fmt.Println("Erreur dans le choix...")
		}
	}

}

func addInventory(c *Character, item string) bool {
	if len(c.Inventaire) >= 10 {
		fmt.Println("Inventaire plein ! Impossible d'ajouter :", item)
		return false
	}
	c.Inventaire = append(c.Inventaire, item)
	return true
}

func removeInventory(c *Character, item string) bool {
	for i, it := range c.Inventaire {
		if it == item {
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			return true
		}
	}
	fmt.Println("Item introuvable dans l'inventaire :", item)
	return false
}
