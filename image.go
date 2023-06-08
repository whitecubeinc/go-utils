package utils

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"net/http"
)

type ImageBuilder struct {
	io.Reader
	image.Image
	FormatName string
}

func NewImageBuilder(r io.Reader) *ImageBuilder {
	img, formatName, err := image.Decode(r)
	if err != nil {
		panic(err)
	}

	return &ImageBuilder{
		Reader:     r,
		Image:      img,
		FormatName: formatName,
	}
}

// Compress return as JPEG format
func (ib *ImageBuilder) Compress() *bytes.Buffer {
	newImg := image.NewRGBA(ib.Image.Bounds())

	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), ib.Image, ib.Image.Bounds().Min, draw.Over)

	opt := jpeg.Options{
		Quality: 70,
	}
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, newImg, &opt)
	if err != nil {
		panic(err)
	}

	return &buf
}

func (ib *ImageBuilder) DrawOnWhiteBackground() *image.RGBA {
	newImg := image.NewRGBA(ib.Image.Bounds())

	// can change it to whichever color you want with
	// a new color.RGBA{} and use image.NewUniform(color.RGBA{<fill in color>}) function
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), ib.Image, ib.Image.Bounds().Min, draw.Over)

	return newImg
}

func (ib *ImageBuilder) Resize() {
	maxWidth := 900

	if ib.Bounds().Max.X > maxWidth {
		ib.Image = resize.Resize(uint(maxWidth), 0, ib.Image, resize.NearestNeighbor)
	}
}

func (ib *ImageBuilder) CutIntoSquare() (images []*bytes.Buffer) {
	ib.Resize()
	newImg := ib.DrawOnWhiteBackground()

	bound := newImg.Bounds()
	option := jpeg.Options{
		Quality: 70,
	}

	y := 0
	for y < bound.Max.Y {
		nextY := y + bound.Max.X
		if nextY > bound.Max.Y {
			nextY = bound.Max.Y
		}

		subImage := newImg.SubImage(image.Rect(0, y, bound.Max.X, nextY))

		var buf bytes.Buffer
		err := jpeg.Encode(&buf, subImage, &option)
		if err != nil {
			panic(err)
		}

		images = append(images, &buf)

		y = nextY
	}

	return images
}

func DownloadImage(imageURL string) *bytes.Buffer {
	response, err := http.Get(imageURL)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		panic(err)
	}

	return &buf
}
