package src

import (
	"bufio"
	"fmt"
	"os"
)

// codes couleur du terminal : je les colle devant le texte, et cReset revient a la normale
var (
	cReset = "\033[0m"
	cGras  = "\033[1m"
	cGris  = "\033[2m"
	cRouge = "\033[31m"
	cJaune = "\033[33m"
	cCyan  = "\033[36m"
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
	fmt.Println(cRouge + "              )" + cReset)
	fmt.Println(cRouge + "         (   ) (" + cReset + "        " + cGras + "L E   C A M P   D ' A U B E F E R" + cReset)
	fmt.Println(cJaune + "          ) (  )" + cReset)
	fmt.Println(cJaune + "         _(____)_" + cReset + "       " + cGris + "Le feu crepite. Le Sergent t'attend." + cReset)
	fmt.Println(cGris + "        '--------'" + cReset)
	fmt.Println()
	fmt.Println(cGris + "   ---------------------------------------------" + cReset)
	fmt.Println("     " + cCyan + "[1]" + cReset + "  Ta fiche de recrue")
	fmt.Println("     " + cCyan + "[2]" + cReset + "  Ta besace")
	fmt.Println("     " + cCyan + "[3]" + cReset + "  La tente du Regisseur")
	fmt.Println("     " + cCyan + "[4]" + cReset + "  Le terrain d'exercice")
	fmt.Println("     " + cCyan + "[0]" + cReset + "  Quitter le camp")
	fmt.Println(cGris + "   ---------------------------------------------" + cReset)
	fmt.Println()
	fmt.Print("     Ton choix : ") // print et pas println car je veux que le joueur tape a cote
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
		case "4":
			trainingFight()
		case "0":
			fmt.Println("   Le Sergent hoche la tete. A demain, recrue.")
			return // return et pas break car break fais que sortrir switch pas dla boucle
		case codeNocturne: // saisie non listee dans le menu
			evenementNocturne()
		default: // prend tout reste ( vide.. ) et fais renvoie une rep
			fmt.Println("   Parle plus clairement, recrue.")
		}
	}
}
