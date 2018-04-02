package command

import (
	"fmt"
	"image"
	"image/draw"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
	"github.com/jiro4989/tkimgutil/imageutil"
)

type TOML struct {
	Image Image `toml:"image"`
}

type Image struct {
	SaveFilenameFormat string     `toml:"save_filename_format"`
	Pattern            [][]string `toml:"pattern"`
}

func CmdGenerate(c *cli.Context) {
	// 出力先ディレクトリの作成
	outDir := c.String("d")
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	var cfg TOML
	if _, err := toml.DecodeFile("./res/config.toml", &cfg); err != nil {
		log.Fatal(err)
	}

	// 画像の幅が必要なので先行して1枚だけload
	src, err := imageutil.ReadImage(cfg.Image.Pattern[0][0])
	if err != nil {
		log.Fatal(err)
	}
	width := src.Bounds().Size().X
	height := src.Bounds().Size().Y
	onFmt := outDir + "/" + cfg.Image.SaveFilenameFormat
	src = nil

	ch := make(chan int, runtime.NumCPU())
	var wg sync.WaitGroup

	for i, p := range cfg.Image.Pattern {
		wg.Add(1)
		dist := image.NewRGBA(image.Rect(0, 0, width, height))

		go func(i int, p []string) {
			defer wg.Done()
			ch <- 1

			for _, f := range p {
				src, err := imageutil.ReadImage(f)
				if err != nil {
					log.Fatal(err)
				}

				// 画像を貼り付け
				rect := image.Rectangle{
					image.ZP,
					image.Pt(width, height),
				}
				draw.Draw(dist, rect, src, image.ZP, draw.Over)
			}

			on := fmt.Sprintf(onFmt, (i + 1))
			if err := imageutil.WriteImage(on, dist); err != nil {
				log.Fatal(err)
			}
			fmt.Println(on)

			<-ch
		}(i, p)
	}

	wg.Wait()
}
