package main

import (
	"fmt"
	"math"
	"strconv" // int to string
	"time"    // pour voir le temps de résolution

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
func maj_compteurs(grille *[TAILLE + 2][TAILLE + 1]int) {

	for line := 0; line < TAILLE; line++ {
		for column := 0; column < TAILLE; column++ {
			if grille[line][column] != 0 {

				grille[line][TAILLE] += 1                                                                                                              // maj des lignes
				grille[TAILLE][column] += 1                                                                                                            // maj des colonnes
				grille[TAILLE+1][(line/int(math.Sqrt(float64(TAILLE))))*int(math.Sqrt(float64(TAILLE)))+(column/int(math.Sqrt(float64(TAILLE))))] += 1 // maj des bloc
				grille[TAILLE][TAILLE] += 1                                                                                                            // maj du nombre total d'indice
			}
		}
	}
}

// generer_possibilite calcule les valeurs possibles pour chaque case vide
// d'une grille qu'il place dans un slice de dimension TAILLE:TAILLE (donc pas la même dimension que la grille !)
// qui EST RENVOYÉ où chaque case contient un slice représentant les valeurs possibles (vide si la case est occupée).
//
// La grille en argument doit être sous forme de pointeur.
func generer_possibilite(grille *[TAILLE + 2][TAILLE + 1]int) [TAILLE][TAILLE][]int {

	possibilite := [TAILLE][TAILLE][]int{}

	// pour chaque case
	for line := 0; line < TAILLE; line++ {
		for column := 0; column < TAILLE; column++ {

			c := []int{}

			for value := 1; value < TAILLE+1; value++ {

				if grille[line][column] == 0 && utils.IsOkayCase(line, column, *grille, TAILLE, value) {

					// si la valeur est possible, on l'ajoute au slice des possibilitées
					c = append(c, value)

				}
			}
			possibilite[line][column] = c
		}
	}
	return possibilite
}

// generer_masque crée un slice de boolean qui indique si la case était là au départ ou non.
// si une case vaut vrai, cela signifie que la valeur correspondante dans la grille était déjà là.
// si c'est faux, cela signifie que la case était vide au départ et que son contenu est choisi par le joueur.
//
// La grille en argument doit être sous forme de pointeur.
func generer_masque (grille *[TAILLE + 2][TAILLE + 1]int) [TAILLE][TAILLE]bool {

	masque := [TAILLE][TAILLE]bool{}

	for line := 0; line < TAILLE; line++ {
		for column := 0; column < TAILLE; column++ {
			
			if grille[line][column] != 0 {
				masque[line][column] = true
			} else {
				masque[line][column] = false
			}
		}
	}
	return masque
}

// print_grille permet d'afficher une grille de sudoku ainsi que ses compteurs.
// Si string_output=true, alors la grille n'est pas affichée mais renvoyée en string.
//
// La grille en argument doit être sous forme de pointeur.
func print_grille(grille *[TAILLE + 2][TAILLE + 1]int, string_output bool) string {

	text := ""
	for line := 0; line < TAILLE; line++ {

		// séparateur de bloc horizontale
		if line%int(math.Sqrt(float64(TAILLE))) == 0 && line != 0 {
			for i := 0; i < (int(math.Sqrt(float64(TAILLE)))+1)*int(math.Sqrt(float64(TAILLE))); i++ {
				text = text + "-"
			}
			text = text + "\n"
		}

		for column := 0; column < TAILLE; column++ {

			// séparateur de bloc vertical
			if column%int(math.Sqrt(float64(TAILLE))) == 0 && column != 0 {
				text = text + "|"
			}

			if grille[line][column] == 0 {
				text = text + "."

			} else {

				// conversion en ASCII pour les valeurs alphabétiques
				// note : rune() utilisé à la place de string() car sinon incompatible avec %c
				if grille[line][column] > 9 {
					text = text + string(grille[line][column]-10+65)
				} else {
					text = text + string(grille[line][column]+48)
				}
			}
		}

		// Compteur de ligne
		text = text + " (" + strconv.Itoa(grille[line][TAILLE]) + ")\n"

	}

	text = text + "\n"

	// compteur de colonne + bloc
	for line := TAILLE; line < TAILLE+2; line++ {
		for column := 0; column < TAILLE; column++ {

			if column%int(math.Sqrt(float64(TAILLE))) == 0 && column != 0 {

				text += "|"
			}
			text += strconv.Itoa(grille[line][column])
		}
		text += "\n"
	}

	text += "Nombre d'indices total :" + strconv.Itoa(grille[TAILLE][TAILLE])

	if string_output {
		return text
	} else {
		fmt.Println(text)
		return ""
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
func algo_backtracking(grille *[TAILLE + 2][TAILLE + 1]int, line int, column int, possibilite *[TAILLE][TAILLE][]int) bool {

	// Condition d'arrêt : lorsque l'on arrive à la dernière case TAILLE-1:TAILLE-1
	if line >= TAILLE || column >= TAILLE {
		return true

		// Si la case n'est pas vide alors on prend la suivante
	} else if grille[line][column] != 0 {
		return algo_backtracking(grille, line+(column+1)/TAILLE, (column+1)%TAILLE, possibilite)

	} else {

		for _, value := range possibilite[line][column] {
			if utils.IsOkayCase(line, column, *grille, TAILLE, value) {
				// Si la valeur est possible on l'attribue à la case
				grille[line][column] = value

				if algo_backtracking(grille, line+(column+1)/TAILLE, (column+1)%TAILLE, possibilite) {
					return true
				}

			}
		}
		grille[line][column] = 0
		return false
	}
}

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
func init_grille (grille *[TAILLE + 2][TAILLE + 1]int) ([TAILLE + 2][TAILLE + 1]int,[TAILLE][TAILLE][]int, [TAILLE][TAILLE]bool) {
	
	// générer la grille
	// mettre à jour ses compteurs
	maj_compteurs(grille)
	// générer la grille de possibilité
	// renvoyer les deux
	return *grille, generer_possibilite(grille), generer_masque(grille)
}

// resolution permet d'appeler les fonctions de résolution de grille sudoku.
// A partir d'une grille et de ses possibilites, si print_nb_solution est vrai : renvoi le nombre de solutions possibles.
// La grille ne sera pas modifiée. Si print_nb_solution est faux, la grille sera résolue selon la première solution trouvée.
// 
// La grille en argument doit être sous forme de pointeur.
func resolution (grille *[TAILLE + 2][TAILLE + 1]int, possibilite *[TAILLE][TAILLE][]int, print_nb_solution bool) int {

	if print_nb_solution {
		return nombre_solutions(grille, 0, 0, possibilite)
	} else {
		algo_backtracking(grille, 0, 0, possibilite)
		return 0
	}
}

// afficher_temps permet d'afficher une durée efficaement.
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

func test() {

	grille, possibilite , _ := init_grille(utils.Grille_sudoku_exemple(TAILLE))		// REMPLACER _ PAR masque QUAND ON L'UTILISE

	fmt.Println("GRILLE DE DEPART : ")
	print_grille(&grille, false)
	fmt.Println()

	before := time.Now()
	// fmt.Println("Nombre de solutions : ", resolution(&grille, &possibilite, true))	// CALCUL DU NB DE SOLUTIONS
	// resolution(&grille, &possibilite, false)		// RESOLUTION DE LA GRILLE
	after := time.Now()

	afficher_temps(after.Sub(before))

	print_grille(&grille, false)
	
	// fmt.Print(grille_to_string(&grille))

}

// github pour faire du sdl en go veandco go sdl 2
// demander pour le "go run UTILS.GO MAIN.GO"
// et l'histoire des packages
