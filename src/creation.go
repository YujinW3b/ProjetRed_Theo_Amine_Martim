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
