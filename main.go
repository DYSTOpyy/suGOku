package main // Necessaire pour build

import (
	"flag"
	"fmt" // I/O
	"math/rand"
	"os"
	"time"
	"golang.org/x/exp/slices"
)

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
	}
	test := []int{1,3} 
	test=slices.Delete(test,1)
	fmt.println(test)
}

func generateFull(size int) [16][16]int {
	var (
		grid [16][16]int
	)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			grid[i][j] = rand.Int() % (size + 1)
		}
	}
	return grid
}

func writeFile(grid [16][16]int, size int) error {
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
