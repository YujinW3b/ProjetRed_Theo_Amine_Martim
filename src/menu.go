package main

import (
	"bufio"
	"fmt"
	"os"
)

// je le cree une seule fois ici, sinon je perds ce qui traine dans le tampon
var lecteur = bufio.NewScanner(os.Stdin)

func lireChoix() string {
	lecteur.Scan()        // ici le programme se fige et attend que le joueur tape
	return lecteur.Text() // Text me rend la ligne sans le \n ni le \r de Windows
}

func afficherMenu() {
	fmt.Println()
	fmt.Println("   L E   C A M P   D ' A U B E F E R")
	fmt.Println()
	fmt.Println("   1. Ta fiche de recrue")
	fmt.Println("   2. Ta besace")
	fmt.Println("   0. Quitter le camp")
	fmt.Println()
	fmt.Print("   Ton choix : ") // Print et pas Println, le curseur reste a cote
}

func menuPrincipal() {
	for { // boucle infinie, seul le return du choix 0 en sort
		afficherMenu()
		choix := lireChoix()

		switch choix {
		case "1":
			fmt.Println("   [temporaire] ta fiche de recrue")
		case "2":
			fmt.Println("   [temporaire] ta besace")
		case "0":
			fmt.Println("   Le Sergent hoche la tete. A demain, recrue.")
			return // pas break, il sortirait du switch et la boucle repartirait
		default: // attrape tout : saisie vide, lettres, chiffres hors menu
			fmt.Println("   Parle plus clairement, recrue.")
		}
	}
}
