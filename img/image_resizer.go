package img

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
)

type IResizer interface {
	ResizeImage(data []byte) ([]byte, error)
}

type ImageResizer struct {
	width  uint
	height uint
}

func New(width, height uint) IResizer {
	return ImageResizer{
		width:  width,
		height: height,
	}
}

func (r ImageResizer) ResizeImage(data []byte) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	newImg := resize.Resize(r.width, r.height, img, resize.Lanczos3)

	var buf bytes.Buffer
	err = jpeg.Encode(&buf, newImg, nil)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
