package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"os"

	_ "image/jpeg"
	_ "image/png"

	"github.com/donatj/ldraw"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s <img>:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	file, err := os.Open(flag.Arg(0))
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	sb := img.Bounds()

	fmt.Print("0 Untitled Model\r\n")
	fmt.Print("0 Name:  Foo\r\n")
	fmt.Print("0 Author:  \r\n")

	for y := sb.Min.Y; y < sb.Max.Y; y += 10 {
		for x := sb.Min.X; x < sb.Max.X; x += 10 {
			c := img.At(x, y)
			bc := ldraw.NearestBrickColor(c)

			var h float64 = -24

			r, g, b, _ := c.RGBA()
			h += (float64(r+g+b) / (0xffff * 3)) * 30

			fmt.Printf("1 %d %f %f %f 1.000000 0.000000 0.000000 0.000000 1.000000 0.000000 0.000000 0.000000 1.000000 3005.dat\r\n", bc, float64(x*2), h, float64(y*2))
		}
	}
}
