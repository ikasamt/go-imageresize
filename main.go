//main.go

package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("no filename")
		os.Exit(1)
	}

	fileName := os.Args[1]
	baseFileName := filepath.Base(fileName)

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	img, formatName, err := image.Decode(file)
	if err != nil {
		log.Println(fileName)
		log.Println(formatName)
		log.Fatal(err)
	}

	// resize
	imgThumb := resize.Thumbnail(400, 400, img, resize.Lanczos3)

	// save
	outputFileName := fmt.Sprintf("r300x300_%s", baseFileName)
	out, err := os.Create(outputFileName)
	defer out.Close()
	if err != nil {
		log.Fatal(err)
	}
	jpeg.Encode(out, imgThumb, nil)
}
