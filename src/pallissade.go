package src

import "fmt"

func lireGraffitisPalissade() {
	effacerEcran()
	fmt.Println()
	fmt.Println("     " + cGras + cViolet6 + "L A   P A L I S S A D E" + cReset)
	fmt.Println()
	fmt.Println(cGris + "     Tu t'approches de la palissade en bois du camp." + cReset)
	fmt.Println(cGris + "     Des noms ont ete graves la par des recrues des nuits precedentes..." + cReset)
	fmt.Println()
	fmt.Println("     Sur une planche, tu dechiffres une serie de titres :")
	fmt.Println("     " + cJaune + "\"Money, money, money\", \"Two for the Price of One\"," + cReset)
	fmt.Println("     " + cJaune + "\"Gimme! Gimme! Gimme!\", \"I saw it in the mirror\"," + cReset)
	fmt.Println("     " + cJaune + "\"Mamma Mia\", \"On and on and on\"..." + cReset)
	fmt.Println("     Ce sont des chansons du groupe : " + cCyan + "ABBA" + cReset)
	fmt.Println()
	fmt.Println("     Un peu plus loin, une autre main a grave :")
	fmt.Println("     " + cJaune + "\"Fighter Squad\" (1961), \"Duel\" (1971)," + cReset)
	fmt.Println("     " + cJaune + "\"A.I. Intelligence artificielle\" (2001), \"Ready Player One\" (2018)..." + cReset)
	fmt.Println("     Ce sont des films de : " + cCyan + "Steven Spielberg" + cReset)
	fmt.Println()
	fmt.Println("     " + cRouge + "Les deux artistes caches dans le camp sont donc :" + cReset)
	fmt.Println("     " + cRouge + "ABBA et Steven Spielberg." + cReset)
	fmt.Println()
	fmt.Print("     Appuie sur Entree pour revenir au camp...")
	lireChoix()
}