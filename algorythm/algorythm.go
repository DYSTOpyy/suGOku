package algo

import (
	"math/rand"
	"git.saussesylva.in/DYSTO_pyy/Sudoku/utils"
	"golang.org/x/exp/slices"
)

type Difficulty int

const (
	Easy Difficulty = 0
	Medium Difficulty = 1
	Hard Difficulty = 2
	Impossible Difficulty = 3
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
func Algo_backtracking(grille *[TAILLE + 2][TAILLE + 1]int, possibilite *[TAILLE][TAILLE][]int, liste []int, size int) bool {
	// Condition d'arrêt : lorsque l'on arrive à la dernière case TAILLE-1:TAILLE-1
	if len(liste) == 0 {
		return true

		// Si la case n'est pas vide alors on prend la suivante
	}
	line,column := int(liste[0]/size),liste[0]%size 
	for _, value := range possibilite[line][column] {
		if utils.IsOkayCase(line, column, *grille, size, value) {
			// Si la valeur est possible on l'attribue à la case
			grille[line][column] = value
			deletedInt:=liste[0]
			liste=slices.Delete(liste,0,1)
			if Algo_backtracking(grille, possibilite,liste, size) {
				return true
			} else {
				liste = slices.Insert(liste,0,deletedInt)
			}

		}
	}
	grille[line][column] = 0
	
	return false
}


func GenSlice(size int, grille [TAILLE + 2][TAILLE + 1]int) []int {
	liste := make([]int,0)
	for line := 0; line < size; line++ {
		for column := 0; column < size; column++ {
			if grille[line][column] == 0 {
				liste = append(liste, size*line+column)
			}
		}
	}
	return liste
}

func GeneratorFull(size int, diff Difficulty) *[TAILLE + 2][TAILLE + 1]int {
	liste := []int{0}
	
	for i:=1;i < size*size;i++ {
		random := rand.Int()%(size*size)
		liste=slices.Insert(liste,random%len(liste),i)
	}
	grille := [TAILLE + 2][TAILLE + 1]int{}
	possibilite := utils.Generer_possibilite(&grille,size)
	Algo_backtracking(&grille,&possibilite,liste,size)
	return &grille
}
// A voir plus tard
func EditDifficulty(liste []int, diff Difficulty) []int{

	switch diff {
	case Easy:
		for i := 0; i < 31; i++ {
			randInt := rand.Int()%len(liste)
			liste=slices.Delete(liste,randInt,randInt+1)
		}
		return liste[:50]
	case Medium:
		for i := 0; i < 46; i++ {
			randInt := rand.Int()%len(liste)
			liste=slices.Delete(liste,randInt,randInt+1)
		}
		return liste[:35]
	case Hard:
		for i := 0; i < 50; i++ {
			randInt := rand.Int()%len(liste)
			liste=slices.Delete(liste,randInt,randInt+1)
		}
		return liste[:31]
	case Impossible:
		for i := 0; i < 54; i++ {
			randInt := rand.Int()%len(liste)
			liste=slices.Delete(liste,randInt,randInt+1)
		}
		return liste[:27]
	}
	return liste
}