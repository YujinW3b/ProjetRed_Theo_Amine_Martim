package src

import "fmt"

const (
	itemPotionDeVie       = "Potion de vie"
	itemPotionDePoison    = "Potion de poison"
	itemLivreBouleDeFeu   = "Livre de Sort : Boule de Feu"
	itemFourrureDeLoup    = "Fourrure de Loup"
	itemPeauDeTroll       = "Peau de Troll"
	itemCuirDeSanglier    = "Cuir de Sanglier"
	itemPlumeDeCorbeau    = "Plume de Corbeau"
	itemAugmentationInv   = "Augmentation d'inventaire"
	itemChapeauAventurier = "Chapeau de l'aventurier"
	itemTuniqueAventurier = "Tunique de l'aventurier"
	itemBottesAventurier  = "Bottes de l'aventurier"

	spellCoupDePoing = "Coup de poing"
	spellBouleDeFeu  = "Boule de Feu"
	spellSouffleDuSergent = "Souffle du Sergent"
)

func findItemIndex(inventaire []string, nom string) int {
	for i, item := range inventaire {
		if item == nom {
			return i
		}
	}
	return -1
}

func countItem(inventaire []string, nom string) int {
	count := 0
	for _, item := range inventaire {
		if item == nom {
			count++
		}
	}
	return count
}

func checkInventoryLimit(c *Character) bool {
	return len(c.Inventaire) < c.InventaireMax
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

func takePot(c *Character) {
	i := findItemIndex(c.Inventaire, itemPotionDeVie)

	switch {
	case i == -1:
		fmt.Println("Sergent : « Ta besace est VIDE, pas de potion, pas de soin ! »")
	case c.Pv >= c.Pvmax:
		fmt.Println("Sergent : « PV au max, on ne gâche pas une potion ! »")
	default:
		removeInventory(c, itemPotionDeVie)
		c.Pv += 50
		if c.Pv > c.Pvmax {
			c.Pv = c.Pvmax
		}
		fmt.Printf("Sergent : « Potion avalée ! PV : %d/%d »\n", c.Pv, c.Pvmax)
	}
}

func upgradeInventorySlot(c *Character) bool {
	if c.InventoryUpgradesUsed >= 3 {
		fmt.Println("Sergent : « Le Régisseur n'a plus de besace renforcée à te vendre, recrue. »")
		return false
	}
	c.InventaireMax += 10
	c.InventoryUpgradesUsed++
	fmt.Printf("Sergent : « Nouvelle besace ! Capacité : %d emplacements. »\n", c.InventaireMax)
	return true
}

func accessInventory(c *Character) {
	for {
		fmt.Println("=== Affichage de l'inventaire ===")
		if len(c.Inventaire) == 0 {
			fmt.Println("Inventaire vide ...")
		} else {
			for itemIndex, itemName := range c.Inventaire {
				fmt.Printf("\t %d - %s\n", itemIndex+1, itemName)
			}
		}
		fmt.Printf("%d/%d\n", len(c.Inventaire), c.InventaireMax)

		fmt.Println("0 - Retour au feu de camp.")
		fmt.Println("Votre choix ?")

		var choix int
		fmt.Scan(&choix)

		if choix == 0 {
			return
		}
		if choix < 1 || choix > len(c.Inventaire) {
			fmt.Println("Sergent : « Erreur dans le choix... »")
			continue
		}

		itemChoisi := c.Inventaire[choix-1]
		switch itemChoisi {
		case itemPotionDeVie:
			takePot(c)
		case itemPotionDePoison:
			removeInventory(c, itemPotionDePoison)
			poisonPot(c)
		case itemLivreBouleDeFeu:
			if spellBook(c) {
				removeInventory(c, itemLivreBouleDeFeu)
			}
		default:
			fmt.Println("Sergent : « Cet objet ne peut pas être utilisé ici. »")
		}
	}
}
