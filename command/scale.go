package command

import (
	"bufio"
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/codegangsta/cli"
	"github.com/disintegration/imaging"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func CmdScale(c *cli.Context) {
	scaleSize := c.Int("s")
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

			fp, err := os.Open(inFile)
			defer fp.Close()
			if err != nil {
				log.Fatal(err)
			}

			img, err := png.Decode(fp)
			if err != nil {
				log.Fatal(err)
			}

			w := img.Bounds().Size().X * scaleSize / 100

			nImg := imaging.Resize(img, w, 0, imaging.Lanczos)
			nfp, err := os.Create(outFile)
			defer nfp.Close()
			if err != nil {
				log.Fatal(err)
			}

			if err := png.Encode(nfp, nImg); err != nil {
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
