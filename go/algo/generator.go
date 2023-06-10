package algo

import (
	"math"
	"math/rand"

	"git.saussesylva.in/DYSTO_pyy/Sudoku/utils"
)

type Difficulty int

const (
	Easy   Difficulty = 1
	Medium Difficulty = 2
	Hard   Difficulty = 3
	Evil   Difficulty = 4
)

const MAX int32 = utils.MAX

// Retourne la fonction de Digging, le nombre de cases données et le minimum par ligne
func EditDifficulty(diff Difficulty) (func(*[MAX][MAX]bool, int32, int32) (int32, int32), int32, int) {
	var (
		nb_case  int32
		min_ligne int
		algoDig  func(*[MAX][MAX]bool, int32, int32) (int32, int32)
	)
	switch diff {
	case Easy:
		nb_case = rand.Int31()%14 + 36
		min_ligne = 4
		algoDig = RandomDig
	case Medium:
		nb_case = rand.Int31()%3 + 32
		min_ligne = 3
		algoDig = JumpingDig
	case Hard:
		nb_case = rand.Int31()%3 + 28
		min_ligne = 2
		algoDig = WanderingDig
	case Evil:
		nb_case = rand.Int31()%5 + 22
		min_ligne = 0
		algoDig = TopBottomDig
	}
	return algoDig, nb_case, min_ligne
}

// Géneration complète d'une grille en fonction d'un paramètre de difficulté
func GeneratorTotal(diff Difficulty) *[MAX + 2][MAX + 1]int {
	grille := [MAX + 2][MAX + 1]int{}
	GenGridFull(&grille)
	liste := Generer_Slice(&grille)
	possibilite := utils.Generer_possibilite(&grille)
	Algo_backtracking(&grille, &possibilite, liste)
	utils.Maj_compteurs(&grille)
	DiggingHoles(&grille, diff)
	return &grille
}

// init_grille génère une grille puis détermine ses compteurs et ses possibilité.
// C'est une fonction à retour multiple, qui renvoie la grille et les possibilités.
//
// La grille en argument doit être sous forme de pointeur.		DOIT DISPARAITRE POUR LAISSER PLACE A UN ARGUMENT "DIFFICULTE"
func Init_grille(difficulteChoisie Difficulty) ([MAX + 2][MAX + 1]int, [MAX][MAX][]int, [MAX][MAX]bool, [MAX][MAX]bool) {
	// générer la grille
	grille := GeneratorTotal(difficulteChoisie)
	// mettre à jour ses compteurs
	utils.Maj_compteurs(grille)
	// générer la grille de possibilité
	possibilite := utils.Generer_possibilite(grille)
	// generer le masque de la grille
	masque := utils.Generer_masque(grille)
	// generer la grille des cases vérifiées
	verifier := [MAX][MAX]bool{}
	// renvoyer les deux
	return *grille, possibilite, masque, verifier
}

// Permet de créer une liste d'entier allant de 0 à la taille de la grille au carré, en enlevant les indices des cases déja occupées
func Generer_Slice(grille *[MAX + 2][MAX + 1]int) []int32 {
	size := utils.Taille
	liste := []int32{}
	for ligne := int32(0); ligne < size; ligne++ {
		for colonne := int32(0) ; colonne < size; colonne++ {
			if grille[ligne][colonne] == 0 {
				liste = append(liste, size * ligne + colonne)
			}
		}
	}
	return liste
}

// Initialise la grille en remplissant les carrés présent sur la diagonale de la grille avec des valeurs aléatoires
func GenGridFull(grille *[MAX + 2][MAX + 1]int) {
	size := utils.Taille
	possibilite := utils.Generer_possibilite(grille)
	liste := []int32{}
	racine := int32(math.Sqrt(float64(size)))
	for i := int32(0); i < size; i++ {
		for j := int32(0); j < size; j = j + racine {
			liste = append(liste, i % racine + int32(i/racine) * size + j * (size+1))
		}
	}
	liste = utils.ListRandomize(liste)
	Algo_backtracking(grille, &possibilite, liste)
}

// Algorythme Top-Down qui permet de vider une grille en fonction de la difficulté choisie
func DiggingHoles(grille *[MAX + 2][MAX + 1]int, diff Difficulty) {
	AlgoDig, nb_case, min_ligne := EditDifficulty(diff)
	digTable := FillDiggable()
	var nbRemovedCase int32 = 1
	taille := utils.Taille
	linToDig, colToDig := rand.Int31()%taille, rand.Int31()%taille
	Dig(grille, digTable, linToDig, colToDig)
	for nbRemovedCase < taille * taille - nb_case && ContinueDigging(digTable) {
		newLinToDig, newColToDig := AlgoDig(digTable, linToDig, colToDig)
		if grille[newLinToDig][taille]-1 < min_ligne || grille[taille][newColToDig]-1 < min_ligne {
			digTable[newLinToDig][newColToDig] = false
			continue
		}
		Copy := *grille
		if IsSolutionUnique(Copy, newLinToDig, newColToDig) {
			linToDig, colToDig = newLinToDig, newColToDig
			Dig(grille, digTable, linToDig, colToDig)
			nbRemovedCase += 1

		} else {
			digTable[newLinToDig][newColToDig] = false
		}
	}
}

// Permet d'initaliser la DigTable avec true dans toutes ses cases
func FillDiggable() *[MAX][MAX]bool {
	digTable := &[MAX][MAX]bool{}
	for i := int32(0); i < utils.Taille; i++ {
		for j := int32(0); j < utils.Taille; j++ {
			digTable[i][j] = true
		}
	}
	return digTable
}

// Renvoie l'indice de ligne et de la colonne à miner, determiner par hasard
func RandomDig(digTable *[MAX][MAX]bool, linPrev int32, colPrev int32) (int32, int32) {
	taille := utils.Taille
	linNew := rand.Int31() % taille
	colNew := rand.Int31() % taille
	for !digTable[linNew][colNew] {
		linNew = rand.Int31() % taille
		colNew = rand.Int31() % taille
	}
	return linNew, colNew
}

// Renvoie l'indice de ligne et de la colonne à miner, determiner par l'algorythme Jumping One
func JumpingDig(digTable *[MAX][MAX]bool, linPrev int32, colPrev int32) (int32, int32) {
	linNew, colNew := linPrev, colPrev
	ToLeft := linNew % 2
	for !digTable[linNew][colNew] {
		if ToLeft == 1 {
			colNew -= 2
			if colNew < 0 {
				colNew += 1 + 2*((colNew+3)%2)
				linNew += 1
				if linNew >= utils.Taille {
					linNew = 0
					colNew = 0 + (colNew % 2)
				}
			}
		} else {
			colNew += 2
			if colNew >= utils.Taille {
				colNew -= 1 + 2*((colNew+1)%2)
				linNew += 1
				if linNew >= utils.Taille {
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
func WanderingDig(digTable *[MAX][MAX]bool, linPrev int32, colPrev int32) (int32, int32) {
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
			if colNew >= utils.Taille {
				colNew = utils.Taille - 1
				linNew += 1
			}
		}
		if linNew >= utils.Taille {
			linNew = 0
			colNew = 0
		}
		ToLeft = linNew % 2
	}
	return linNew, colNew
}

// Renvoie l'indice de ligne et de la colonne à miner, determiner par l'algorythme Left to Right then Top to Bottom
func TopBottomDig(digTable *[MAX][MAX]bool, linPrev int32, colPrev int32) (int32, int32) {
	linNew, colNew := linPrev, colPrev
	for !digTable[linNew][colNew] {
		colNew += 1
		if colNew >= utils.Taille {
			colNew -= utils.Taille
			linNew += 1
			if linNew >= utils.Taille {
				linNew = 0
			}
		}
	}
	return linNew, colNew
}

// Verifie l'unicité d'une solution et renvoie true si la solution est unique, false sinon
func IsSolutionUnique(grille [MAX + 2][MAX + 1]int, linToDig int32, colToDig int32) bool {
	listValue := []int{}
	for i := 1; i < int(utils.Taille + 1); i++ {
		if i != grille[linToDig][colToDig] {
			listValue = append(listValue, i)
		}
	}
	for _, value := range listValue {
		if utils.CaseBonne(linToDig, colToDig, grille, value) {
			grille[linToDig][colToDig] = value
			possibilite := utils.Generer_possibilite(&grille)
			slice := Generer_Slice(&grille)
			if Algo_backtracking(&grille, &possibilite, slice) {
				return false
			}
		}
	}
	return true
}

// Permet de miner la case indiqué, en mettant à jour la grille et la DigTable
func Dig(grille *[MAX + 2][MAX + 1]int, digTable *[MAX][MAX]bool, linToDig int32, colToDig int32) {
	digTable[linToDig][colToDig] = false
	grille[linToDig][colToDig] = 0
	utils.Maj_compteurs(grille)
}

// Fonction qui renvoie un boolean indiquant si il reste encore des cases à miner dans notre grille
func ContinueDigging(digTable *[MAX][MAX]bool) bool {
	for i := int32(0); i < utils.Taille; i++ {
		for j := int32(0); j < utils.Taille; j++ {
			if digTable[i][j] {
				return true
			}
		}
	}
	return false
}
