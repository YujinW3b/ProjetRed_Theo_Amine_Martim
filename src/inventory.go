package src

import "fmt"

<<<<<<< HEAD
func findItemIndex(inventaire []string, nom string) int {
	for i, item := range inventaire {
		if item == nom {
			return i
		}
	}
	return -1
}

func checkInventoryLimit(c *Character) bool {
	return len(c.Inventaire) < c.Inventaire
}

=======
func checkInventoryLimit(c *Character) bool {
	return len(c.Inventaire) < 10 // T18 : a remplacer par une capacite stockee dans Character
}

>>>>>>> e4b247ede91c150e006bf6d8de2eb9938c0cd736
func accessInventory(c *Character, item string) {
	fmt.Println("=== Affichage de l'inventaire ===")
	for itemIndex, itemName := range c.Inventaire {
		fmt.Printf("\t %d - %s\n", (itemIndex + 1), itemName)
	}
	if len(c.Inventaire) == 0 {
		fmt.Println("Inventaire vide ...")
	}
	fmt.Printf("%d/%d\n", len(c.Inventaire), 10)
	for {
		fmt.Println("0 - retour au feu de champs.")
		fmt.Println("Votre choix ?")
		var chose int
		fmt.Scan(&chose)
		switch chose {
		case 0:
			return
		default:
			fmt.Println("Erreur dans le choix...")
		}
	}
<<<<<<< HEAD
=======

>>>>>>> e4b247ede91c150e006bf6d8de2eb9938c0cd736
}

func addInventory(c *Character, item string) bool {
	if !checkInventoryLimit(c) {
		fmt.Println("Sergent : « Ta besace déborde, recrue. »")
		return false
	}
	c.Inventaire = append(c.Inventaire, item)
	return true
}

func removeInventory(c *Character, item string) bool {
	i := findItemIndex(c.Inventaire, item)
	if i == -1 {
		fmt.Println("Item introuvable dans l'inventaire :", item)
		return false
	}
	c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
	return true
}
