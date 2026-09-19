package main

// le perso est global : tous les fichiers du package peuvent le lire
var joueur Character

func main() {
	joueur.initCharactere("Recrue", "Humain")
	menuPrincipal()
}
