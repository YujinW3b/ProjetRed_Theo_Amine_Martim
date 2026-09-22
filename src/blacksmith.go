package src

import "fmt"

const prixForgeron = 5

func recetteOK(c *Character, materiaux map[string]int) bool {
	for item, qte := range materiaux {
		if countItem(c.Inventaire, item) < qte {
			return false
		}
	}
	return true
}

func totalMateriaux(materiaux map[string]int) int {
	total := 0
	for _, qte := range materiaux {
		total += qte
	}
	return total
}

func fabriquer(c *Character, nomEquipement string, materiaux map[string]int) {
	if !recetteOK(c, materiaux) {
		fmt.Println("Sergent : « Il te manque des matériaux, recrue. »")
		return
	}
	if c.Or < prixForgeron {
		fmt.Println("Sergent : « T'as pas assez d'or pour ça, recrue. »")
		return
	}
	if len(c.Inventaire)-totalMateriaux(materiaux)+1 > c.InventaireMax {
		fmt.Println("Sergent : « Ta besace déborde, recrue. »")
		return
	}

	for item, qte := range materiaux {
		for i := 0; i < qte; i++ {
			removeInventory(c, item)
		}
	}
	c.Or -= prixForgeron
	addInventory(c, nomEquipement)
	fmt.Printf("Sergent : « %s forgé ! Or restant : %d »\n", nomEquipement, c.Or)
}

func accessForgeron(c *Character) {
	for {
		fmt.Println("=== Le Forgeron ===")
		fmt.Println("1 - Chapeau de l'aventurier (1 Plume de Corbeau + 1 Cuir de Sanglier, 5 or)")
		fmt.Println("2 - Tunique de l'aventurier (2 Fourrures de Loup + 1 Peau de Troll, 5 or)")
		fmt.Println("3 - Bottes de l'aventurier (1 Fourrure de Loup + 1 Cuir de Sanglier, 5 or)")
		fmt.Println("0 - Retour au feu de camp.")
		fmt.Println("Votre choix ?")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 0:
			return
		case 1:
			fabriquer(c, itemChapeauAventurier, map[string]int{itemPlumeDeCorbeau: 1, itemCuirDeSanglier: 1})
		case 2:
			fabriquer(c, itemTuniqueAventurier, map[string]int{itemFourrureDeLoup: 2, itemPeauDeTroll: 1})
		case 3:
			fabriquer(c, itemBottesAventurier, map[string]int{itemFourrureDeLoup: 1, itemCuirDeSanglier: 1})
		default:
			fmt.Println("Sergent : « Erreur dans le choix... »")
		}
	}
}
