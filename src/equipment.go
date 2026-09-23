package src

import "fmt"

type Equipment struct {
	Tete  string
	Torse string
	Pieds string																																																																																																																																																																										
}

func bonusPvEquipement(item string) int {
	switch item {
	case itemChapeauAventurier:
		return 10
	case itemTuniqueAventurier:
		return 25
	case itemBottesAventurier:
		return 15
	}
	return 0
}

func recalculerPvmax(c *Character) {
	c.Pvmax = c.PvmaxBase
	c.Pvmax += bonusPvEquipement(c.Equipment.Tete)
	c.Pvmax += bonusPvEquipement(c.Equipment.Torse)
	c.Pvmax += bonusPvEquipement(c.Equipment.Pieds)

	if c.Pv > c.Pvmax {
		c.Pv = c.Pvmax
	}
}

func equiperItem(c *Character, item string) {
	var emplacementActuel *string

	switch item {
	case itemChapeauAventurier:
		emplacementActuel = &c.Equipment.Tete
	case itemTuniqueAventurier:
		emplacementActuel = &c.Equipment.Torse
	case itemBottesAventurier:
		emplacementActuel = &c.Equipment.Pieds
	default:
		fmt.Println("Sergent : « Cet objet ne s'équipe pas, recrue. »")
		return
	}

	if !removeInventory(c, item) {
		fmt.Println("Sergent : « Tu n'as pas cet équipement dans ta besace, recrue. »")
		return
	}

	ancienItem := *emplacementActuel

	if ancienItem != "" {
		if !checkInventoryLimit(c) {
			addInventory(c, item)
			fmt.Println("Sergent : « Ta besace déborde, impossible d'échanger cet équipement, recrue. »")
			return
		}
		addInventory(c, ancienItem)
		fmt.Printf("Sergent : « %s remplacé, remis dans ta besace. »\n", ancienItem)
	}

	*emplacementActuel = item
	recalculerPvmax(c)
	fmt.Printf("Sergent : « %s équipé ! PV max : %d »\n", item, c.Pvmax)
}
