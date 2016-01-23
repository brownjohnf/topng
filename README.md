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
* Support resizing

## Usage

From the internet:

```
curl https://some.url.with/webp.webp | topng > out.png
```

Local file:

```
topng < in.gif > out.png
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

