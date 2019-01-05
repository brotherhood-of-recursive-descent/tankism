package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path"
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
	var outputDir string
	flag.StringVar(&filename, "filename", "", "path to xml")
	flag.StringVar(&outputDir, "outputDir", "output", "path to output directory")
	flag.Parse()

	if len(filename) == 0 {
		printUsage()
		os.Exit(1)
	}

	// load data into struct
	fmt.Println("loading textureAtlas: ", filename)
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
	defer spritesheet.Close()

	spritesheetImage, _ := png.Decode(spritesheet)

	if err != nil {
		fmt.Println("failed loading image referece: ", err)
		os.Exit(1)
	}
	fmt.Printf("loading image reference %s\n", spriteFilepath)

	// create output directory
	err = createDirectory(outputDir)
	if err != nil {
		fmt.Println("failed creating outputDirectory: ", err)
		os.Exit(1)
	}

	// extract individual textures
	for _, entry := range textureAtlas.SubTextures {
		subImage, err := getSubImage(spritesheetImage, entry)
		if err != nil {
			fmt.Println("skipping image due to: ", err)
			continue
		}

		if err := writePngImage(subImage, path.Join(outputDir, entry.Name)); err != nil {
			fmt.Println("skipping image due to: ", err)
			continue
		}
		fmt.Println("wrote image: ", entry.Name)
	}
}

// getSubImage extracts a subtexture from a given spritesheet
func getSubImage(spritesheet image.Image, texture SubTexture) (image.Image, error) {
	subImage := spritesheet.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(texture.X,
		texture.Y,
		texture.X+texture.Width,
		texture.Y+texture.Height))

	return subImage, nil
}

// writePngImage writes the given image to a png file
func writePngImage(subImage image.Image, name string) error {
	f, err := os.Create(name)
	defer f.Close()

	if err != nil {
		return err
	}

	if err := png.Encode(f, subImage); err != nil {
		return err
	}

	return nil
}

// createDirectory creates the directory structure if not present
func createDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}
