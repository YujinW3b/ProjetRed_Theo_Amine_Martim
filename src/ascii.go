package src

// MARTIM - M5 : habillage ASCII, couleurs, UX.
// Les dessins sont des planches toutes faites (autorise par les profs).
// Les fonctions du bas (barre de vie, cadre, typewriter) sont a lire et a savoir expliquer.

import (
	"fmt"
	"time"
)

// le vert manquait dans la palette de menu.go, je l'ajoute ici
var cVert = "\033[32m"

// ---------------------------------------------------------------------------
// LES PLANCHES ASCII
// J'utilise des accents graves et pas des guillemets : comme ca les antislashs
// des dessins passent tels quels, sans avoir a les doubler.
// ---------------------------------------------------------------------------

var artGobelin = `
        .-"""-.
       / .===. \
       \/ 6 6 \/
       ( \___/ )
    ___ooo___ooo___`

var artGobelinKO = `
        .-"""-.
       / .===. \
       \/ x x \/
       ( \___/ )
    ___ooo___ooo___`

var artForge = `
          ___
      ___/   \___
     /           \
     \___________/
          | |
         _|_|_`

var artTente = `
           /\
          /  \
         /    \
        /______\
        |  []  |
        |______|`

var artPotion = `
         ___
        |   |
        |___|
       /     \
      |  ~~~  |
      |_______|`

var artVictoire = `
__   _____ ___ _____ ___ ___ ___ ___ 
\ \ / /_ _/ __|_   _/ _ \_ _| _ \ __|
 \ V / | | (__  | || (_) | ||   / _| 
  \_/ |___\___| |_| \___/___|_|_\___|`

var artKO = `
 _  __  ___   
| |/ / / _ \  
| ' < | (_) | 
|_|\_(_)___(_)`

// afficherArt affiche une planche dans la couleur demandee
func afficherArt(art string, couleur string) {
	fmt.Println(couleur + art + cReset)
}

// ---------------------------------------------------------------------------
// LES OUTILS D'AFFICHAGE
// ---------------------------------------------------------------------------

// effacerEcran vide le terminal et remonte le curseur en haut
func effacerEcran() {
	fmt.Print("\033[2J\033[H")
}

// barreDeVie fabrique une barre du style [##########------]
// Le principe : 16 cases au total. Je calcule combien doivent etre pleines,
// puis je complete avec des tirets.
func barreDeVie(pv int, pvmax int) string {
	if pvmax <= 0 { // securite : jamais de division par zero
		return "[----------------]"
	}

	pleines := pv * 16 / pvmax // division entiere : a 1 PV ca donne 0 case, c'est voulu
	if pleines < 0 {
		pleines = 0
	}
	if pleines > 16 {
		pleines = 16
	}

	barre := "["
	for i := 0; i < pleines; i++ {
		barre = barre + "#"
	}
	for i := pleines; i < 16; i++ {
		barre = barre + "-"
	}
	barre = barre + "]"

	// la couleur depend de ce qu'il reste : vert au-dessus de la moitie,
	// jaune entre le quart et la moitie, rouge en dessous
	if pv*2 > pvmax {
		return cVert + barre + cReset
	}
	if pv*4 > pvmax {
		return cJaune + barre + cReset
	}
	return cRouge + barre + cReset
}

// largeurCadre est la largeur interieure de tous les ecrans du jeu.
// Une seule valeur pour tout le monde, sinon les bords ne s'alignent pas.
var largeurCadre = 46

// cadreHaut ouvre un ecran avec son titre
func cadreHaut(titre string) {
	fmt.Println()
	fmt.Print("   +")
	for i := 0; i < largeurCadre; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
	cadreLigne(cGras + titre + cReset)
	fmt.Print("   +")
	for i := 0; i < largeurCadre; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

// cadreBas ferme l'ecran
func cadreBas() {
	fmt.Print("   +")
	for i := 0; i < largeurCadre; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

// cadreLigne ecrit une ligne entre les deux bords.
// Le piege : je compte la longueur du texte VISIBLE, sans les codes couleur,
// sinon le bord droit part n'importe ou.
func cadreLigne(texte string) {
	fmt.Print("   |  " + texte)
	for i := longueurVisible(texte); i < largeurCadre-3; i++ {
		fmt.Print(" ")
	}
	fmt.Println("|")
}

// longueurVisible compte les caracteres affiches, en sautant les codes couleur.
// Un code commence par l'echappement \033 et se termine par la lettre m.
func longueurVisible(texte string) int {
	compte := 0
	dansUnCode := false

	for _, c := range texte {
		if c == '\033' {
			dansUnCode = true
		} else if dansUnCode {
			if c == 'm' {
				dansUnCode = false
			}
		} else {
			compte++
		}
	}

	return compte
}

// typewriter affiche un texte lettre par lettre, comme une machine a ecrire.
// A reserver aux repliques du Sergent : sur un menu ce serait insupportable.
func typewriter(texte string) {
	for _, c := range texte {
		fmt.Print(string(c))
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println()
}
