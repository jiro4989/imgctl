package command

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/codegangsta/cli"
	"github.com/disintegration/imaging"
	"github.com/jiro4989/tkimgutil/imageutil"
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

			src, err := imageutil.ReadImage(inFile)
			if err != nil {
				log.Fatal(err)
			}

			w := src.Bounds().Size().X * scaleSize / 100

			dist := imaging.Resize(src, w, 0, imaging.Lanczos)
			if err := imageutil.WriteImage(outFile, dist); err != nil {
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
