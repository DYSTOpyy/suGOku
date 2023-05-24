package main// Necessaire pour build

import (
	"flag"
	"fmt" // I/O
	"math/rand"
	"os"
	"time"
	// "golang.org/x/exp/slices"
	"git.saussesylva.in/DYSTO_pyy/Sudoku/utils"
)



const maxSize int  = 16

func main() {
	size := flag.Int("size", 9, "Taille du Sudoku")
	isGen := flag.Bool("g", false, "Generator mode")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	if *isGen {
		grid := generateFull(*size)
		err := writeFile(grid, *size)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(utils.IsOkayCase(1,2,grid,3))
	}
	
	// Tests sur les slices
	// test := []int{1,3,4,7} 
	// fmt.Println(test)
	// test=slices.Delete(test,1,2)
	// fmt.Println(test)
}

func generateFull(size int) [maxSize+2][maxSize+1]int {
	var (
		grid [maxSize+2][maxSize+1]int
	)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			grid[i][j] = rand.Int() % (size + 1)
		}
	}
	return grid
}

func writeFile(grid [maxSize+2][maxSize+1]int, size int) error {
	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			char := fmt.Sprint(grid[i][j])
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
	return nil
}
