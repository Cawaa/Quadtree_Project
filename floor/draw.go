package floor

import (
	"image"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw affiche dans une image (en général, celle qui représente l'écran),
// la partie du sol qui est visible (qui doit avoir été calculée avec Get avant).
func (f Floor) Draw(screen *ebiten.Image, camerax, cameray int) {
    for y := range f.content {
		for x := range f.content[y] {
			if configuration.Global.WaterAnimated {
				// Si l'animation de l'eau est activÃ©e, on affiche l'eau animÃ©e (4) et les autres blocs.
				if f.content[y][x] != -1 && f.content[y][x] != 4 {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x*configuration.Global.TileSize), float64(y*configuration.Global.TileSize))

					shiftX := f.content[y][x] * configuration.Global.TileSize

					screen.DrawImage(assets.FloorImage.SubImage(
						image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
					).(*ebiten.Image), op)
				}
			} else {
				// Sinon, on affiche uniquement les blocs.
				if f.content[y][x] != -1 {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x*configuration.Global.TileSize), float64(y*configuration.Global.TileSize))

					shiftX := f.content[y][x] * configuration.Global.TileSize

					screen.DrawImage(assets.FloorImage.SubImage(
						image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
					).(*ebiten.Image), op)
				}
			}
			if configuration.Global.WaterAnimated {
				// Si l'animation de l'eau est activÃ©e, on affiche l'eau animÃ©e (4) et les autres blocs.
				if f.content[y][x] == 4 {
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x*configuration.Global.TileSize), float64(y*configuration.Global.TileSize))
					shiftY := f.animation * configuration.Global.TileSize
					screen.DrawImage(assets.WaterImage.SubImage(
						image.Rect(0, shiftY, 0+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
					).(*ebiten.Image), op)
				}
			}
		}
	}
}
