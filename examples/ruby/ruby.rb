#!/usr/bin/env ruby

=begin
  ------
  DISCLAIMER: This as an EXAMPLE only, and does not reflect best practices.
  Your usage will likely vary, but this illustrates how this can be
  integrated into a ruby project without writing temp files, etc.
  ------

  To run these examples, ensure you have topng somewhere in your path, and some
  version of Ruby installed. This was tested only on Ruby 2.2.3.

  Then run:

  $ ./ruby.rb

  And you should see three (3) files in this directory:

  out_convert.png
  out_resize.png
  out_proportional.png

  They should look normal.
=end

module ToPNG
  class << self
    # Public: convert takes a source and destination and converts the source
    # file to a PNG and saves it at the destination.
    #
    # src   - String path to source file;
    # dest  - String path to destination file;
    #
    # Returns String destination.
    def convert(src, dest)
      input = File.read(src)

      open("|topng", "w+") do |f|
        f.print input

        File.open(dest, "w") do |out|
          out.print f.read
        end
      end

      dest
    end

    # Public: resize takes a source and destination and converts the source
    # file to a PNG and saves it at the destination after resizing it.
    #
    # src     - String path to source file;
    # dest    - String path to destination file;
    # height  - Height of the output image in pixels;
    # width   - Width of the output image in pixels;
    #
    # Returns String destination.
    def resize(src, dest, height: nil, width: nil)
      input = File.read(src)

      args = ""
      args += " -h #{height}" if height
      args += " -w #{width}" if width

      open("|topng#{args}", "w+") do |f|
        f.print input

        File.open(dest, "w") do |out|
          out.print f.read
        end
      end

      dest
    end
  end
end

# Just convert webp to png
ToPNG.convert("./in.webp", "./out_convert.png")

# Resize it to something weird
ToPNG.resize("./in.webp", "./out_resize.png", height: 100, width: 20)

# Resize it proportionally
ToPNG.resize("./in.webp", "./out_proportional.png", height: 100)

