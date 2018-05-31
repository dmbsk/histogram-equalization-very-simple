package HistogramEqualization

import (
	"image"
	"image/color"

	"../pixel"
)

type Grayscale struct {
	Intensivity color.Gray
	Amount uint32
	Frequency uint32
	Cdf float64
}

func RGBAtoGrayscale(clr color.Color) color.Gray{
	const rScale float32 = 0.299
	const gScale float32 = 0.587
	const bScale float32 = 0.114

	oldR, oldG, oldB, _ := clr.RGBA()
	gray := rScale * float32(oldR) + gScale * float32(oldG) + bScale * float32(oldB)
	pixel := color.Gray{uint8( gray / 256) }

	return pixel
}

func RGBAImageToGrayscale(colors [] color.Color) []color.Gray{
	grays := make([]color.Gray, len(colors))
	for i,color := range colors {
		grays[i] = RGBAtoGrayscale(color)
	}

	return grays
}

func Equalization(img image.Image) image.Image{
	bounds := img.Bounds()
	pixelsAmount := bounds.Dx() * bounds.Dy()

	colors := pixel.GetColors(img)
	grays := RGBAImageToGrayscale(colors)
	min, max := MinAndMax(grays)
	if min == 0 {
		min++
	}
	var scale float64
	scale = float64(255) / float64(max - min)

	newImg := image.NewGray(img.Bounds())
	for i := 0; i < pixelsAmount; i++ {
		x := i % bounds.Dx()
		y := i / bounds.Dx()

		newGray := color.Gray{}
		newGray.Y = uint8(scale * (float64(grays[i].Y) - float64(min)))
		newImg.SetGray(x, y, newGray)
	}

	return newImg
}

func MinAndMax(grays []color.Gray) (uint8, uint8) {
	var min uint8 = 255
	var max uint8 = 0
	for _, gray := range grays {
		if gray.Y < min {
			min = gray.Y
		}
		if gray.Y > max {
			max = gray.Y
		}
	}

	return min, max
}