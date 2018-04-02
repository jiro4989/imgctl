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
	"github.com/jiro4989/tkimgutil/command/config"
	"github.com/jiro4989/tkimgutil/imageutil"
)

func CmdGenerate(c *cli.Context) {
	cfgPath := c.String("config")

	// 設定ファイルの読み込み
	var cfg config.TOML
	if _, err := toml.DecodeFile(cfgPath, &cfg); err != nil {
		log.Fatal(err)
	}

	for _, cfgImg := range cfg.Images {
		outDir := cfgImg.OutDir + "/generate"
		gen := cfgImg.Generate

		if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		// 画像の幅が必要なので先行して1枚だけload
		src, err := imageutil.ReadImage(gen.Pattern[0][0])
		if err != nil {
			log.Fatal(err)
		}
		// loadした画像から横幅、縦幅、保存するファイル名書式を取得
		width := src.Bounds().Size().X
		height := src.Bounds().Size().Y
		onFmt := outDir + "/" + gen.SaveFilenameFormat
		src = nil // 使わないのでデータを破棄

		ch := make(chan int, runtime.NumCPU())
		var wg sync.WaitGroup

		for i, p := range gen.Pattern {
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
}
