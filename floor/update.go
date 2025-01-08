package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos int) {
	f.animation++
	if f.animation == 32 {
		f.animation = 0
	}

	
	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY

	switch configuration.Global.FloorKind {
	case GridFloor:
		f.updateGridFloor(topLeftX, topLeftY)
	case FromFileFloor:
		f.updateFromFileFloor(topLeftX, topLeftY)
	case QuadTreeFloor:
		f.updateQuadtreeFloor(topLeftX, topLeftY)
	}
}

// le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(topLeftX, topLeftY int) {
	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			absX := topLeftX
			if absX < 0 {
				absX = -absX
			}
			absY := topLeftY
			if absY < 0 {
				absY = -absY
			}
			f.content[y][x] = ((x + absX%2) + (y + absY%2)) % 2
		}
	}
}


// updateFromFileFloor met à jour le contenu du sol en fonction des coordonnées
// fournies (topLeftX, topLeftY) en utilisant les données de fullContent.
//
// Paramètres:
// - topLeftX: Coordonnée X du coin supérieur gauche
// - topLeftY: Coordonnée Y du coin supérieur gauche
//
// La fonction parcourt chaque cellule de content et calcule les coordonnées
// absolues correspondantes dans fullContent. Si les coordonnées absolues sont
// valides (c'est-à-dire qu'elles se trouvent dans les limites de fullContent),
// la cellule de content est mise à jour avec la valeur correspondante de fullContent.
// Si les coordonnées sont hors des limites, la cellule de content est remplie
// avec une valeur par défaut (-1).
func (f *Floor) updateFromFileFloor(topLeftX, topLeftY int) {
    for y := 0; y < len(f.content); y++ {
        for x := 0; x < len(f.content[y]); x++ {
            // Calcul des coordonnées absolues dans fullContent
            absX := topLeftX + x
            absY := topLeftY + y

            // Vérifier si les coordonnées absolues sont valides dans fullContent
            if absY >= 0 && absY < len(f.fullContent) && absX >= 0 && absX < len(f.fullContent[absY]) {
                // Mettre à jour la cellule de content avec la valeur de fullContent
                f.content[y][x] = f.fullContent[absY][absX]
            } else {
                // Si hors des limites, remplir avec une valeur par défaut (-1)
                f.content[y][x] = -1
            }
        }
    }
}
// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(topLeftX, topLeftY int) {
	f.quadtreeContent.GetContent(topLeftX, topLeftY, f.content)
}
