package src

// je le declare ici, au niveau du package, pour que tous les fichiers du jeu le voient
var joueur Character

// Start lance le jeu. C'est la seule fonction appelee depuis l'exterieur de src,
// d'ou la majuscule : en Go, un nom qui commence par une majuscule est visible hors du package
func Start() {
	characterCreation()
	menuPrincipal()
}
