package src

import "fmt"

func prixMarchand(c *Character, item string) int {
	switch item {
	case itemPotionDeVie:
		if !c.PotionGratuiteRecue {
			return 0
		}
		return 3
	case itemPotionDePoison:
		return 6
	case itemLivreBouleDeFeu:
		return 25
	case itemFourrureDeLoup:
		return 4
	case itemPeauDeTroll:
		return 7
	case itemCuirDeSanglier:
		return 3
	case itemPlumeDeCorbeau:
		return 1
	case itemAugmentationInv:
		return 30
	}
	return -1
}

func acheterItem(c *Character, item string) {
	prix := prixMarchand(c, item)

	if c.Or < prix {
		fmt.Println("Sergent : « T'as pas assez d'or pour ça, recrue. »")
		return
	}

	if item == itemAugmentationInv {
		if c.InventoryUpgradesUsed >= 3 {
			fmt.Println("Sergent : « Le Régisseur n'a plus de besace renforcée à te vendre, recrue. »")
			return
		}
		c.Or -= prix
		upgradeInventorySlot(c)
		fmt.Printf("Vous avez acheté : %s (Or restant : %d)\n", item, c.Or)
		return
	}

	if !checkInventoryLimit(c) {
		fmt.Println("Sergent : « Ta besace déborde, recrue. »")
		return
	}

	c.Or -= prix
	addInventory(c, item)
	if item == itemPotionDeVie && !c.PotionGratuiteRecue {
		c.PotionGratuiteRecue = true
	}
	fmt.Printf("Vous avez acheté : %s (Or restant : %d)\n", item, c.Or)
}

func accessMarchand(c *Character) {
	for {
		fmt.Println("=== Le Régisseur ===")
		fmt.Printf("1 - Potion de vie (%d or)\n", prixMarchand(c, itemPotionDeVie))
		fmt.Printf("2 - Potion de poison (%d or)\n", prixMarchand(c, itemPotionDePoison))
		fmt.Printf("3 - Livre de Sort : Boule de Feu (%d or)\n", prixMarchand(c, itemLivreBouleDeFeu))
		fmt.Printf("4 - Fourrure de Loup (%d or)\n", prixMarchand(c, itemFourrureDeLoup))
		fmt.Printf("5 - Peau de Troll (%d or)\n", prixMarchand(c, itemPeauDeTroll))
		fmt.Printf("6 - Cuir de Sanglier (%d or)\n", prixMarchand(c, itemCuirDeSanglier))
		fmt.Printf("7 - Plume de Corbeau (%d or)\n", prixMarchand(c, itemPlumeDeCorbeau))
		fmt.Printf("8 - Augmentation d'inventaire (%d or) [%d/3 utilisées]\n", prixMarchand(c, itemAugmentationInv), c.InventoryUpgradesUsed)
		fmt.Println("0 - Retour au feu de camp.")
		fmt.Printf("Votre or : %d\n", c.Or)
		fmt.Println("Votre choix ?")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 0:
			return
		case 1:
			acheterItem(c, itemPotionDeVie)
		case 2:
			acheterItem(c, itemPotionDePoison)
		case 3:
			acheterItem(c, itemLivreBouleDeFeu)
		case 4:
			acheterItem(c, itemFourrureDeLoup)
		case 5:
			acheterItem(c, itemPeauDeTroll)
		case 6:
			acheterItem(c, itemCuirDeSanglier)
		case 7:
			acheterItem(c, itemPlumeDeCorbeau)
		case 8:
			acheterItem(c, itemAugmentationInv)
		default:
			fmt.Println("Sergent : « Erreur dans le choix... »")
		}
	}
}
