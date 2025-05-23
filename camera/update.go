package camera

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

// Update met à jour la position de la caméra à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Camera) Update(characterPosX, characterPosY int) {

	switch configuration.Global.CameraMode {
	case Static:
		c.updateStatic()
	case FollowCharacter:
		c.updateFollowCharacter(characterPosX, characterPosY)
	}
}

// updateStatic est la mise-à-jour d'une caméra qui reste
// toujours à la position (0,0). Cette fonction ne fait donc
// rien.
func (c *Camera) updateStatic() {}

// updateFollowCharacter est la mise-à-jour d'une caméra qui
// suit toujours le personnage. Elle prend en paramètres deux
// entiers qui indiquent les coordonnées du personnage et place
// la caméra au même endroit.
func (c *Camera) updateFollowCharacter(characterPosX, characterPosY int) {
    if configuration.Global.CinematographicCamera {
        smoothingFactor := configuration.Global.CameraSmoothing // Facteur de lissage (0.0 < facteur < 1.0)
        c.X = int(float64(c.X)*(1-smoothingFactor) + float64(characterPosX)*smoothingFactor)
        c.Y = int(float64(c.Y)*(1-smoothingFactor) + float64(characterPosY)*smoothingFactor)
    } else {
        c.X = characterPosX
        c.Y = characterPosY
    }
}