package main

import (
	"fmt"
	"math"
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

// maj_compteurs met à jour les compteurs du nombre d'indices par ligne/colonne/bloc
// d'une grille qui sont situés respectivement sur la colonne TAILLE,
// la ligne TAILLE et la ligne TAILLE+1, ainsi que le nombre total d'indice
// de la grille situé en TAILLE:TAILLE.
//
// La grille en argument doit être sous forme de pointeur.
func maj_compteurs(grille *[TAILLE + 2][TAILLE + 1]int, size int) {

	for line := 0; line < size; line++ {
		for column := 0; column < size; column++ {
			if grille[line][column] != 0 {

				grille[line][size] += 1                                                                                                              // maj des lignes
				grille[size][column] += 1                                                                                                            // maj des colonnes
				grille[size+1][(line/int(math.Sqrt(float64(size))))*int(math.Sqrt(float64(size)))+(column/int(math.Sqrt(float64(size))))] += 1 // maj des bloc
				grille[size][size] += 1                                                                                                            // maj du nombre total d'indice
			}
		}
	}
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
			if utils.IsOkayCase(line, column, *grille, TAILLE, value) {
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
func init_grille (grille *[TAILLE + 2][TAILLE + 1]int, size int) ([TAILLE + 2][TAILLE + 1]int,[TAILLE][TAILLE][]int, [TAILLE][TAILLE]bool) {
	
	// générer la grille
	// mettre à jour ses compteurs
	maj_compteurs(grille, size)
	// générer la grille de possibilité
	// renvoyer les deux
	return *grille, utils.Generer_possibilite(grille,size), utils.Generer_masque(grille,size)
}

// resolution permet d'appeler les fonctions de résolution de grille sudoku.
// A partir d'une grille et de ses possibilites, si print_nb_solution est vrai : renvoi le nombre de solutions possibles.
// La grille ne sera pas modifiée. Si print_nb_solution est faux, la grille sera résolue selon la première solution trouvée.
// 
// La grille en argument doit être sous forme de pointeur.
func resolution (grille *[TAILLE + 2][TAILLE + 1]int, possibilite *[TAILLE][TAILLE][]int, size int,print_nb_solution bool) int {	
	liste := algo.GenSlice(size ,*grille)
	if print_nb_solution {
		return nombre_solutions(grille, 0, 0, possibilite)
	} else {
		fmt.Println(algo.Algo_backtracking(grille, possibilite, liste, size))
		return 0
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
	size := 9
	// grille, possibilite , _ := init_grille(utils.Grille_sudoku_exemple(size),size)		// REMPLACER _ PAR masque QUAND ON L'UTILISE

	// fmt.Println("GRILLE DE DEPART : ")
	// utils.Print_grille(&grille,size, false)
	// fmt.Println()

	// before := time.Now()
	// // fmt.Println("Nombre de solutions : ", resolution(&grille, &possibilite, true))	// CALCUL DU NB DE SOLUTIONS
	// resolution(&grille, &possibilite, size, false)		// RESOLUTION DE LA GRILLE
	// after := time.Now()

	// afficher_temps(after.Sub(before))

	// utils.Print_grille(&grille,size, false)

	before := time.Now()
	NewGrille := algo.GeneratorFull(size,algo.Easy)
	after := time.Now()
	afficher_temps(after.Sub(before))
	utils.Print_grille(NewGrille,size,false)
	// GrilleGen, possibilite2 , _ := init_grille(NewGrille,size)
	// fmt.Println("GRILLE DE DEPART : ")
	// utils.Print_grille(&GrilleGen,size, false)
	// resolution(&GrilleGen, &possibilite2, size, false)	
	// utils.Print_grille(&GrilleGen,size, false)

	// fmt.Print(grille_to_string(&grille))

}

// github pour faire du sdl en go veandco go sdl 2
// demander pour le "go run UTILS.GO MAIN.GO"
// et l'histoire des packages
