package main

import (
	"bufio"
	"fmt"
	"os"
)

var lecteur = bufio.NewScanner(os.Stdin)

func retireEspaces(s string) string {
	debut := 0    // debut pointe sur le premier caractere
	fin := len(s) // fin vaut len(s) cest une position apres le dernier

	for debut < fin && s[debut] == ' ' { // regarde chaque position et verif si cest < a la fin + bien un espace
		debut++ // sa passe caractere a coté +1 a chaque fois et uen fois sa tombe sur 1, c'est 1 on stop
	}

	for fin > debut && s[fin-1] == ' ' { // pareil en partant de fin-1, on baisse a chaque fois cest un espace vide et quand on tombe sur le nombre sa return par exemple return s[0:1]
		fin-- // on commence par trouve le bon "debut" puis la bonne "fin" et on return
	}

	return s[debut:fin]
}

func lireChoix() string {
	lecteur.Scan()                       // ici le programme se fige et attend que le joueur tape
	return retireEspaces(lecteur.Text()) // Text me rend la ligne sans le \n ni le \r de Windows
} // je prends ce que lecteur.Text() me donne, je le fais passer par retireEspaces, et c'est ce qui en sort que je renvoie.

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
			joueur.displayInfo()
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
