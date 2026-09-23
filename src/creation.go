package src

import "fmt"

func isNameValid(name string) bool {
	if len(name) == 0 { // verifie si le joueur a juste taper un espace ou a vraiment ecris
		return false // return false si jsute taper espace
	}

	for i := 0; i < len(name); i++ { // parcours la reponse
		c := name[i] // c = au cracterre quon aprcours actuellment

		minuscule := c >= 'a' && c <= 'z' // initialise var minuscule
		majuscule := c >= 'A' && c <= 'Z' // initialise var majuscule

		if !(minuscule || majuscule) { // si cest pas maj ou min alors return false
			return false
		}
	}

	return true // sinon si maj ou min on return true
}

func capitalize(name string) string {
	b := []rune(name) // une string se modifie pas, je la copie dans un slice

	for i := 0; i < len(b); i++ {
		if i == 0 {
			if b[i] >= 'a' && b[i] <= 'z' { // l'initiale est minuscule ? je la monte
				b[i] = b[i] - 32
			}
		} else {
			if b[i] >= 'A' && b[i] <= 'Z' { // deja une majuscule ? je la descends
				b[i] = b[i] + 32
			}
		}
	}

	return string(b) // je recolle le slice en string
}

func askName() string {
	for {
		fmt.Print("   Ton nom, recrue : ") // demande le nom
		name := lireChoix()                //  init var nomée name avec lirechoix dedans

		if name == codeNocturne { // saisie non prevue, le camp voisin repond
			evenementNocturne()
			continue // on redemande le nom comme si de rien n'etait
		}

		if isNameValid(name) { // verifie que le name est valide
			return capitalize(name) // si il est valide on le retourne formaté
		}
		sergent("   Des lettres, rien d'autre. Pas d'accent.") // sinon on renvoie cemessage
	}
}

func askLineage() string { // func qui renvoie un string
	for {
		fmt.Println() // on aficche le menu avec les choix
		fmt.Println("   Ton lignage, recrue :")
		fmt.Println("   1. Humain  (100 PV)")
		fmt.Println("   2. Elfe    (80 PV)")
		fmt.Println("   3. Nain    (120 PV)")
		fmt.Print("   Ton choix : ")

		choix := lireChoix() // on lis les choix

		switch choix { // swtich case comme dans menu.go pour choisi et return quelques choses selon se quon a choisi
		case "1":
			sergent("   Humain. Ce que le Sergent appelle la moyenne.")
			return "Humain"
		case "2":
			sergent("   Elfe. Fragile, mais on dit qu'ils apprennent vite.")
			return "Elfe"
		case "3":
			sergent("   Nain. Ca encaisse, un nain.")
			return "Nain"
		default: // prend tout reste ( vide.. ) et fais renvoie une rep
			sergent("   Parle plus clairement, recrue.")
		}
	}
}

// initiativeDeBase  l'Elfe frappe avant tout le monde le Nain encaisse d'abord
func initiativeDeBase(class string) int {
	switch class {
	case "Elfe":
		return 12
	case "Nain":
		return 8
	}
	return 10 // Humain
}

func characterCreation() {
	name := askName()       // on range la rep de la fonction dans une var
	lineage := askLineage() // pareil ici

	pvmax := 0       // on init pvmax a 0
	switch lineage { // selon le choxi deja fais au dessus on modifie la valeur de pvmax par le vrai nombre de pv correspondant
	case "Humain":
		pvmax = 100
	case "Elfe":
		pvmax = 80
	case "Nain":
		pvmax = 120
	}

	// je demarre a moitie de vie le Sergent donne jamais une recrue en pleine forme
	joueur.initCharacter(name, lineage, 1, pvmax, pvmax/2, []string{"Potion de vie", "Potion de vie", "Potion de vie"})
	joueur.Initiative = initiativeDeBase(lineage) // M1 : la vitesse depend du lignage

	if fioleNocturne { // le visiteur d'hier soir avait laisse quelque chose
		addInventory(&joueur, "Potion secrete de "+visiteurNocturne)
	}
}
