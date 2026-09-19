package main

import "fmt"

const potionLife = "potions 2 vie"

type Character struct {
	Name                string
	Class               string
	Pvmax               int
	Pv                  int
	Inventaire          []string
	PotionGratuiteRecue bool
}

func (c *Character) initCharactere(name string, class string) {
	c.Name = name
	c.Class = class
	c.Pv = 50
	c.Pvmax = 100
	c.Inventaire = []string{}
}

func (c *Character) displayInfo() {
	fmt.Println("=== info ===")
	fmt.Printf("\tnom : %s\n", c.Name)
	fmt.Printf("\tClasse : %s\n", c.Class)
	fmt.Printf("\tPvmax : %d\n", c.Pvmax)
	fmt.Printf("\tPv : %d\n", c.Pv)
}

// MARTIM - T01 Character, T02 initCharacter, T03 displayInfo, T08 isDead,
// T13 l or, T16 la structure Equipment.
// initCharacter recoit ses valeurs EN PARAMETRES, elle ne les decide jamais.
