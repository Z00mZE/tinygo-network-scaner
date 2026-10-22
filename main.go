package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/axellse/tdisplays3-tinygo/display"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

const (
	WIDTH  int16 = 320
	HEIGHT int16 = 170

	MS100 time.Duration = 113 * time.Millisecond
)

func main() {
	power := machine.GPIO15
	power.Configure(machine.PinConfig{Mode: machine.PinOutput})
	power.High()
	time.Sleep(100 * time.Millisecond)

	dspl, err := display.Init()
	if err != nil {
		panic(err)
	}

	for y := range HEIGHT {
		for x := range WIDTH {
			c := color.RGBA{R: 255, A: 255}
			switch {
			case y >= 57 && y < 114:
				c = getRainbowRGB(uint8((x * y) % 255))
			case y >= 114:
				c = color.RGBA{B: 255, A: 255}
			}

			dspl.SetPixel(x, y, c)
		}
	}

	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	tinyfont.WriteLine(dspl, &freesans.Bold18pt7b, 10, 35, "HELLO", white)
	tinyfont.WriteLine(dspl, &freesans.Bold18pt7b, 30, 35, "HELLO", white)

	if err := dspl.Display(); err != nil {
		panic(err)
	}

	for ts := range time.NewTicker(MS100).C {
		for y := range HEIGHT {
			for x := range WIDTH {
				switch {
				case y >= 57 && y < 114:
					dspl.SetPixel(x, y, getRainbowRGB(uint8(ts.UnixMilli()%255)))
				default:
					break
				}

			}
		}

		if err := dspl.Display(); err != nil {
			panic(err)
		}
	}
}

func getRainbowRGB(i uint8) color.RGBA {
	var out color.RGBA
	switch {
	case i < 85:
		out = color.RGBA{R: i * 3, G: 255 - i*3, A: 255}
	case i < 170:
		i -= 85
		out = color.RGBA{R: 255 - i*3, B: i * 3, A: 255}
	default:
		i -= 170
		out = color.RGBA{G: i * 3, B: 255 - i*3, A: 255}
	}
	println("color", i, "->", out.R, out.G, out.B)
	return out
}
