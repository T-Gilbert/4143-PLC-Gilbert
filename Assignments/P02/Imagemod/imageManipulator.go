
package imageManipulator

import (
	"github.com/T-Gilbert/4143-PLC-Gilbert/tree/main/Assignments/P02"
)

type ImageManipulator struct {
	Image *P02.Context
}

func NewImageManipulator(width, height int) *ImageManipulator {
	img := P02.NewContext(width, height)
	return &ImageManipulator{Image: img}
}

func (im *ImageManipulator) SaveToFile(LazyBear string) error {
	return im.Image.SavePNG(LazyBear)
}

func (im *ImageManipulator) DrawRectangle(x, y, width, height float64) {
	im.Image.DrawRectangle(x, y, width, height)
	im.Image.Stroke()
}
