package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"image/png"
)

type TextureAtlas struct {
	ImagePath   string       `xml:"imagePath,attr"`
	SubTextures []SubTexture `xml:"SubTexture"`
}

type SubTexture struct {
	Name   string `xml:"name,attr"`
	X      int    `xml:"x,attr"`
	Y      int    `xml:"y,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
}

func main() {

	// parse flags
	var filename string
	flag.StringVar(&filename, "filename", "", "path to xml")
	flag.Parse()

	if len(filename) == 0 {
		fmt.Print("missing filename, exit.\n")
		os.Exit(1)
	}

	// load data into struct
	fmt.Printf("loading file '%s'...\n", filename)
	xmlFile, err := os.Open(filename)
	defer xmlFile.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, _ := ioutil.ReadAll(xmlFile)
	var textureAtlas TextureAtlas
	xml.Unmarshal(b, &textureAtlas)
	fmt.Printf("found %d textures\n", len(textureAtlas.SubTextures))

	// load spritesheet and extract sprites
	basepath, _ := filepath.Abs(filepath.Dir(filename))
	spriteFilepath := filepath.Join(basepath, textureAtlas.ImagePath)
	spritesheet, err := os.Open(spriteFilepath)
	fmt.Printf("loading spritesheet %s\n", spriteFilepath)
	defer spritesheet.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spritesheetImage, _ := png.Decode(spritesheet)
	subImage := spritesheetImage.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(textureAtlas.SubTextures[0].X,
		textureAtlas.SubTextures[0].Y,
		textureAtlas.SubTextures[0].X+textureAtlas.SubTextures[0].Width,
		textureAtlas.SubTextures[0].Y+textureAtlas.SubTextures[0].Height))

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, subImage); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
