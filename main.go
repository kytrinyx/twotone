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
	dropTransparency = flag.Bool("drop-transparency", false, "Replace transparent background with background color.")
	threshold        = flag.Int("threshold", 127, "Default background threshold.")
	bgColor          = flag.String("bg", "", "Background color. Default is transparent.")
	fgColor          = flag.String("fg", "000000", "Foreground color.")
	inFile           = flag.String("in", "", "Name of the file to use as input.")
	outFile          = flag.String("out", "out.png", "Name of the file to use as output.")
)

func main() {
	flag.Parse()

	bg := color.Color(color.RGBA{0, 0, 0, 0}) // default background
	if *bgColor != "" {
		bg = toColor(*bgColor)
	}

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

	var (
		stride, pix int
	)

	switch in := m.(type) {
	case *image.RGBA:
		stride = in.Stride
		pix = len(in.Pix)
	case *image.NRGBA:
		stride = in.Stride
		pix = len(in.Pix)
	}
	out := &image.RGBA{Pix: make([]uint8, pix), Stride: stride, Rect: m.Bounds()}

	max := m.Bounds().Max
	for x := 0; x < max.X; x++ {
		for y := 0; y < max.Y; y++ {
			v := m.At(x, y)
			if !*dropTransparency && isTransparent(v) {
				continue
			}
			if isTransparent(v) || isBackground(v) {
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

const n = 256 // scaling factor that I don't understand
func isBackground(c color.Color) bool {
	r, g, b, _ := c.RGBA()
	v := uint32((*threshold) * n)
	return r > v && g > v && b > v
}

func isTransparent(c color.Color) bool {
	r, g, b, a := c.RGBA()
	return r == 0 && g == 0 && b == 0 && a == 0
}

func toColor(s string) color.Color {
	r, _ := strconv.ParseInt(s[0:2], 16, 0)
	g, _ := strconv.ParseInt(s[2:4], 16, 0)
	b, _ := strconv.ParseInt(s[4:], 16, 0)

	return color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}
}
