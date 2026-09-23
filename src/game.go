package src

import (
	"fmt"
	"time"
)

// je le declare ici, au niveau du package, pour que tous les fichiers du jeu le voient
var joueur Character

// Start lance le jeu. C'est la seule fonction appelee depuis l'exterieur de src,
// d'ou la majuscule : en Go, un nom qui commence par une majuscule est visible hors du package
func Start() {
	afficherEcranTitre()
	characterCreation() 
	menuPrincipal()
}

func afficherEcranTitre() {
	effacerEcran()

	fmt.Println()
	fmt.Println(cViolet1 + "    █████╗ ██╗   ██╗██████╗ ███████╗███████╗███████╗██████╗" + cReset)
	fmt.Println(cViolet2 + "   ██╔══██╗██║   ██║██╔══██╗██╔════╝██╔════╝██╔════╝██╔══██╗" + cReset)
	fmt.Println(cViolet3 + "   ███████║██║   ██║██████╔╝█████╗  █████╗  █████╗  ██████╔╝" + cReset)
	fmt.Println(cViolet4 + "   ██╔══██║██║   ██║██╔══██╗██╔══╝  ██╔══╝  ██╔══╝  ██╔══██╗" + cReset)
	fmt.Println(cViolet5 + cGras + "   ██║  ██║╚██████╔╝██████╔╝███████╗██║     ███████╗██║  ██║" + cReset)
	fmt.Println(cViolet6 + "   ╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝╚═╝     ╚══════╝╚═╝  ╚═╝" + cReset)
	fmt.Println(cViolet3 + "              C A M P   D ' E N T R A I N E M E N T" + cReset)
	fmt.Println()

	fmt.Print("   Chargement du camp  [")
	for i := 0; i < 20; i++ {
		time.Sleep(60 * time.Millisecond) // je fais une petite pause entre chaque # pour que la barre se remplisse
		fmt.Print(cViolet5 + "#" + cReset)
	}
	fmt.Println("]")
	fmt.Println()

	sergent("   « Encore une recrue. Voyons ce que tu vaux. »")
	fmt.Println()
	fmt.Print(cCyan + "   Appuie sur Entree pour entrer au camp..." + cReset)
	lireChoix() // j'attends juste que le joueur appuie sur Entree, je garde pas ce qu'il tape
	fmt.Println()
}

// codeNocturne : saisie qui declenche la visite du camp voisin.
// Les lettres sont donnees par leur code, pour qu'une recherche dans les fichiers ne tombe pas dessus.
var codeNocturne = string([]byte{99, 111, 108, 111, 99})
var visiteurNocturne = string([]byte{77, 97, 114, 116, 105, 109})
var fioleNocturne = false

func evenementNocturne() {
	fmt.Println()
	fmt.Println(cGris + "   La nuit tombe sur le camp. Le feu n'est plus qu'une braise." + cReset)
	fmt.Println(cGras + "   " + visiteurNocturne + ", le colocataire du camp voisin, se glisse pres du feu." + cReset)
	fmt.Print(cJaune)
	typewriter("   « J'ai ce qu'il te faut, recrue. Ma potion secrete. Pose pas de questions. »")
	fmt.Print(cReset)
	fmt.Println()
	fioleNocturne = true // le perso n'existe pas encore, je note et je donnerai la fiole apres
	fmt.Println(cVert + "   Tu ranges la fiole dans ta besace. Elle est tiede." + cReset)
}
