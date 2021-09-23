package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var palette = []color.Color{color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}}

func main() {
	os.Args = []string{"lissajous", "web"}
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			qcycles, err := strconv.Atoi(strings.Join(r.Form["cycles"], ""))
			if err != nil {
				qcycles = 20.0
			}
			lissajous(w, float64(qcycles))
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	outfile, err := os.Create("out.gif")
	if err != nil {
		log.Fatal(err)
	}
	lissajous(outfile, 20.0)
}
func lissajous(out io.Writer, qcycles float64) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	cycles := qcycles
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(t)%4)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
