package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	height := len(floorContent)
	width := len(floorContent[0])

	// On Construie la racine de l'arbre
	root := makeNode(floorContent, 0, 0, width, height)

	// Retourner le quadtree
	q = Quadtree{
		width:  width,
		height: height,
		root:   root,
	}
	return
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
