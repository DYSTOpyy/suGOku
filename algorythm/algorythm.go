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
	if len(liste) == 0 {
		return true
	}
	line, column := utils.IndexToLinCol(liste[0])
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

func DiggingHoles(grille *[TAILLE + 2][TAILLE + 1]int, diff Difficulty) {
	AlgoDig, nb_case, min_line := EditDifficulty(diff)
	digTable := FillDiggable()
	nbRemovedCase := 1
	linToDig, colToDig := rand.Int()%utils.Size,rand.Int()%utils.Size
	Dig(grille,digTable,linToDig,colToDig)
	for nbRemovedCase < utils.Size * utils.Size - nb_case && ContinueDigging(digTable) {
		newLinToDig, newColToDig := AlgoDig(digTable,linToDig,colToDig)
		if (grille[newLinToDig][utils.Size] - 1 < min_line || grille[utils.Size][newColToDig] - 1 < min_line){
			digTable[newLinToDig][newColToDig] = false
			continue
		}
		Copy := *grille
		if IsSolutionUnique(Copy,newLinToDig,newColToDig) {
			linToDig,colToDig = newLinToDig, newColToDig
			Dig(grille,digTable,linToDig,colToDig)
			nbRemovedCase += 1
			
		} else {
			digTable[newLinToDig][newColToDig] = false
		}
	}
}

// Renvoie l'indice de ligne et de la colonne à miner, determiner par hasard 
func RandomDig(digTable *[TAILLE][TAILLE]bool, linPrev int, colPrev int) (int, int) {
	linNew := rand.Int() % utils.Size
	colNew := rand.Int() % utils.Size
	for !digTable[linNew][colNew] {
		linNew = rand.Int() % utils.Size
		colNew = rand.Int() % utils.Size
	}
	return linNew, colNew
}

// Renvoie l'indice de ligne et de la colonne à miner, determiner par l'algorythme Jumping One 

func JumpingDig(digTable *[TAILLE][TAILLE]bool, linPrev int, colPrev int) (int, int) {
	linNew, colNew := linPrev, colPrev
	ToLeft := linNew % 2
	for !digTable[linNew][colNew] {
		if ToLeft == 1{
			colNew -= 2
			if colNew < 0 {
				colNew += 1 + 2 * ((colNew + 3) % 2)
				linNew += 1
				if linNew >= utils.Size{
					linNew = 0
					colNew = 0 + (colNew % 2)
				}
			}
		} else {
			colNew += 2
			if colNew >= utils.Size{
				colNew -= 1 + 2 * ((colNew + 1) % 2)
				linNew += 1
				if linNew >= utils.Size{
					linNew = 0
					colNew = 0 + (colNew % 2)
				}
			}
		}
		ToLeft = linNew % 2
	}
	return linNew, colNew 
}

// Renvoie l'indice de ligne et de la colonne à miner, determiner par l'algorythme Wandering along S
func WanderingDig(digTable *[TAILLE][TAILLE]bool, linPrev int, colPrev int) (int, int) {
	linNew, colNew := linPrev, colPrev
	ToLeft := linNew % 2
	for !digTable[linNew][colNew] {
		if ToLeft == 1 {
			colNew -= 1
			if colNew < 0 {
				colNew = 0
				linNew += 1				
			}
		} else {
			colNew += 1
			if colNew >= utils.Size {
				colNew = utils.Size - 1
				linNew += 1
			}
		}
		if linNew >= utils.Size{
			linNew = 0
			colNew = 0
		}
		ToLeft = linNew % 2
	}
	return linNew , colNew
}

// Renvoie l'indice de ligne et de la colonne à miner, determiner par l'algorythme Left to Right then Top to Bottom
func TopBottomDig(digTable *[TAILLE][TAILLE]bool, linPrev int, colPrev int) (int, int) {
	linNew, colNew := linPrev, colPrev
	for !digTable[linNew][colNew] {
		colNew +=  1
		if colNew >= utils.Size{
			colNew -= utils.Size
			linNew += 1
			if linNew >= utils.Size{
				linNew = 0
			}
		}
	}
	return linNew, colNew
}

func IsSolutionUnique(grille [TAILLE + 2][TAILLE + 1]int, linToDig int,colToDig int) bool{
	listValue := []int{}
	for i := 1; i < utils.Size + 1; i++ {
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

func Dig (grille *[TAILLE + 2][TAILLE + 1]int, digTable *[TAILLE][TAILLE]bool, linToDig int, colToDig int) {
	digTable[linToDig][colToDig] = false
	grille[linToDig][colToDig] = 0
	utils.Maj_compteurs(grille)
}