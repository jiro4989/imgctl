package command

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/codegangsta/cli"
	"github.com/disintegration/gift"
	jimage "github.com/jiro4989/lib-go/image"
)

func CmdFlip(c *cli.Context) {
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

			src, err := jimage.ReadFile(inFile)
			if err != nil {
				log.Fatal(err)
			}

			g := gift.New(gift.FlipHorizontal())
			dist := image.NewRGBA(g.Bounds(src.Bounds()))
			g.Draw(dist, src)

			if err := jimage.WriteFile(outFile, dist); err != nil {
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
