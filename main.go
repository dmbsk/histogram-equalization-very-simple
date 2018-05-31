package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"./pixel"
	"./imageOpener"
	"./imageCreator"
	"./HistogramEqualization"
)

func main() {
	path := "./images/source"
	images := imageOpener.GetImages(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for i, image := range images {
		equalizedImage := HistogramEqualization.Equalization(image)
		grays := pixel.GetGrays(equalizedImage)

		drawedImage := imageCreator.DrawImageGray(images[i].Bounds().Dx(), images[i].Bounds().Dy(), &grays)
		imageCreator.SaveImage(drawedImage, 80, files[i].Name())
		fmt.Printf("Image number %v completed\n\n", i + 1)

	}

	debug := "DONE"
	fmt.Println(debug)
	var _,_ = fmt.Println()
}
