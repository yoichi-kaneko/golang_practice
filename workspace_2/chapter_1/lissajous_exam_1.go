package main

import (
	"image"
	"image/color/palette"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var lineColor1 = color.RGBA{0, 200, 0, 255}
var lineColor2 = color.RGBA{0, 0, 200, 255}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const(
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette.WebSafe)

		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			var targetColor color.RGBA
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)

			if t < cycles * math.Pi {
				targetColor = lineColor1
			} else {
				targetColor = lineColor2
			}
			img.Set(size + int(x * size + 0.5), size + int(y * size + 0.5), targetColor)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
