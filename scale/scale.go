package scale

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/disintegration/imaging"
	"github.com/jiro4989/imgctl/imageio"
	"github.com/jiro4989/lib-go/image"
)

func ScaleImages(outDir string, files []string, scaleSize int) ([]string, error) {
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		return nil, err
	}

	ch := make(chan int, runtime.NumCPU())
	var wg sync.WaitGroup

	var scaled []string
	for _, inFile := range files {
		wg.Add(1)
		go func(inFile string) {
			defer wg.Done()
			ch <- 1
			base := filepath.Base(inFile)
			outFile := outDir + "/" + base

			src, err := image.ReadFile(inFile)
			if err != nil {
				log.Fatal(err)
			}

			w := src.Bounds().Size().X * scaleSize / 100

			dist := imaging.Resize(src, w, 0, imaging.Lanczos)
			if err := imageio.WriteFile(outFile, dist); err != nil {
				log.Fatal(err)
			}

			scaled = append(scaled, outFile)
			<-ch
		}(inFile)
	}
	wg.Wait()

	return scaled, nil
}
