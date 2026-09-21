package main

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
