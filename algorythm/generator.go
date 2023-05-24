package algo

import (
	"math"
	"math/rand"
	
	"git.saussesylva.in/DYSTO_pyy/Sudoku/utils"
)

type Difficulty int

const (
	Easy Difficulty = 0
	Medium Difficulty = 1
	Hard Difficulty = 2
	Evil Difficulty = 3
)



func GenSlice(grille *[TAILLE + 2][TAILLE + 1]int) []int {
	size := utils.Size
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

func GenGridInit( grille *[TAILLE + 2][TAILLE + 1]int){
	size := utils.Size
	possibilite := utils.Generer_possibilite(grille)
	liste := []int{}
	root := int(math.Sqrt(float64(size)))
	for i := 0; i < size; i++ {
		for j:=0; j < size ; j=j+root{
			liste = append(liste, i%root+int(i/root)*size+j*(size)+j)
		}
	}
	liste = utils.ListRandomize(liste)
	Algo_backtracking(grille,&possibilite,liste)
	utils.Print_grille(grille,false)
}

func GeneratorFull(diff Difficulty) *[TAILLE + 2][TAILLE + 1]int {
	grille := [TAILLE + 2][TAILLE + 1]int{}
	GenGridInit(&grille)
	liste := GenSlice(&grille)
	possibilite := utils.Generer_possibilite(&grille)
	Algo_backtracking(&grille,&possibilite,liste)
	utils.Maj_compteurs(&grille)
	utils.Print_grille(&grille,false)
	DiggingHoles(&grille,diff)
	return &grille
}


// Retourne la fonction de Digging, le nombre de cases données et le minimum par ligne
func EditDifficulty(diff Difficulty) (func(*[TAILLE + 2][TAILLE + 1]int,*[TAILLE][TAILLE]bool,int,int) (int,int),int,int) {
	var (
		nb_case int
		min_line int
		algoDig func(*[TAILLE + 2][TAILLE + 1]int,*[TAILLE][TAILLE]bool,int,int) (int,int)
	)
	switch diff {
	case Easy:
		nb_case = rand.Int()%14+36
		min_line = 4
		algoDig = RandomDig
	case Medium:
		nb_case = rand.Int()%3+32
		min_line = 3
		algoDig = JumpingDig
	case Hard:
		nb_case = rand.Int()%3+28
		min_line = 2
		algoDig = WanderingDig
	case Evil:
		nb_case = rand.Int()%5+22
		min_line = 0
		algoDig = TopBottomDig
	}
	return algoDig, nb_case, min_line
}

func FillDiggable() *[TAILLE][TAILLE]bool{
	digTable := &[TAILLE][TAILLE]bool{}
	for i := 0; i < utils.Size; i++ {
		for j := 0; j < utils.Size; j++ {
			digTable[i][j] = true
		}
	}
	return digTable
}

func ContinueDigging(digTable *[TAILLE][TAILLE]bool) bool {
	for i := 0; i < utils.Size; i++ {
		for j := 0; j < utils.Size; j++ {
			if digTable[i][j] {
				return true
			}
		}
	}
	return false
}