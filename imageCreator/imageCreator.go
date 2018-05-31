package imageCreator

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"golang.org/x/image/bmp"
	"image/png"
	"log"
	"os"
	"strings"
)

type imageCreator struct {
	outPutPath, imageFormat string
}

func DrawImageRGB(width, height int, srcColors *[]color.Color) image.Image{
	imgRect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(imgRect)
	bounds := img.Bounds()

	for i, c := range *srcColors{
		x := i % bounds.Dx()
		y := i / bounds.Dx()

		img.Set(x, y, c)
	}

	fmt.Println("Drawing RGB completed...")
	return img
}

func DrawImageGray(width, height int, srcColors *[]color.Gray) image.Image{
	imgRect := image.Rect(0, 0, width, height)
	img := image.NewGray(imgRect)
	bounds := img.Bounds()

	for i, c := range *srcColors{
		x := i % bounds.Dx()
		y := i / bounds.Dx()

		img.Set(x, y, c)
	}

	fmt.Println("Drawing RGB completed...")
	return img
}

func SaveImage(img image.Image, quality int, fileName string){
	if quality > 100 || quality < 0{
		fmt.Printf("Quality %v to was out of range 0 - 100. Automaticly set to 80\n", quality)
		quality = 80
	}

	var opt jpeg.Options
	opt.Quality = quality
	out := createPlaceholderImage("./Images/output", fileName)
	ext := strings.Split(fileName, ".")[1]
	var err error
	switch ext {
	case "jpg":
		err = jpeg.Encode(out, img, &opt)
	case "bmp":
		err = bmp.Encode(out, img)
	case "png":
		err = png.Encode(out, img)
	}

	if err != nil {
		fmt.Printf("Failed to save image: %v\n", fileName)
		log.Fatal(err)
	}

	fmt.Printf("%v saved succesfully\n", fileName)
}

func createPlaceholderImage(path, name string) *os.File{
	out, err := os.Create(path + "/" + name)
	if err != nil {
		fmt.Println("Failed to create placeholder image")
		log.Fatal(err)
	}
	return out
}
