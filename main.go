package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

func openSVGStream(svgPath string) *oksvg.SvgIcon {
	fmt.Print("Read SVG stream")
	in, err := os.Open(svgPath)
	defer in.Close()
	if err != nil {
		panic(err)
	}
	icon, err := oksvg.ReadIconStream(in)
	if err != nil {
		panic(err)
	}
	return icon
}

func SaveToPngFile(filePath string, m image.Image) error {
	// Create the file
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	// Create Writer from file
	b := bufio.NewWriter(f)
	// Write the image into the buffer
	// err = png.Encode(b, m)
	err = jpeg.Encode(b, m, &jpeg.Options{Quality: 100})
	if err != nil {
		return err
	}
	err = b.Flush()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	svgPath := flag.String("svg", "", "svg file path")
	jpegPath := flag.String("jpeg", "", "jpeg file path")
	flag.Parse()
	if *svgPath == "" {
		panic("you should enter svg file path")
	}
	if *jpegPath == "" {
		panic("you should enter png file path")
	}

	svgIcon, err := oksvg.ReadIcon(*svgPath, oksvg.IgnoreErrorMode)
	if err != nil {
		panic(err)
	}
	w, h := int(svgIcon.ViewBox.W), int(svgIcon.ViewBox.H)

	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	fmt.Print("Draw SVG to PNG\n")
	scannerGV := rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())
	raster := rasterx.NewDasher(w, h, scannerGV)
	svgIcon.Draw(raster, 1.0)
	err = SaveToPngFile(*jpegPath, rgba)
	if err != nil {
		panic(err)
	}
	fmt.Print("Convert success")
}
