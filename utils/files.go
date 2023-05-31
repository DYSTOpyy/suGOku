package utils

import (
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

// Permet de revenir
func GetPackagePath() string {
	Path,_ := filepath.Abs("test/dir/test2")
	for filepath.Base(Path) != "Sudoku" {
		Path,_ = filepath.Split(Path)
		Path = filepath.Clean(Path)
	}
	return Path
}


func SaveFile(grid [TAILLE+2][TAILLE+1]int, mask [TAILLE][TAILLE]bool) error {
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
			char := fmt.Sprint(mask[i][j])
			if char == "true" {
				char = "1"
			} else {
				char = "0"
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

func ImportFile() (*[TAILLE+2][TAILLE+1]int, error){
	PackagePath := GetPackagePath()
	ImportPath := filepath.Join(PackagePath,"files/import.txt")
	buffer,err := os.ReadFile(ImportPath)
	nbChar := len(buffer)
	fmt.Println(nbChar)
	if err != nil {
		return nil, err
	}
	root := int(math.Sqrt(float64(nbChar))) 
	if root * root != nbChar {
		return nil, errors.New("SizeError")
	}
	Size = root
	grille := [TAILLE + 2][TAILLE + 1]int{}
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			char := string(buffer[i*Size+j])
			if char == "." {
				char = "0"
			}
			
			Case, err := strconv.ParseInt(char,17,0)
			if err != nil {
				return nil, errors.New("ConversionError")
			}
			grille[i][j] = int(Case)
		}
	}
	Maj_compteurs(&grille)
	return &grille,nil
}