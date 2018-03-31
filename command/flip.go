package command

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/codegangsta/cli"
	"github.com/disintegration/gift"
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

			reader, err := os.Open(inFile)
			defer reader.Close()
			if err != nil {
				log.Fatal(err)
			}

			src, err := png.Decode(reader)
			if err != nil {
				log.Fatal(err)
			}

			g := gift.New(gift.FlipHorizontal())
			dist := image.NewRGBA(g.Bounds(src.Bounds()))
			g.Draw(dist, src)

			writer, err := os.Create(outFile)
			defer writer.Close()
			if err != nil {
				log.Fatal(err)
			}

			if err := png.Encode(writer, dist); err != nil {
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
