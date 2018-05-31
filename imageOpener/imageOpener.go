package imageOpener

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
)

func GetImages(dir string) []image.Image {
	var images []image.Image

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		img := LoadImage(path)
		images = append(images, img)
		return nil
	})
	if err != nil{
		log.Fatal(err)
	}
	return images
}

func LoadImage(path string) image.Image {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	fileNameSplited := strings.Split(path, "/")
	name := fileNameSplited[len(fileNameSplited) - 1]
	img, err := decodeAll(f, name)
	if err != nil {
		fmt.Println("Failed to decode image")
		log.Fatal(err)
	}

	return img
}

func decodeAll(f io.Reader, name string) (image.Image, error) {
	ext := strings.Split(name, ".")
	switch ext[len(ext) - 1] {
	case "jpg":
		img, err := jpeg.Decode(f)
		return img, err
	case "bmp":
		img, err := bmp.Decode(f)
		return img, err
	case "png":
		img, err := png.Decode(f)
		return img, err
	}

	img, err := jpeg.Decode(f)
	return img, err
}