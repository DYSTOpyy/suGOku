package algo

import (
	"git.saussesylva.in/DYSTO_pyy/Sudoku/utils"
)

// algo_backtracking résout n'importe quelle grille de sudoku avec la méthode du backtracking.
// Selon sa dimension, cela peut prendre du temps. La grille est parcouru en commençant par la case line:column
// qui sera généralement 0:0. Les valeurs testées sont les valeurs répertoriées dans possibilite afin d'accéler le résultat.
// Les possibilitées ne sont pas mises à jour à chaque nouvelle case insérée, sinon le temps augmente exponentiellement.
// Rien ne dis que la solution est unique : c'est la première qu'il trouve.
//
// La grille en argument doit être sous forme de pointeur, afin qu'elle puisse être directement complétée après exécution (!!).
// possibilite est également un un pointeur, puisqu'il n'est jamais modifié dedans.
func algo_backtracking(grille *[utils.TAILLE + 2][utils.TAILLE + 1]int, line int, column int, possibilite *[utils.TAILLE][utils.TAILLE][]int) bool {

	// Condition d'arrêt : lorsque l'on arrive à la dernière case TAILLE-1:TAILLE-1
	if line >= utils.TAILLE || column >= utils.TAILLE {
		return true

		// Si la case n'est pas vide alors on prend la suivante
	} else if grille[line][column] != 0 {
		return algo_backtracking(grille, line+(column+1)/utils.TAILLE, (column+1)%utils.TAILLE, possibilite)

	} else {

		for _, value := range possibilite[line][column] {
			if utils.IsOkayCase(line, column, *grille, utils.TAILLE, value) {
				// Si la valeur est possible on l'attribue à la case
				grille[line][column] = value

				if algo_backtracking(grille, line+(column+1)/utils.TAILLE, (column+1)%utils.TAILLE, possibilite) {
					return true
				}

			}
		}
		grille[line][column] = 0
		return false
	}
}