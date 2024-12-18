package floor

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	switch configuration.Global.FloorKind {
	case FromFileFloor:
		f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
	case QuadTreeFloor:
		f.quadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FloorFile))
	}
}

// lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
func readFloorFromFile(fileName string) (floorContent [][]int) {
	// Ouvrir le fichier
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Scanner le fichier ligne par ligne
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // lire une ligne
		row := []int{}         // initialiser une ligne du tableau

		// Parcourir chaque caractère de la ligne et le convertir en entier
		for _, char := range line {
			cell, err := strconv.Atoi(string(char)) // conversion en entier
			if err != nil {
				log.Fatalf("Erreur de conversion dans le fichier %s: %v", fileName, err)
				return
			}
			row = append(row, cell) // ajouter l'entier à la ligne
		}

		// Ajouter la ligne au tableau floorContent
		floorContent = append(floorContent, row)
	}	
	return floorContent
}
