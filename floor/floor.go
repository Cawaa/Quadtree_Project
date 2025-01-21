package floor

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"

// Floor représente les données du terrain. Pour le moment
// aucun champs n'est exporté.
//
//   - content : partie du terrain qui doit être affichée à l'écran
//   - fullContent : totalité du terrain (utilisé seulement avec le type
//     d'affichage du terrain "fromFileFloor")
//   - quadTreeContent : totalité du terrain sous forme de quadtree (utilisé
//     avec le type d'affichage du terrain "quadtreeFloor")
type Floor struct {
	content         [][]int
	fullContent     [][]int
	quadtreeContent quadtree.Quadtree
	animation 		int
}

// types d'affichage du terrain disponibles
const (
	GridFloor int = iota
	FromFileFloor
	QuadTreeFloor
)

// GetHeight retourne la hauteur (en cases) du terrain
// à partir du tableau fullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetHeight() (height int) {
	return len(f.fullContent)
}

// GetWidth retourne la largeur (en cases) du terrain
// à partir du tableau fullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetWidth() (width int) {
	if len(f.fullContent) > 0 {
		width = len(f.fullContent[0])
	}
	return
}


// GetHeightQuad retourne la hauteur du terrain représenté par le quadtree.
// La hauteur est déterminée par le nombre de lignes dans le tableau représentant le quadtree.
func (f Floor) GetHeightQuad() (height int) {
	// Convertir le quadtree en tableau pour accéder à ses dimensions
	array := f.quadtreeContent.ToArray()

	// La hauteur est déterminée par le nombre de lignes dans le tableau
	return len(array)
}

// GetWidthQuad retourne la largeur du terrain représenté par le quadtree.
// La largeur est déterminée par le nombre de colonnes dans le tableau représentant le quadtree.
func (f Floor) GetWidthQuad() (width int) {
	// Convertir le quadtree en tableau pour accéder à ses dimensions
	array := f.quadtreeContent.ToArray()
	if len(array) > 0 {
		// La largeur est déterminée par le nombre de colonnes dans le tableau
		return len(array[0])
	}
	return 0
}