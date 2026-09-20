package main

// je le declare ici et pas dans main() pour que menu.go puisse le voir aussi
var joueur Character

func main() {

	joueur.initCharacter("Recrue", "Humain", 1, 100, 50, []string{})
	menuPrincipal()
}
