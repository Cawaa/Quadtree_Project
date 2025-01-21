package configuration

import (
	"encoding/json"
	"log"
	"os"
	"fmt"
	"bufio"
)

// Configuration définit les élèments de la configuration
// du jeu. Pour ajouter un élèment de configuration il
// suffit d'ajouter un champs dans cette structure.
//
// Les champs directement lus dans le fichier de configuration sont :
//   - DebugMode : indique si on est en mode debug ou pas
//   - NumTileX, NumTileY : les nombres de cases affichées à l'écran
//     en largeur et hauteur.
//   - TileSize : la taille en pixels du côté d'une case.
//   - NumCharacterAnimImages : le nombre de d'images de l'animation du
//     personnage.
//   - NumFramePerCharacterAnimImage : le nombre d'appels à update ou
//     de 1/60 de seconde) qui ont lieu entre deux images de l'animation
//     du personnage.
//   - NumTileForDebug : le nombre de cases à ajouter à droite de l'écran
//     pour afficher les informations de debug
//   - CameraMode : le type de caméra à utiliser (0 pour une caméra fixe
//     et 1 pour une caméra qui suit le personnage).
//   - FloorKind : détermine la méthode à utiliser pour afficher le terrain
//     (quadrillage, lecture dans un fichier, quadtree, etc)
//   - FloorFile : le chemin d'un fichier où lire les informations sur le
//     terrain si nécessaire
//	 - Waterblock : si true impossible d'aller sur les blocs d'eau
// 	 - RandomFloor : si true, le terrain sera généré aléatoirement
//   - WaterAnimated : si true, l'eau sera animée
//   - Zoom : si true, la fonctionnalité de zoom sera activée
//   - Teleportation : si true, la fonctionnalité de téléportation sera activée
//   - CinematographicCamera : si true, la caméra sera utilisée pour des effets cinématographiques
//   - CameraSmoothing : la valeur de lissage de la caméra
//   - InfiniteGenExtension : si true, le terrain sera généré de manière infinie
//   - RandomGenExtension : si true, le terrain sera généré aléatoirement
//   - RandomWidht, RandomHeight : la largeur et la hauteur du terrain généré aléatoirement
//   - FloorWidth, FloorHeigth : la largeur et la hauteur du terrain
//
// Les champs calculés à partir des précédents sont :
//   - ScreenWidth, ScreenHeight : la largeur et la hauteur de l'écran
//     en pixels (hors zone d'affichage pour le debug)
//   - ScreenCenterTileX, ScreenCenterTileY : les coordonnées de la case
//     au centre de l'écran, où sera placé la caméra.

type Configuration struct {
	DebugMode                     bool
	NumTileX, NumTileY            int
	TileSize                      int
	NumCharacterAnimImages        int
	NumFramePerCharacterAnimImage int
	NumTileForDebug               int
	CameraMode                    int
	FloorKind                     int
	FloorFile                     string
	Waterblock 					  bool
	RandomFloor 				  bool
	WaterAnimated 				  bool
	Zoom						  bool
	Teleportation 				  bool
	CinematographicCamera         bool
	CameraSmoothing               float64
	InfiniteGenExtension 	      bool 
	RandomGenExtension bool
	RandomWidht        int
	RandomHeight       int
	FloorWidth                       int
	FloorHeigth                      int




	ScreenWidth, ScreenHeight            int `json:"-"`
	ScreenCenterTileX, ScreenCenterTileY int `json:"-"`
}

// Global est la variable qui contient la configuration
// du jeu. Sa valeur est fixée à partir de la lecture d'un
// fichier de configuration par la fonction Load. C'est
// cette variable qu'il faut lire (configuration.Global)
// pour accéder à la configuration depuis d'autres paquets.
var Global Configuration

// Load se charge de lire un fichier de configuration, de
// remplir les champs obtenus par simple lecture, puis
// d'appeler la fonction qui remplit les champs calculés.
func Load(configurationFileName string) {
	content, err := os.ReadFile(configurationFileName)
	if err != nil {
		log.Fatal("Error while opening configuration file: ", err)
	}

	err = json.Unmarshal(content, &Global)
	if err != nil {
		log.Fatal("Error while reading configuration file: ", err)
	}

	setComputedFields()
}

// setComputedFields se charge de remplir les champs calculés
// de la configuration à partir des autres champs.
func setComputedFields() {
	if Global.RandomGenExtension {
		Global.FloorWidth = Global.RandomWidht
		Global.FloorHeigth = Global.RandomHeight
	} else {
		widhtFile, heightFile := getWidhtHeightOfFile(Global.FloorFile)
		Global.FloorWidth = widhtFile
		Global.FloorHeigth = heightFile
	}
	Global.ScreenWidth = Global.NumTileX * Global.TileSize
	Global.ScreenHeight = Global.NumTileY * Global.TileSize
	Global.ScreenCenterTileX = Global.NumTileX / 2
	Global.ScreenCenterTileY = Global.NumTileY / 2
}

// getWidhtHeightOfFile lit un fichier et retourne la largeur et la hauteur du contenu du fichier.
// La largeur est déterminée par le nombre de caractères dans la première ligne du fichier,
// et la hauteur est déterminée par le nombre de lignes dans le fichier.
func getWidhtHeightOfFile(fileName string) (width, height int) {
	// Ouvrir le fichier en lecture
	file, err := os.Open(fileName)
	if err != nil {
		// Si une erreur se produit lors de l'ouverture du fichier, afficher l'erreur et retourner
		fmt.Println(err)
		return
	}
	defer file.Close() // Fermer le fichier lorsque la fonction se termine

	// Créer un scanner pour lire le fichier ligne par ligne
	scanner := bufio.NewScanner(file)

	// Initialiser les variables pour stocker la largeur et la hauteur
	var ligneCount int
	var colonneCount int

	// Lire le fichier ligne par ligne
	for scanner.Scan() {
		// Incrémentation du compteur de lignes
		ligneCount++

		// Lire la première ligne pour déterminer la largeur
		ligne := scanner.Text()
		if ligneCount == 1 {
			// La largeur est déterminée par le nombre de caractères dans la première ligne
			colonneCount = len(ligne)
		}
	}

	// Retourner la largeur et la hauteur
	return colonneCount, ligneCount
}

