package src

import "fmt"

const inventaireCapaciteInitiale = 10

type Character struct {
	Name                  string
	Class                 string
	Level                 int
	Pvmax                 int
	Pv                    int
	Mana                  int
	ManaMax               int
	Inventaire            []string
	InventaireMax         int
	Skill                 []string
	Or                    int
	PotionGratuiteRecue   bool
	InventoryUpgradesUsed int
	Equipment             Equipment
}

type Monster struct {
	Name   string
	Pvmax  int
	Pv     int
	Attack int
}

func (c *Character) initCharacter(name string, class string, level int, pvmax int, pv int, inventaire []string) {
	c.Name = name
	c.Class = class
	c.Level = level
	c.Pvmax = pvmax
	c.Pv = pv
	c.ManaMax = 30
	c.Mana = c.ManaMax
	c.Inventaire = inventaire
	c.InventaireMax = inventaireCapaciteInitiale
	c.Skill = []string{spellCoupDePoing, spellSouffleDuSergent}
	c.Or = 100
	c.PotionGratuiteRecue = false
	c.InventoryUpgradesUsed = 0
}

func (c *Character) displayInfo() {
	fmt.Println("=== info ===")
	fmt.Printf("\tnom : %s\n", c.Name)
	fmt.Printf("\tClasse : %s\n", c.Class)
	fmt.Printf("\tNiveau : %d\n", c.Level)
	fmt.Printf("\tPvmax : %d\n", c.Pvmax)
	fmt.Printf("\tPv : %d\n", c.Pv)
	fmt.Printf("\tMana : %d/%d\n", c.Mana, c.ManaMax)
	fmt.Printf("\tOr : %d\n", c.Or)
	fmt.Printf("\tInventaire (%d/%d) : %v\n", len(c.Inventaire), c.InventaireMax, c.Inventaire)
	fmt.Printf("\tSorts connus : %v\n", c.Skill)
	fmt.Printf("\tCasque : %v\n", c.Equipment.Tete)
	fmt.Printf("\tBuste : %v\n", c.Equipment.Torse)
	fmt.Printf("\tBottes : %v\n", c.Equipment.Pieds)
}

func knowsSkill(c *Character, spell string) bool {
	for _, s := range c.Skill {
		if s == spell {
			return true
		}
	}
	return false
}

func spellBook(c *Character) bool {
	if knowsSkill(c, spellBouleDeFeu) {
		fmt.Println("Sergent : « Tu sais déjà faire ça. »")
		return false
	}
	c.Skill = append(c.Skill, spellBouleDeFeu)
	fmt.Println("Sergent : « Nouveau sort appris : Boule de Feu ! »")
	return true
}
<<<<<<< HEAD
func (c *Character) isDead() bool {
if c.Pv> 0 { 
	return false
}		
	fmt.Println("Vous êtes mort ..., Revivre ?" )
c.Pv= c.Pvmax/2 
return true 
}
=======
>>>>>>> 6331bf2aa6e2cbf1ea9f22e2108f9c2f821a125c
