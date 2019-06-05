package flip

import (
	"bufio"
	"image"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/disintegration/gift"
	"github.com/jiro4989/imgctl/imageio"
)

func FlipImages(outDir, saveFileNameFormat string, files []string) ([]string, error) {
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	ch := make(chan int, runtime.NumCPU())
	var wg sync.WaitGroup

	var flipped []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		wg.Add(1)
		inFile := sc.Text()
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

			flipped = append(flipped, outFile)
			<-ch
		}(inFile)
	}

	wg.Wait()

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return flipped, nil
}
