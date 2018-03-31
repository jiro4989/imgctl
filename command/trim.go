package command

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/codegangsta/cli"
	"github.com/oliamb/cutter"
)

func CmdTrim(c *cli.Context) {
	var (
		x   = c.Int("x")
		y   = c.Int("y")
		w   = c.Int("width")
		h   = c.Int("height")
		pt1 = image.Pt(x, y)
		pt2 = image.Pt(x+w, y+h)
	)
	outDir := c.String("d")
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	ch := make(chan int, runtime.NumCPU())
	var wg sync.WaitGroup

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		wg.Add(1)
		inFile := sc.Text()
		go func(inFile string) {
			defer wg.Done()
			ch <- 1
			base := filepath.Base(inFile)
			outFile := outDir + "/" + base

			reader, err := os.Open(inFile)
			defer reader.Close()
			if err != nil {
				log.Fatal(err)
			}

			img, err := png.Decode(reader)
			if err != nil {
				log.Fatal(err)
			}

			cImg, err := cutter.Crop(img, cutter.Config{
				Width:  w,
				Height: h,
				Anchor: image.Pt(x, y),
				Mode:   cutter.TopLeft,
			})
			if err != nil {
				log.Fatal(err)
			}

			oImg := image.NewRGBA(image.Rectangle{pt1, pt2})
			draw.Draw(oImg, oImg.Bounds(), cImg, pt1, draw.Over)

			writer, err := os.Create(outFile)
			defer writer.Close()
			if err != nil {
				log.Fatal(err)
			}

			if err := png.Encode(writer, oImg); err != nil {
				log.Fatal(err)
			}

			fmt.Println(outFile)
			<-ch
		}(inFile)
	}

	wg.Wait()

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
}
