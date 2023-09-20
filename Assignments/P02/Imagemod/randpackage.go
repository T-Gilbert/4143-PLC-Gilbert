package imagemod

import (
	"github.com/fogleman/gg"
)

// ImageManipulator represents an image manipulation tool.
type ImageManipulator struct {
	Image     *gg.Context
	ImagePath string // Add a field to store the image path
}
