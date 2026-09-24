package src

import "fmt"

const inventaireCapaciteInitiale = 10

type Character struct {
	Name                  string
	Class                 string
	Level                 int
	Initiative            int // M1 : qui commence le tour de combat
	Pvmax                 int
	PvmaxBase             int
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
<<<<<<< HEAD
	Name     string
	Pvmax    int
	Pv       int
	Attack   int
	XPDonnee int
=======
	Name   string
	Pvmax  int
	Pv     int
	Attack int
	Initiative int // M1 : la vitesse du monstre
>>>>>>> 7eac4e221e6aefe12e5ae31de4f1bedba45e1a57
}

func (c *Character) initCharacter(name string, class string, level int, pvmax int, pv int, inventaire []string) {
	c.Name = name
	c.Class = class
	c.Level = level
	c.PvmaxBase = pvmax
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

func (c *Character) isDead() bool {
	if c.Pv > 0 {
		return false
	}
	fmt.Println("Vous êtes mort ..., Revivre ?")
	c.Pv = c.Pvmax / 2
	return true
}
