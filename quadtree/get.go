package quadtree

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	numTileY := len(contentHolder)
	numTileX := len(contentHolder[0])

	for y:=0; y<numTileY; y++ {
		for x:=0; x<numTileX;x++ {
			absX := topLeftX + x
			absY := topLeftY + y 

			if absX < 0 || absY < 0 || absX >= q.width || absY >= q.height {
				contentHolder[y][x] = -1
			} else {
				contentHolder[x][y] = q.root.getValueAt(absX,absY)
			}
		}
	}
}

func (n *node) getValueAt(x, y int) int {
	if n.isLeaf {
		return n.content
	}
	midX := n.topLeftX + n.width/2
	midY := n.topLeftY + n.height/2

	if x < midX && y < midY {
		return n.topLeftNode.getValueAt(x,y)
	} else if x >= midX && y >= midY {
		return n.topRightNode.getValueAt(x,y)
	} else if x < midX && y >= midY {
		return n.bottomLeftNode.getValueAt(x,y)
	} else {
		return n.bottomRightNode.getValueAt(x,y)
	}
}
