package main

import "fmt"

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
