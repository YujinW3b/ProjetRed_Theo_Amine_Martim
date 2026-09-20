package main

import "fmt"

const potionLife = "potions 2 vie"

type Character struct {
	Name       string
	Class      string
	Level      int
	Pvmax      int
	Pv         int
	Inventaire []string
}

func (c *Character) initCharacter(name string, class string, level int, pvmax int, pv int, inventaire []string) {
	c.Name = name
	c.Class = class
	c.Level = level
	c.Pvmax = 100
	c.Pv = 50
	c.Inventaire = inventaire
}

func (c *Character) displayInfo() {
	fmt.Println("=== info ===")
	fmt.Printf("\tnom : %s\n", c.Name)
	fmt.Printf("\tClasse : %s\n", c.Class)
	fmt.Printf("\tNiveau : %d\n", c.Level)
	fmt.Printf("\tPvmax : %d\n", c.Pvmax)
	fmt.Printf("\tPv : %d\n", c.Pv)
	fmt.Printf("\tInventaire : %v\n", c.Inventaire)
}

func main() {
	var character Character

	character.initCharacter(
		"Hero",
		"Elfe",
		1,
		100,
		40,
		[]string{"Potion", "Potion", "Potion"},
	)

	character.displayInfo()
}

// MARTIM - T01 Character, T02 initCharacter, T03 displayInfo, T08 isDead,
// T13 l or, T16 la structure Equipment.
// initCharacter recoit ses valeurs EN PARAMETRES, elle ne les decide jamais.
