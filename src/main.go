package main

import "fmt"

// je le declare ici et pas dans main() pour que menu.go puisse le voir aussi
var joueur Character

func main() {
	// TEST TEMPOR
	fmt.Println(isNameValid("Theo"))       // true
	fmt.Println(isNameValid(""))           // false
	fmt.Println(isNameValid("Bo3"))        // false
	fmt.Println(isNameValid("le sergent")) // false, l'espace
	fmt.Println(isNameValid("Théo"))       // false, le "e" accentue tient sur 2 bytes et aucun n'est entre 'a' et 'z'

	joueur.initCharacter("Recrue", "Humain", 1, 100, 50, []string{})
	menuPrincipal()
}
