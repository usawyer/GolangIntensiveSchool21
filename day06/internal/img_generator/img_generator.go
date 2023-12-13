package img

import (
	"github.com/pkg/errors"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Image struct {
	Width  int
	Height int
	Img    *image.RGBA
}

func NewImage(width, height int) *Image {
	img := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: width, Y: height}})

	return &Image{
		Width:  width,
		Height: height,
		Img:    img,
	}
}

func (i Image) GenerateAmazingLogo() error {
	for x := 0; x < i.Width; x++ {
		for y := 0; y < i.Height; y++ {
			i.Img.Set(x, y, color.RGBA{R: 113, G: 156, B: 186, A: 255})
		}
	}

	i.drawSquare(140, 10, 40, color.RGBA{R: 103, G: 183, B: 175, A: 255})
	i.drawSquare(180, 50, 40, color.RGBA{R: 85, G: 215, B: 162, A: 255})
	i.drawSquare(220, 90, 60, color.RGBA{R: 67, G: 235, B: 153, A: 255})
	i.drawNumbers(15, 140, 30)
	i.addLabel(260, 18, "bfile")

	err := i.save("amazing_logo.png")
	if err != nil {
		return err
	}
	return nil
}

func (i Image) drawNumbers(width, height int, scale int) {
	pattern := [][]int{
		{1, 1, 1, 0, 0, 1, 0},
		{0, 0, 0, 1, 0, 0, 1},
		{0, 1, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 0, 0, 1},
	}

	for y := 0; y < len(pattern); y++ {
		for x := 0; x < len(pattern[y]); x++ {
			if pattern[y][x] == 1 {
				imgX := x*scale + width
				imgY := y*scale + height
				for k := 0; k < scale; k++ {
					for j := 0; j < scale; j++ {
						i.Img.Set(imgX+k, imgY+j, color.RGBA{G: 255, B: 174, A: 255})
					}
				}
			}
		}
	}
}

func (i Image) drawSquare(width, height int, size int, color color.RGBA) {
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			i.Img.Set(x+width, y+height, color)
		}
	}
}

func (i Image) addLabel(x, y int, label string) {
	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}

	d := &font.Drawer{
		Dst:  i.Img,
		Src:  image.NewUniform(color.Black),
		Face: basicfont.Face7x13,
		Dot:  point,
	}

	d.DrawString(label)
}

func (i Image) save(filePath string) error {
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()

	if err != nil {
		return errors.Wrap(err, "cannot create file")
	}

	err = png.Encode(imgFile, i.Img.SubImage(i.Img.Rect))
	if err != nil {
		return errors.Wrap(err, "cannot encode file")
	}
	return nil
}
