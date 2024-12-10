package quadtree

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.

func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	numTileY := len(contentHolder)    // Nombre de lignes visibles
	numTileX := len(contentHolder[0]) // Nombre de colonnes visibles

	// On Parcour chaque case de contentHolder (tableau 2D utilisé pour représenter la zone visible du terrain sur l'écran)
	for y := 0; y < numTileY; y++ {
		for x := 0; x < numTileX; x++ {
			// On Calcul les coordonnées absolues dans le terrain
			absX := topLeftX + x
			absY := topLeftY + y

			// On Vérifie si les coordonnées sont hors limites
			if absX < 0 || absY < 0 || absX >= q.width || absY >= q.height {
				contentHolder[y][x] = -1 // ici la Zone est hors limites
			} else {
				// sinon on Obtien la valeur depuis le quadtree
				contentHolder[y][x] = q.root.getValueAt(absX, absY)
			}
		}
	}
}

// getValueAt récupère la valeur d'une case donnée dans le quadtree.
func (n *node) getValueAt(x, y int) int {
	// Si le nœud est une feuille on retourner sa valeur
	if n.isLeaf {
		return n.content
	}

	// on determine ici dans quel sous-nœud chercher
	midX := n.topLeftX + n.width/2
	midY := n.topLeftY + n.height/2

	if x < midX && y < midY { // Ici en haut a gauche
		return n.topLeftNode.getValueAt(x, y)
	} else if x >= midX && y < midY { // la en Haut a droite
		return n.topRightNode.getValueAt(x, y)
	} else if x < midX && y >= midY { // en bas a gauche
		return n.bottomLeftNode.getValueAt(x, y)
	} else { // en bas a droite
		return n.bottomRightNode.getValueAt(x, y)
	}
}

