package flip

import (
	"image"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/disintegration/gift"
	"github.com/jiro4989/imgctl/imageio"
)

func FlipImages(outDir string, files []string) ([]string, error) {
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		return nil, err
	}

	ch := make(chan int, runtime.NumCPU())
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	var flipped []string
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

			g := gift.New(gift.FlipHorizontal())
			dist := image.NewRGBA(g.Bounds(src.Bounds()))
			g.Draw(dist, src)

			if err := imageio.WriteFile(outFile, dist); err != nil {
				log.Fatal(err)
			}

			mutex.Lock()
			flipped = append(flipped, outFile)
			mutex.Unlock()
			<-ch
		}(inFile)
	}

	wg.Wait()

	return flipped, nil
}
