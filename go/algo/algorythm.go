package algo

import (
	"git.saussesylva.in/DYSTO_pyy/Sudoku/utils"
	"golang.org/x/exp/slices"
)

// algo_backtracking résout n'importe quelle grille de sudoku avec la méthode du backtracking.
// Selon sa dimension, cela peut prendre du temps. La grille est parcouru en commençant par la case ligne:colonne
// qui sera généralement 0:0. Les valeurs testées sont les valeurs répertoriées dans possibilite afin d'accéler le résultat.
// Les possibilitées ne sont pas mises à jour à chaque nouvelle case insérée, sinon le temps augmente exponentiellement.
// Rien ne dis que la solution est unique : c'est la première qu'il trouve.
//
// La grille en argument doit être sous forme de pointeur, afin qu'elle puisse être directement complétée après exécution (!!).
// possibilite est également un un pointeur, puisqu'il n'est jamais modifié dedans.
func Algo_backtracking(grille *[MAX + 2][MAX + 1]int, possibilite *[MAX][MAX][]int, liste []int32) bool {
	// Condition d'arrêt : lorsque la liste est vide
	if len(liste) == 0 {
		return true
	}
	ligne, colonne := utils.IndexToLinCol(liste[0])
	for _, value := range possibilite[ligne][colonne] {
		if utils.CaseBonne(ligne, colonne, *grille, value) {
			// Si la valeur est possible on l'attribue à la case
			grille[ligne][colonne] = value
			deletedInt := liste[0]
			liste = slices.Delete(liste, 0, 1)
			if Algo_backtracking(grille, possibilite, liste) {
				return true
			} else {
				liste = slices.Insert(liste, 0, deletedInt)
			}

		}
	}
	grille[ligne][colonne] = 0
	return false
}

// nombre_solutions détermine le nombre de solutions d'une grille de sudoku.
// Prend beaucoup de temps pour de grandes matrices. La grille est parcouru en commençant par la case ligne:colonne
// qui sera généralement 0:0. Les valeurs testées sont les valeurs répertoriées dans possibilite afin d'accéler le résultat.
// Les possibilitées ne sont pas mises à jour à chaque nouvelle case insérée, sinon le temps augmente exponentiellement.
//
// La grille en argument doit être sous forme de pointeur.
// possibilite est également un un pointeur, puisqu'il n'est jamais modifié dedans.
func Nombre_solutions(grille *[MAX + 2][MAX + 1]int, possibilite *[MAX][MAX][]int, liste []int32) int32 {
	ligne, colonne := utils.IndexToLinCol(liste[0])
	liste = slices.Delete(liste,0,1)
	count := int32(0)
	for _, value := range possibilite[ligne][colonne] {
		Copy := *grille
		Copy[ligne][colonne] = value
		if Algo_backtracking(&Copy,possibilite,liste){
			count += 1
		}
	}
	return count
}

// resolution permet d'appeler les fonctions de résolution de grille sudoku.
// A partir d'une grille et de ses possibilites, si print_nb_solution est vrai : renvoi le nombre de solutions possibles.
// La grille ne sera pas modifiée. Si print_nb_solution est faux, la grille sera résolue selon la première solution trouvée.
// 
// La grille en argument doit être sous forme de pointeur.
func Resolution (grille *[MAX + 2][MAX + 1]int, possibilite *[MAX][MAX][]int, afficher_nb_solution bool) int32 {	
	liste := Generer_Slice(grille)
	if afficher_nb_solution {
		return Nombre_solutions(grille, possibilite, liste)
	} else {
		if Algo_backtracking(grille, possibilite, liste){
			return 1
		} else {
			return 0
		}
	}
}

