package imageutil

import (
	"github.com/nfnt/resize"
	"image"
)

func ResizeThumbnail(image image.Image) image.Image {
	img := resize.Thumbnail(300, 300, image, resize.NearestNeighbor)
	return img
}
