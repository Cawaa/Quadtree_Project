package quadtree

import (
    "math/rand"
    "os"
    "bufio"
    "log"
    "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"fmt"
)
// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
//
// Paramètres:
// - floorContent: Un tableau 2D représentant le contenu du sol.
//
// Retourne:
// - Un quadtree représentant le terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	// On calcule la largeur et la hauteur du terrain
	if configuration.Global.RandomFloor {
		for i := range floorContent {
			for j := range floorContent[i] {
				floorContent[i][j] = rand.Intn(5)
			}
		}
		// Enregistrer le terrain généré dans un fichier
		err := writeFloorToFile("../floor-files/generated_floor.txt", floorContent)
		if err != nil {
			log.Fatalf("Erreur lors de l'écriture du fichier: %v", err)
	}
}

	// On obtient la hauteur du tableau
	height := len(floorContent)
	// On obtient la largeur du tableau
	width := len(floorContent[0])

	// On construit la racine de l'arbre en utilisant la fonction makeNode
	root := makeNode(floorContent, 0, 0, width, height)

	// On retourne le quadtree avec les dimensions et la racine
	q = Quadtree{
		width:  width,
		height: height,
		root:   root,
	}
	return q
}

// makeNode construit un nœud de l'arbre comme on veux donc de façon récursive
func makeNode(grid [][]int, topLeftX, topLeftY, width, height int) *node {
	//on Vérifie si la zone est homogène
	isHomogeneous, terrainType := isZoneHomogeneous(grid, topLeftX, topLeftY, width, height)

	if isHomogeneous {
		// Si la zone est homogène, on créé une feuille
		return &node{
			topLeftX: topLeftX,
			topLeftY: topLeftY,
			width:    width,
			height:   height,
			content:  terrainType,
			isLeaf:   true,
		}
	}

	// et on Divise la zone en quatre sous-zones
	halfWidth := width / 2
	halfHeight := height / 2

	// On Construit un nœud interne
	return &node{
		topLeftX:        topLeftX,
		topLeftY:        topLeftY,
		width:           width,
		height:          height,
		content:         -1,
		isLeaf:          false,
		topLeftNode:     makeNode(grid, topLeftX, topLeftY, halfWidth, halfHeight),
		topRightNode:    makeNode(grid, topLeftX+halfWidth, topLeftY, width-halfWidth, halfHeight),
		bottomLeftNode:  makeNode(grid, topLeftX, topLeftY+halfHeight, halfWidth, height-halfHeight),
		bottomRightNode: makeNode(grid, topLeftX+halfWidth, topLeftY+halfHeight, width-halfWidth, height-halfHeight),
	}
}

// isZoneHomogeneous est notre fonction qui vérifie si toutes les cases d'une zone ont le même type
func isZoneHomogeneous(grid [][]int, startX, startY, width, height int) (bool, int) {
	firstValue := grid[startY][startX]
	for y := startY; y < startY+height; y++ {
		for x := startX; x < startX+width; x++ {
			if grid[y][x] != firstValue {
				return false, -1
			}
		}
	}
	return true, firstValue
}


// writeFloorToFile écrit le contenu du tableau floorContent dans un fichier
func writeFloorToFile(fileName string, floorContent [][]int) error {
    file, err := os.Create(fileName)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    for _, row := range floorContent {
        for _, cell := range row {
            _, err := fmt.Fprintf(writer, "%d", cell)
            if err != nil {
                return err
            }
        }
        _, err = fmt.Fprintln(writer)
        if err != nil {
            return err
        }
    }
    return writer.Flush()
}