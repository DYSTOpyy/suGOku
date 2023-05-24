package main

import (
	"fmt"
	"time"    // pour voir le temps de résolution

	"git.saussesylva.in/DYSTO_pyy/Sudoku/algorythm"
	"git.saussesylva.in/DYSTO_pyy/Sudoku/utils"
)

// Taille de la grille de sudoku
const TAILLE int = utils.TAILLE

// grille_to_string convertit une grille en chaine de caractère, lisible depuis un .txt .
// 
// La grille en argument doit être sous forme de pointeur.
func grille_to_string(grille *[TAILLE + 2][TAILLE + 1]int) string {

	text := ""
	for line := 0; line < TAILLE; line++ {
		for column := 0; column < TAILLE; column++ {
			if grille[line][column] == 0 {
				text = text + "."
			} else if grille[line][column] > 9 {
				text = text + string(grille[line][column]-10+65)
			} else {
				text = text + string(grille[line][column]+48)
			}
		}
	}
	return text
}









// algo_backtracking résout n'importe quelle grille de sudoku avec la méthode du backtracking.
// Selon sa dimension, cela peut prendre du temps. La grille est parcouru en commençant par la case line:column
// qui sera généralement 0:0. Les valeurs testées sont les valeurs répertoriées dans possibilite afin d'accéler le résultat.
// Les possibilitées ne sont pas mises à jour à chaque nouvelle case insérée, sinon le temps augmente exponentiellement.
// Rien ne dis que la solution est unique : c'est la première qu'il trouve.
//
// La grille en argument doit être sous forme de pointeur, afin qu'elle puisse être directement complétée après exécution (!!).
// possibilite est également un un pointeur, puisqu'il n'est jamais modifié dedans.
// func algo_backtracking(grille *[TAILLE + 2][TAILLE + 1]int, line int, column int, possibilite *[TAILLE][TAILLE][]int) bool {

// 	// Condition d'arrêt : lorsque l'on arrive à la dernière case TAILLE-1:TAILLE-1
// 	if line >= TAILLE || column >= TAILLE {
// 		return true

// 		// Si la case n'est pas vide alors on prend la suivante
// 	} else if grille[line][column] != 0 {
// 		return algo_backtracking(grille, line+(column+1)/TAILLE, (column+1)%TAILLE, possibilite)

// 	} else {

// 		for _, value := range possibilite[line][column] {
// 			if utils.IsOkayCase(line, column, *grille, TAILLE, value) {
// 				// Si la valeur est possible on l'attribue à la case
// 				grille[line][column] = value

// 				if algo_backtracking(grille, line+(column+1)/TAILLE, (column+1)%TAILLE, possibilite) {
// 					return true
// 				}

// 			}
// 		}
// 		grille[line][column] = 0
// 		return false
// 	}
// }

// nombre_solutions détermine le nombre de solutions d'une grille de sudoku.
// Prend beaucoup de temps pour de grandes matrices. La grille est parcouru en commençant par la case line:column
// qui sera généralement 0:0. Les valeurs testées sont les valeurs répertoriées dans possibilite afin d'accéler le résultat.
// Les possibilitées ne sont pas mises à jour à chaque nouvelle case insérée, sinon le temps augmente exponentiellement.
//
// La grille en argument doit être sous forme de pointeur.
// possibilite est également un un pointeur, puisqu'il n'est jamais modifié dedans.
func nombre_solutions(grille *[TAILLE + 2][TAILLE + 1]int, line int, column int, possibilite *[TAILLE][TAILLE][]int) int {

	// Condition d'arrêt : lorsque l'on arrive à la dernière case TAILLE-1:TAILLE-1
	if line >= TAILLE || column >= TAILLE {
		return 1

		// Si la case n'est pas vide alors on prend la suivante
	} else if grille[line][column] != 0 {
		return nombre_solutions(grille, line+(column+1)/TAILLE, (column+1)%TAILLE, possibilite)

	} else {

		var n int = 0

		for _, value := range possibilite[line][column] {
			if utils.IsOkayCase(line, column, *grille, value) {
				// Si la valeur est possible on l'attribue à la case
				grille[line][column] = value

				n = n + nombre_solutions(grille, line+(column+1)/TAILLE, (column+1)%TAILLE, possibilite)

			}
		}
		grille[line][column] = 0
		return n
	}
}

// init_grille génère une grille puis détermine ses compteurs et ses possibilité.
// C'est une fonction à retour multiple, qui renvoie la grille et les possibilités.
// 
// La grille en argument doit être sous forme de pointeur.		DOIT DISPARAITRE POUR LAISSER PLACE A UN ARGUMENT "DIFFICULTE"
func init_grille (grille *[TAILLE + 2][TAILLE + 1]int,) ([TAILLE + 2][TAILLE + 1]int,[TAILLE][TAILLE][]int, [TAILLE][TAILLE]bool) {
	// générer la grille
	// mettre à jour ses compteurs
	utils.Maj_compteurs(grille)
	// générer la grille de possibilité
	// renvoyer les deux
	return *grille, utils.Generer_possibilite(grille), utils.Generer_masque(grille)
}

// resolution permet d'appeler les fonctions de résolution de grille sudoku.
// A partir d'une grille et de ses possibilites, si print_nb_solution est vrai : renvoi le nombre de solutions possibles.
// La grille ne sera pas modifiée. Si print_nb_solution est faux, la grille sera résolue selon la première solution trouvée.
// 
// La grille en argument doit être sous forme de pointeur.
func resolution (grille *[TAILLE + 2][TAILLE + 1]int, possibilite *[TAILLE][TAILLE][]int, print_nb_solution bool) int {	
	liste := algo.GenSlice(grille)
	if print_nb_solution {
		return nombre_solutions(grille, 0, 0, possibilite)
	} else {
		if algo.Algo_backtracking(grille, possibilite, liste){
			return 1
		} else {
			return 0
		}
	}
}

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

	before := time.Now()
	NewGrille := algo.GeneratorFull(algo.Easy)
	after := time.Now()
	utils.Print_grille(NewGrille,false)
	afficher_temps(after.Sub(before))
	possibilite2 := utils.Generer_possibilite(NewGrille)
	fmt.Println(resolution(NewGrille, &possibilite2, false))	
	
	// GrilleGen, possibilite2 , _ := init_grille(NewGrille)
	// fmt.Println("GRILLE DE DEPART : ")
	// utils.Print_grille(&GrilleGen, false)
	// resolution(&GrilleGen, &possibilite2, false)	
	// utils.Print_grille(&GrilleGen, false)

	// fmt.Print(grille_to_string(&grille))

}

// github pour faire du sdl en go veandco go sdl 2
// demander pour le "go run UTILS.GO MAIN.GO"
// et l'histoire des packages
