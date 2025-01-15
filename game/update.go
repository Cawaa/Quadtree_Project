package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update met à jour les données du jeu à chaque 1/60 de seconde.
// Il faut bien faire attention à l'ordre des mises-à-jour car elles
// dépendent les unes des autres (par exemple, pour le moment, la
// mise-à-jour de la caméra dépend de celle du personnage et la définition
// du terrain dépend de celle de la caméra).
func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		configuration.Global.DebugMode = !configuration.Global.DebugMode
	}

	g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y))
	g.camera.Update(g.character.X, g.character.Y)
	g.floor.Update(g.camera.X, g.camera.Y)
	
	
	if configuration.Global.Zoom{
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			configuration.Global.ScreenHeight -= configuration.Global.TileSize * configuration.Global.NumTileX / 2
			configuration.Global.ScreenWidth -= configuration.Global.TileSize * configuration.Global.NumTileY / 2
			configuration.Global.ScreenCenterTileX -= configuration.Global.TileSize / 6
			configuration.Global.ScreenCenterTileY -= configuration.Global.TileSize / 6
		}
	


		if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
			configuration.Global.ScreenHeight += configuration.Global.TileSize * configuration.Global.NumTileX / 2
			configuration.Global.ScreenWidth += configuration.Global.TileSize * configuration.Global.NumTileY / 2
			configuration.Global.ScreenCenterTileX += configuration.Global.TileSize / 6
			configuration.Global.ScreenCenterTileY += configuration.Global.TileSize / 6
		}
	}


	return nil
}
