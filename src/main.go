package main

// le perso est global : tous les fichiers du package peuvent le lire
var joueur Charactere

func main() {
	joueur.initCharactere("Recrue", "Humain")
	menuPrincipal()
}