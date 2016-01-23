// The topng command reads an image from the standard input
// and writes it as a PNG image to the standard output.
package main

import (
	"fmt"
	"image"
	_ "image/gif"  // register GIF decoder
	_ "image/jpeg" // register JPEG decoder
	"image/png"
	"io"
	"os"

	_ "golang.org/x/image/webp"
)

func main() {
	if err := toPNG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "png: %v\n", err)
		os.Exit(1)
	}
}

func toPNG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return png.Encode(out, img)
}
