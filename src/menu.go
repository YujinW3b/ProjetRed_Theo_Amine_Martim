package src

import (
	"bufio"
	"fmt"
	"os"
)

var lecteur = bufio.NewScanner(os.Stdin) // un seul scanner pour tout le jeu

func retireEspaces(s string) string {
	debut := 0    // je set debut = 0
	fin := len(s) // je set fin sur dernier caractere

	for debut < fin && s[debut] == ' ' { // javance espace de gauche jusqua caractere petit a petit et fin évite de depasser la limite
		debut++ // jeface rien je deplace juste la borne
	}

	for fin > debut && s[fin-1] == ' ' { // pareil mais en partant de la fin de droite jusqua caracterre
		fin-- // fin-1 car la fin est 1 caractere apres la vrai fin
	}

	return s[debut:fin] // a détaillé
}

func lireChoix() string {
	lecteur.Scan()                       // programme figer attend joueur donne reponse
	return retireEspaces(lecteur.Text()) // text rend ligne sans \n ni le \r
} // je nettoie maintenant comme ca switch gere rien

func afficherMenu() {
	fmt.Println()
	fmt.Println("   L E   C A M P   D ' A U B E F E R")
	fmt.Println()
	fmt.Println("   1. Ta fiche de recrue")
	fmt.Println("   2. Ta besace")
	fmt.Println("   3. La tente du Regisseur")
	fmt.Println("   0. Quitter le camp")
	fmt.Println()
	fmt.Print("   Ton choix : ") // print et pas println car je veux que le joueur taper a coté et pas en dessous
}

func menuPrincipal() {
	for { // boucle infinie il y a que return qui fais sortir de la boucle
		afficherMenu()
		choix := lireChoix()

		switch choix {
		case "1":
			joueur.displayInfo()
		case "2":
			fmt.Println("   [temporaire] ta besace")
		case "3":
			accessMarchand(&joueur)
		case "0":
			fmt.Println("   Le Sergent hoche la tete. A demain, recrue.")
			return // return et pas break car break fais que sortrir switch pas dla boucle
		default: // prend tout reste ( vide.. ) et fais renvoie une rep
			fmt.Println("   Parle plus clairement, recrue.")
		}
	}
}
