package floor

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"fmt"
	"math/rand"
	"time"
	
	
	
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
	
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() 
		row := []int{}         // initialiser une ligne du tableau
		for _, char := range line {
			cell, err := strconv.Atoi(string(char)) // conversion en entier
			if err != nil {
				log.Fatalf("Erreur de conversion dans le fichier %s: %v", fileName, err)
				return
			}
			row = append(row, cell) // ajouter l'entier à la ligne
		}

		
		floorContent = append(floorContent, row)
	}	
	return floorContent
}

// GenerateNewChunk génère un nouveau chunk de terrain en fonction de la direction donnée.
// La direction peut être :
//   - 0 : à droite
//   - 1 : en bas
//   - 2 : à gauche
//   - 3 : en haut
func (f *Floor) GenerateNewChunk(direction int) {
	// Convertir le quadtree en tableau pour accéder à ses données
	currentArray := f.quadtreeContent.ToArray()

	// Obtenir la largeur et la hauteur du fichier de terrain
	widhtFile, heightFile := getWidhtHeightOfFile(configuration.Global.FloorFile)

	// Générer un nouveau chunk en fonction de la direction
	switch direction {
	case 1: // Bas
		// Générer un nouveau chunk en bas du terrain actuel
		var newChunkSud [][]int
		if configuration.Global.RandomGenExtension {
			// Générer un chunk aléatoire si la configuration le permet
			newChunkSud = generateRandomFloor(len(currentArray[1]), configuration.Global.RandomHeight)
		} else {
			// Copier les lignes du fichier de terrain pour générer le nouveau chunk
			for i := 0; i < heightFile; i++ {
				newChunkSud = append(newChunkSud, currentArray[i])
			}
		}
		// Ajouter le nouveau chunk au terrain actuel
		currentArray = append(currentArray, newChunkSud...)
		// Mettre à jour la hauteur du terrain
		configuration.Global.FloorHeigth = len(currentArray)

	case 0: // Droite
		// Générer un nouveau chunk à droite du terrain actuel
		var newChunkEst [][]int
		if configuration.Global.RandomGenExtension {
			// Générer un chunk aléatoire si la configuration le permet
			newChunkEst = generateRandomFloor(configuration.Global.RandomWidht, len(currentArray))
		} else {
			// Copier les lignes du fichier de terrain pour générer le nouveau chunk
			for i := 0; i < len(currentArray)/heightFile; i++ {
				newChunkEst = append(newChunkEst, readFloorFromFile(configuration.Global.FloorFile)...)
			}
		}
		// Ajouter les colonnes du nouveau chunk au terrain actuel
		if configuration.Global.RandomGenExtension {
			for i := 0; i < len(currentArray); i++ {
				for j := 0; j < configuration.Global.RandomWidht; j++ {
					currentArray[i] = append(currentArray[i], newChunkEst[i][j])
				}
			}
		} else {
			for i := 0; i < len(currentArray); i++ {
				for j := 0; j < widhtFile; j++ {
					currentArray[i] = append(currentArray[i], newChunkEst[i][j])
				}
			}
		}
		// Mettre à jour la largeur du terrain
		configuration.Global.FloorWidth = len(currentArray[0])

	case 2: // Gauche
		// Générer un nouveau chunk à gauche du terrain actuel
		var newChunkOuest [][]int
		if configuration.Global.RandomGenExtension {
			// Générer un chunk aléatoire si la configuration le permet
			newChunkOuest = generateRandomFloor(configuration.Global.RandomWidht, len(currentArray))
		} else {
			// Copier les lignes du fichier de terrain pour générer le nouveau chunk
			for i := 0; i < len(currentArray)/heightFile; i++ {
				newChunkOuest = append(newChunkOuest, readFloorFromFile(configuration.Global.FloorFile)...)
			}
		}
		// Ajouter les colonnes du nouveau chunk au terrain actuel
		for i := 0; i < len(currentArray); i++ {
			currentArray[i] = append(newChunkOuest[i], currentArray[i]...)
		}
		// Mettre à jour la largeur du terrain
		configuration.Global.FloorWidth = len(currentArray[0])

	case 3: // Haut
		// Générer un nouveau chunk en haut du terrain actuel
		var newChunkNord [][]int
		if configuration.Global.RandomGenExtension {
			// Générer un chunk aléatoire si la configuration le permet
			newChunkNord = generateRandomFloor(len(currentArray[1]), configuration.Global.RandomHeight)
		} else {
			for i := 0; i < heightFile; i++ {
				newChunkNord = append(newChunkNord, currentArray[i])
			}
		}
		currentArray = append(newChunkNord, currentArray...)
		configuration.Global.FloorHeigth = len(currentArray)
	default:
		panic("direction invalide")
	}
	f.quadtreeContent = quadtree.MakeFromArray(currentArray)
	
}

// getWidhtHeightOfFile lit un fichier et retourne la largeur et la hauteur du contenu du fichier.
// La largeur est déterminée par le nombre de caractères dans la première ligne du fichier,
// et la hauteur est déterminée par le nombre de lignes dans le fichier.
func getWidhtHeightOfFile(fileName string) (width, height int) {
	file, err := os.Open(fileName)
	if err != nil {
		
		fmt.Println(err)
		return
	}
	defer file.Close() 

	
	scanner := bufio.NewScanner(file)

	// Initialiser les variables pour stocker la largeur et la hauteur
	var ligneCount int
	var colonneCount int
	
	for scanner.Scan() {
		ligneCount++

		// Lire la première ligne pour déterminer la largeur
		ligne := scanner.Text()
		if ligneCount == 1 {
			// La largeur est déterminée par le nombre de caractères dans la première ligne
			colonneCount = len(ligne)
		}
	}


	return colonneCount, ligneCount
}

// generateRandomFloor génère un tableau de taille aléatoire avec des valeurs aléatoires.
// La fonction prend en paramètre la largeur et la hauteur du tableau à générer.
func generateRandomFloor(width, height int) [][]int {
	// Initialiser le générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	// Créer un tableau de taille width x height
	floorContent := make([][]int, height)
	for i := range floorContent {
		floorContent[i] = make([]int, width)
	}

	// Remplir le tableau avec des valeurs aléatoires
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Générer un nombre aléatoire entre 0 et 5
			floorContent[y][x] = rand.Intn(5)
		}
	}

	
	return floorContent
}