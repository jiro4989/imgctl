package crop

import (
	"image"
	"image/draw"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/jiro4989/imgctl/imageio"
	"github.com/oliamb/cutter"
)

func CropImages(outDir string, files []string, x, y, w, h int) ([]string, error) {
	pt1 := image.Pt(x, y)
	pt2 := image.Pt(x+w, y+h)
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		return nil, err
	}

	ch := make(chan int, runtime.NumCPU())
	var wg sync.WaitGroup

	var cropped []string
	for _, inFile := range files {
		wg.Add(1)
		go func(inFile string) {
			defer wg.Done()
			ch <- 1
			base := filepath.Base(inFile)
			outFile := outDir + "/" + base

			src, err := imageio.ReadFile(inFile)
			if err != nil {
				log.Fatal(err)
			}

			cImg, err := cutter.Crop(src, cutter.Config{
				Width:  w,
				Height: h,
				Anchor: image.Pt(x, y),
				Mode:   cutter.TopLeft,
			})
			if err != nil {
				log.Fatal(err)
			}

			dist := image.NewRGBA(image.Rectangle{pt1, pt2})
			draw.Draw(dist, dist.Bounds(), cImg, pt1, draw.Over)

			if err := imageio.WriteFile(outFile, dist); err != nil {
				log.Fatal(err)
			}

			cropped = append(cropped, outFile)
			<-ch
		}(inFile)
	}

	wg.Wait()

	return cropped, nil
}
