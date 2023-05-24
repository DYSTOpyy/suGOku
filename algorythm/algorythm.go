package algo

import (
	"math/rand"
	"git.saussesylva.in/DYSTO_pyy/Sudoku/utils"
	"golang.org/x/exp/slices"
)

const TAILLE int = utils.TAILLE

// algo_backtracking résout n'importe quelle grille de sudoku avec la méthode du backtracking.
// Selon sa dimension, cela peut prendre du temps. La grille est parcouru en commençant par la case line:column
// qui sera généralement 0:0. Les valeurs testées sont les valeurs répertoriées dans possibilite afin d'accéler le résultat.
// Les possibilitées ne sont pas mises à jour à chaque nouvelle case insérée, sinon le temps augmente exponentiellement.
// Rien ne dis que la solution est unique : c'est la première qu'il trouve.
//
// La grille en argument doit être sous forme de pointeur, afin qu'elle puisse être directement complétée après exécution (!!).
// possibilite est également un un pointeur, puisqu'il n'est jamais modifié dedans.
func Algo_backtracking(grille *[TAILLE + 2][TAILLE + 1]int, possibilite *[TAILLE][TAILLE][]int, liste []int) bool {
	// Condition d'arrêt : lorsque la liste est vide
	size := utils.Size
	if len(liste) == 0 {
		return true
	}
	line, column := int(liste[0]/size), liste[0]%size
	for _, value := range possibilite[line][column] {
		if utils.IsOkayCase(line, column, *grille, value) {
			// Si la valeur est possible on l'attribue à la case
			grille[line][column] = value
			deletedInt := liste[0]
			liste = slices.Delete(liste, 0, 1)
			if Algo_backtracking(grille, possibilite, liste) {
				return true
			} else {
				liste = slices.Insert(liste, 0, deletedInt)
			}

		}
	}
	grille[line][column] = 0

	return false
}

func DiggingHoles( grille *[TAILLE + 2][TAILLE + 1]int, diff Difficulty) {
	AlgoDig, nb_case, min_line := EditDifficulty(diff)
	digTable := FillDiggable()
	nbRemovedCase := 1
	linToDig, colToDig := rand.Int()%utils.Size,rand.Int()%utils.Size
	digTable[linToDig][colToDig] = false
	grille[linToDig][colToDig] = 0
	utils.Maj_compteurs(grille)
	for ContinueDigging(digTable) {
		newLinToDig, newColToDig := AlgoDig(grille,digTable,linToDig,colToDig)
		if nbRemovedCase + 1 > utils.Size*utils.Size-nb_case {
			digTable[newLinToDig][newColToDig] = false
			continue
		}
		if (grille[newLinToDig][utils.Size] - 1 < min_line || grille[utils.Size][newColToDig] - 1 < min_line){
			digTable[newLinToDig][newColToDig] = false
			continue
		}
		Copy := *grille
		if IsSolutionUnique(Copy,newLinToDig,newColToDig) {
			linToDig,colToDig = newLinToDig, newColToDig
			digTable[linToDig][colToDig] = false
			grille[linToDig][colToDig] = 0
			nbRemovedCase += 1
			utils.Maj_compteurs(grille)
		}
	}
}

// Renvoie l'indice de ligne et de la colonne à miner, determiner par hasard
func RandomDig(grille *[TAILLE + 2][TAILLE + 1]int, digTable *[TAILLE][TAILLE]bool, linPrev int, colPrev int) (int, int) {
	linNew := rand.Int() % utils.Size
	colNew := rand.Int() % utils.Size
	for !digTable[linNew][colNew] {
		linNew = rand.Int() % utils.Size
		colNew = rand.Int() % utils.Size
	}
	return linNew, colNew
}

// Renvoie l'indice de ligne et de la colonne à miner, determiner par l'algorythme Jumping One ~ 10 milliseconds
func JumpingDig(grille *[TAILLE + 2][TAILLE + 1]int, digTable *[TAILLE][TAILLE]bool, linPrev int, colPrev int) (int, int) {
	return linPrev,colPrev
}

// Renvoie l'indice de ligne et de la colonne à miner, determiner par l'algorythme Wandering along S
func WanderingDig(grille *[TAILLE + 2][TAILLE + 1]int, digTable *[TAILLE][TAILLE]bool, linPrev int, colPrev int) (int, int) {
	return linPrev,colPrev
}

// Renvoie l'indice de ligne et de la colonne à miner, determiner par l'algorythme Left to Right then Top to Bottom
func TopBottomDig(grille *[TAILLE + 2][TAILLE + 1]int, digTable *[TAILLE][TAILLE]bool, linPrev int, colPrev int) (int, int) {
	return linPrev,colPrev
}

func IsSolutionUnique(grille [TAILLE + 2][TAILLE + 1]int, linToDig int,colToDig int) bool{
	listValue := []int{}
	for i := 1; i < utils.Size+1; i++ {
		if i!= grille[linToDig][colToDig]{
			listValue = append(listValue, i)
		}
	}
	for _,value := range listValue {
		if utils.IsOkayCase(linToDig,colToDig,grille,value){
			grille[linToDig][colToDig] = value
			possibilite := utils.Generer_possibilite(&grille)
			slice := GenSlice(&grille)
			if Algo_backtracking(&grille,&possibilite,slice){
				return false
			}
		}
	}
	return true
}