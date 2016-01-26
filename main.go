// The topng command reads an image from the standard input
// and writes it as a PNG image to the standard output.
package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"  // register GIF decoder
	_ "image/jpeg" // register JPEG decoder
	"image/png"
	"io"
	"os"

	"github.com/nfnt/resize"
	_ "golang.org/x/image/webp"
)

var w = flag.Int("w", 0, "resize-to width")
var h = flag.Int("h", 0, "resize-to height")

func main() {
	flag.Parse()

	if err := toPNG(os.Stdin, os.Stdout, uint(*w), uint(*h)); err != nil {
		fmt.Fprintf(os.Stderr, "topng: %v\n", err)
		os.Exit(1)
	}
}

func toPNG(in io.Reader, out io.Writer, width uint, height uint) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format: ", kind)

	if width > 0 || height > 0 {
		fmt.Fprintln(os.Stderr, "resizing image")
		newImg := resize.Resize(width, height, img, resize.Lanczos3)

		return png.Encode(out, newImg)
	}

	return png.Encode(out, img)
}
