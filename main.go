/*
Command twotone converts black/white PNGs to other twotone combinations.

Usage:

export TAN=EAE0CC
export BROWN=2D232A
twotone -bg=$TAN -fg=$BROWN -in=fixtures/test.png -out=fixtures/out.png
*/
package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
)

var (
	bgColor = flag.String("bg", "ffffff", "Background color.")
	fgColor = flag.String("fg", "000000", "Foreground color.")
	inFile  = flag.String("in", "", "Name of the file to use as input.")
	outFile = flag.String("out", "out.png", "Name of the file to use as output.")
)

func main() {
	flag.Parse()

	bg := toColor(*bgColor)
	fg := toColor(*fgColor)

	r, err := os.Open(*inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	m, _, err := image.Decode(r)
	if err != nil {
		log.Fatal(err)
	}
	in := m.(*image.RGBA)

	out := &image.RGBA{Pix: make([]uint8, len(in.Pix)), Stride: in.Stride, Rect: in.Rect}
	max := in.Rect.Max
	for x := 0; x < max.X; x++ {
		for y := 0; y < max.Y; y++ {
			v := in.At(x, y)
			if isBackground(v) {
				out.Set(x, y, bg)
				continue
			}
			out.Set(x, y, fg)
		}
	}

	w, _ := os.Create(*outFile)
	defer func() {
		if err := w.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	png.Encode(w, out)
}

func isBackground(c color.Color) bool {
	rgb := c.(color.RGBA)
	return rgb.R > 127 && rgb.G > 127 && rgb.B > 127
}

func toColor(s string) color.Color {
	r, _ := strconv.ParseInt(s[0:2], 16, 0)
	g, _ := strconv.ParseInt(s[2:4], 16, 0)
	b, _ := strconv.ParseInt(s[4:], 16, 0)

	return color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}
}
