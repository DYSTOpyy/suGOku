package main

import (
	"fmt"
	"time"    // pour voir le temps de résolution
	"git.saussesylva.in/DYSTO_pyy/Sudoku/algorythm"
	"git.saussesylva.in/DYSTO_pyy/Sudoku/utils"
)


// afficher_temps permet d'afficher une durée efficacement.
// Elle détermine si le meilleur affichage est en secondes, en minutes ou millisecondes.
func afficher_temps (temps time.Duration) {
	fmt.Print("TEMPS ÉCOULÉ : ")
	if temps.Milliseconds() > 60000 {
		fmt.Println(temps.Minutes(), "MINUTE.")
	} else if temps.Milliseconds() > 1000 {
		fmt.Println(temps.Seconds(), "SECONDE.")
	} else {
		fmt.Println(temps.Milliseconds(), "MILLISECONDES.")
	}
}

func main() {
	// grille, possibilite , _ := init_grille(utils.Grille_sudoku_exemple())		// REMPLACER _ PAR masque QUAND ON L'UTILISE

	// fmt.Println("GRILLE DE DEPART : ")
	// utils.Print_grille(&grille, false)
	// fmt.Println()

	// before := time.Now()
	// // fmt.Println("Nombre de solutions : ", resolution(&grille, &possibilite, true))	// CALCUL DU NB DE SOLUTIONS
	// resolution(&grille, &possibilite,  false)		// RESOLUTION DE LA GRILLE
	// after := time.Now()

	// afficher_temps(after.Sub(before))

	// utils.Print_grille(&grille, false)
	grille, mask, err := utils.ImportFile()
	if err != nil {
		fmt.Println(err)
	} else {
		possibilite := utils.Generer_possibilite(grille)
		utils.Print_grille(grille)
		utils.RestartGrille(grille,mask)
		utils.Print_grille(grille)
		fmt.Println(algo.Resolution(grille,&possibilite,true))

	} 
	// before := time.Now()
	// NewGrille := algo.GeneratorFull(algo.Easy)
	// after := time.Now()
	// utils.Print_grille(NewGrille,false)
	// afficher_temps(after.Sub(before))
	// utils.SaveFile(*NewGrille,utils.Generer_masque(NewGrille))

	// possibilite2 := utils.Generer_possibilite(NewGrille)
	// fmt.Println(resolution(NewGrille, &possibilite2, false))
	// utils.Print_grille(NewGrille,false)
	// GrilleGen, possibilite2 , _ := init_grille(NewGrille)
	// fmt.Println("GRILLE DE DEPART : ")
	// utils.Print_grille(&GrilleGen, false)
	// resolution(&GrilleGen, &possibilite2, false)	
	// utils.Print_grille(&GrilleGen, false)

	// fmt.Print(grille_to_string(&grille))

}