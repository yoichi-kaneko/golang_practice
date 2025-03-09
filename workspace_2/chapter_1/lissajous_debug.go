package main

import (
	"image/color"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous()
}

func lissajous() {
	const(
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	fmt.Println(reflect.TypeOf(cycles))
}
