package utils

import (
	"errors"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

// Renvoie le chemin absolu du dossier racine du projet
func GetPackagePath() string {
	Path,_ := filepath.Abs("")
	for filepath.Base(Path) != "Sudoku" {
		Path,_ = filepath.Split(Path)
		Path = filepath.Clean(Path)
	}
	return Path
}

// Enregistre une grille dans un fichier save.txt, présent dans le dossier files, sous la forme de 2 blocs de chiffres :
// le premier pour la grille
// et le second pour le masque
// Ce programme peut renvoyer une erreur de type PathError ou une erreur d'écriture
func SaveFile(grid *[TAILLE+2][TAILLE+1]int, mask [TAILLE][TAILLE]bool) error {
	PackagePath := GetPackagePath()
	FilePath := filepath.Join(PackagePath,"files/save.txt")
	file, err := os.Create(FilePath)
	if err != nil {
		return err
	}
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			char := strconv.FormatInt(int64(grid[i][j]),17)
			if char == "0" {
				char = "."
			}
			_, err2 := file.WriteString(char)
			if err2 != nil {
				file.Close()
				return err2
			}
		}
	}
	_, err2 := file.WriteString(" ")
	if err2 != nil {
		file.Close()
		return err2
	}
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			char := "0"
			if mask[i][j] {
				char = "1"
			}
			_, err2 := file.WriteString(char)
			if err2 != nil {
				file.Close()
				return err2
			}
		}
	}
	return nil
}

// Permet d'importer un sudoku totalemant neuf, ou de charger une sauvegarde
// Le programme renvoie la grille, le masque associé ainsi qu'une erreur
// Peut renvoyer une erreur d'ouverture, de taille ou de conversion
func ImportFile() (*[TAILLE+2][TAILLE+1]int, *[TAILLE][TAILLE]bool, error){
	var ImportPath string
	PackagePath := GetPackagePath()
	ImportPath = filepath.Join(PackagePath,"files/save.txt")
	buffer,err := os.ReadFile(ImportPath)
	if err != nil {
		return nil, nil, err
	}
	nbChar := len(buffer)
	var root int
	var isSave bool
	if  root = int(math.Sqrt(float64(nbChar))) ; root * root == nbChar {
		isSave = false
		Size = root
	} else if root = int(math.Sqrt(float64((nbChar-1)/2))) ; root * root * 2 + 1 == nbChar{
		Size = root
		isSave = true
	} else {
		return nil, nil, errors.New("SizeError")
	}
	grille := [TAILLE + 2][TAILLE + 1]int{}
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			char := string(buffer[i*Size+j])
			if char == "." {
				char = "0"
			}
			Case, err := strconv.ParseInt(char,17,0)
			if err != nil {
				return nil, nil,err
			}
			grille[i][j] = int(Case)
		}
	}
	mask := [TAILLE][TAILLE]bool{}
	if isSave {
		for i := 0; i < Size; i++ {
			for j := 0; j < Size; j++ {
				char := string(buffer[(Size * Size + 1) + i * Size + j])
				if char == "0" {
					mask[i][j] = false
				} else {
					mask[i][j] = true
				}
			}
		}
	} else {
		mask = Generer_masque(&grille)
	}
	Maj_compteurs(&grille)
	return &grille,&mask,nil
}

