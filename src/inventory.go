package main

import "fmt"

func findItemIndex(inventaire []string, nom string) int {
	for i, item := range inventaire {
		if item == nom {
			return i
		}
	}
	return -1
}

func checkInventoryLimit(c *Character) bool {
	return len(c.Inventaire) < c.
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
