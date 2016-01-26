[![Build Status](https://travis-ci.org/brownjohnf/topng.svg?branch=master)](https://travis-ci.org/brownjohnf/topng)

# topng

GIF, JPEG, PNG, WebP -> PNG converter.

## Install

```
go install github.com/brownjohnf/goutils/topng
```

## TODO

* Support flags:
  * -h, help
  * -f, output file
  * -i, input file

## Usage

`topng` accepts a GIF, JPEG, PNG, or WebP image as a stream on stdin, and
streams a PNG on stdout. It will also output the input file type as it
was detected on stderr.

Resizing is supported via the -w and -h flags for height and width. If neither
is passed, the image will not be resized. If only one is passed, the image
will be resized proportionally.

```
Usage of ./topng:
  -h int
        resize-to height
  -w int
        resize-to width
```

### Example Usage

Converting images from the internet:

```
curl https://some.url.with/webp.webp | topng > out.png
```

Converting local files, resizing to 100px wide:

```
topng -w 100 < in.gif > out.png
```

### Examples

The following will grab a webp reference file from Google, and pipe it through
the converter:

```
curl https://www.gstatic.com/webp/gallery3/4_webp_ll.webp | topng > out.png
```

## Building

```
$ go get github.com/brownjohnf/goutils/topng
$ cd $GOPATH
$ go build # or go install
```

