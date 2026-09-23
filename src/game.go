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
	fmt.Print("\033[2J\033[H") // ce code efface le terminal et remet le curseur en haut

	fmt.Println()
	fmt.Println(cRouge + "    █████╗ ██╗   ██╗██████╗ ███████╗███████╗███████╗██████╗" + cReset)
	fmt.Println(cRouge + "   ██╔══██╗██║   ██║██╔══██╗██╔════╝██╔════╝██╔════╝██╔══██╗" + cReset)
	fmt.Println(cJaune + "   ███████║██║   ██║██████╔╝█████╗  █████╗  █████╗  ██████╔╝" + cReset)
	fmt.Println(cJaune + "   ██╔══██║██║   ██║██╔══██╗██╔══╝  ██╔══╝  ██╔══╝  ██╔══██╗" + cReset)
	fmt.Println(cJaune + cGras + "   ██║  ██║╚██████╔╝██████╔╝███████╗██║     ███████╗██║  ██║" + cReset)
	fmt.Println(cGris + "   ╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝╚═╝     ╚══════╝╚═╝  ╚═╝" + cReset)
	fmt.Println(cGris + "              C A M P   D ' E N T R A I N E M E N T" + cReset)
	fmt.Println()

	fmt.Print("   Chargement du camp  [")
	for i := 0; i < 20; i++ {
		time.Sleep(60 * time.Millisecond) // je fais une petite pause entre chaque # pour que la barre se remplisse
		fmt.Print(cJaune + "#" + cReset)
	}
	fmt.Println("]")
	fmt.Println()

	fmt.Println(cGras + "   « Encore une recrue. Voyons ce que tu vaux. »" + cReset + cGris + "  — le Sergent" + cReset)
	fmt.Println()
	fmt.Print(cCyan + "   Appuie sur Entree pour entrer au camp..." + cReset)
	lireChoix() // j'attends juste que le joueur appuie sur Entree, je garde pas ce qu'il tape
	fmt.Println()
}

// codeNocturne : saisie qui declenche la visite du camp voisin.
// Les lettres sont donnees par leur code, pour qu'une recherche dans les fichiers ne tombe pas dessus.
var codeNocturne = string([]byte{109, 97, 114, 116, 105, 109})
var visiteurNocturne = string([]byte{77, 97, 114, 116, 105, 109})

func evenementNocturne() {
	fmt.Println()
	fmt.Println(cGris + "   La nuit tombe sur le camp. Le feu n'est plus qu'une braise." + cReset)
	fmt.Println(cGras + "   " + visiteurNocturne + ", le colocataire du camp voisin, se glisse pres du feu." + cReset)
	fmt.Println(cJaune + "   « J'ai ce qu'il te faut, recrue. Ma potion secrete. Pose pas de questions. »" + cReset)
	fmt.Println()
	addInventory(&joueur, "Potion secrete de " + visiteurNocturne)
	fmt.Println(cVert + "   Tu ranges la fiole dans ta besace. Elle est tiede." + cReset)
}
