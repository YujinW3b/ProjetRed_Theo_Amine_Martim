package main

import "fmt"

const potionLife = "potions 2 vie"

type Charactere struct {
	Name       string
	Class      string
	Pvmax      int
	Pv         int
	Inventaire map[string]int
}

func (c *Charactere) initCharactere(name string, class string) {
	c.Name = name
	c.Class = class
	c.Pv = 50
	c.Pvmax = 100
	c.Inventaire = map[string]int{potionLife: 5}
}

func (c *Charactere) displayInfo() {
	fmt.Println("=== info ===")
	fmt.Printf("\tnom : %s\n", c.Name)
	fmt.Printf("\tClasse : %s\n", c.Class)
	fmt.Printf("\tPvmax : %d\n", c.Pvmax)
	fmt.Printf("\tPv : %d\n", c.Pv)
}

// MARTIM - T01 Character, T02 initCharacter, T03 displayInfo, T08 isDead,
// T13 l or, T16 la structure Equipment.
// initCharacter recoit ses valeurs EN PARAMETRES, elle ne les decide jamais.