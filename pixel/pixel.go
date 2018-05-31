package pixel

import (
	"fmt"
	"image"
	"image/color"
)

type Pixel struct {
	R, G, B, A uint32
}

func PrintRGBA(pixel * Pixel) {
	fmt.Printf("RED: %v \nGREEN: %v \nBLUE: %v \nALFA: %v \n", pixel.R, pixel.G, pixel.B, pixel.A)
}

func PrintGray(pixel * color.Gray) {
	fmt.Printf("Intesivity: %v \n", pixel.Y)
}

func PrintRGBA_arr(arr *[]Pixel){
	for _, pixel := range *arr{
		PrintRGBA(&pixel)
	}
}

func PrintGray_arr(arr *[]color.Gray){
	for _, pixel := range *arr{
		PrintGray(&pixel)
	}
}

func GetPixels(img image.Image) []Pixel{
	bounds := img.Bounds()
	pixelsAmount := bounds.Dx() * bounds.Dy()
	pixels := make([]Pixel, pixelsAmount)

	for i := 0; i < pixelsAmount; i++ {
		x := i % bounds.Dx()
		y := i / bounds.Dx()
		p := SetPixel(img.At(x, y))
		pixels[i] = p
	}

	return pixels
}

func GetColors(img image.Image) []color.Color{
	bounds := img.Bounds()
	pixelsAmount := bounds.Dx() * bounds.Dy()
	colors := make([]color.Color, pixelsAmount)

	for i := 0; i < pixelsAmount; i++ {
		x := i % bounds.Dx()
		y := i / bounds.Dx()
		p := img.At(x, y)
		colors[i] = p
	}

	return colors
}

func GetGrays(img image.Image) []color.Gray{
	bounds := img.Bounds()
	pixelsAmount := bounds.Dx() * bounds.Dy()
	grays := make([]color.Gray, pixelsAmount)

	for i := 0; i < pixelsAmount; i++ {
		x := i % bounds.Dx()
		y := i / bounds.Dx()
		p := RGBAtoGrayscale(img.At(x, y))
		grays[i] = p
	}

	return grays
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

func SetPixel(color color.Color) Pixel{
	p := Pixel{}
	p.R, p.G, p.B, p.A = color.RGBA();
	return p;
}
