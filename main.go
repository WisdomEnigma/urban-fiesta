package main

import (
	"image/png"
	"log"
	"os"
	"strings"

	prints "github.com/WisdomEnigma/urban-fiesta/fingerprint"
)

func main() {
	file, err := os.Open("ms-icon-310x310.png")
	if err != nil {
		log.Fatal("Error opening file:", err)
		return
	}

	defer file.Close()

	fstat, err := os.Stat(file.Name())
	if err != nil {
		log.Fatal("Error stat:", err)
		return
	}

	log.Println("Icon:", fstat)

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal("Error decoding image:", err)
		return
	}

	_imagePrint := &prints.Image_Print{}
	repeat := make([]string, img.Bounds().Max.X)

	for i := 0; i < img.Bounds().Max.X; i++ {
		r, g, b, k := img.At(i, i).RGBA()
		repeat[i] = _imagePrint.CalculateHashColor(r, g, b, k, int64(i))
	}

	_hasher := strings.Join(repeat, "")
	log.Println("value:", string(_hasher[:]))
}
