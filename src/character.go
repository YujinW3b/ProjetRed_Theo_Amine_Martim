package src

import "fmt"

// Const potions
const potionLife = "potions 2 vie"

type Character struct {
	Name                string
	Class               string
	Level               int
	Pvmax               int
	Pv                  int
	Inventaire          []string
	PotionGratuiteRecue bool
	Equipment           Equipment
}

func (c *Character) initCharacter(name string, class string, level int, pvmax int, pv int, inventaire []string) {
	c.Name = name
	c.Class = class
	c.Level = level
	c.Pvmax = pvmax
	c.Pv = pv
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
	fmt.Printf("\tCasque : %v\n", c.Equipment.Tete)
	fmt.Printf("\tBuste : %v\n", c.Equipment.Torse)
	fmt.Printf("\tBottes : %v\n", c.Equipment.Pieds)
}

type Monster struct {
	Name   string
	Pvmax  int
	Pv     int
	Attack int
}
