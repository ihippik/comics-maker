package main

import (
	"image"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/fogleman/gg"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	var cfgFile, fontFile, imgFile, resultFile string

	app := cli.NewApp()
	app.Name = "сomics maker"
	app.Usage = "overlay text on images"
	app.Author = "hippik80@gmail.com"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "font, f",
			Value:       "font.ttf",
			Usage:       "font file name",
			Destination: &fontFile,
		},
		cli.StringFlag{
			Name:        "config, c",
			Value:       "blocks.yml",
			Usage:       "load configuration from `FILE`",
			Destination: &cfgFile,
		},
		cli.StringFlag{
			Name:        "image, i",
			Value:       "template.png",
			Usage:       "template image file",
			Destination: &imgFile,
		},
		cli.StringFlag{
			Name:        "result, r",
			Value:       "result.png",
			Usage:       "result file name",
			Destination: &resultFile,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "make",
			Aliases: []string{"m"},
			Usage:   "make image with text layer",
			Action: func(c *cli.Context) error {
				return makeImg(cfgFile, fontFile, imgFile, resultFile)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func makeImg(fileCfg, fontFile, imgFile, resultFile string) error {
	cfg, err := initConfig(fileCfg)
	if err != nil {
		logrus.Fatalln(err)
	}
	if err = cfg.Validate(); err != nil {
		logrus.Fatalln(err)
	}

	var img image.Image
	if isValidUrl(imgFile) {
		response, err := http.Get(imgFile)
		if err != nil {
			logrus.WithError(err).WithField("url", imgFile).Fatalln("can`t load img")
		}
		defer response.Body.Close()

		img, _, err = image.Decode(response.Body)
		if err != nil {
			logrus.WithError(err).Fatalln("can`t decode img")
		}
	} else {
		img, err = gg.LoadPNG(imgFile)
		if err != nil {
			logrus.WithError(err).Fatalln("can`t open img")
		}
	}
	dc := gg.NewContextForImage(img)

	for _, block := range cfg.Config.Blocks {
		dcImg := drawText(&block, fontFile)
		if cfg.Config.Debug {
			drawDebugRect(dcImg, 0, 0, block.X2-block.X1, block.Y2-block.Y1)
		}
		img := dcImg.Image()
		dc.DrawImage(img, int(block.X1), int(block.Y1))
	}
	err = dc.SavePNG(resultFile)
	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.Infoln("save to", resultFile)
	return nil
}

func drawText(b *Block, fontFile string) *gg.Context {
	dc := gg.NewContext(int(b.X2-b.X1), int(b.Y2-b.Y1))
	if err := dc.LoadFontFace(fontFile, b.Size); err != nil {
		logrus.WithError(err).Fatalln("can`t load font")
	}
	dc.SetRGB(0, 0, 0)
	textAlign := alignMap[b.TextAlign]
	dc.DrawStringWrapped(b.Text, 0, 0, 0, 0, b.X2-b.X1, b.Spacing, textAlign)
	return dc
}

func drawDebugRect(dc *gg.Context, x1, y1, x2, y2 float64) {
	dc.DrawLine(x1, y1, x2, y1)
	dc.DrawLine(x1, y1, x1, y2)
	dc.DrawLine(x1, y2, x2, y2)
	dc.DrawLine(x2, y1, x2, y2)
	dc.SetRGB(50, 255, 255)
	dc.SetLineWidth(1)
	dc.Stroke()
}

func isValidUrl(txt string) bool {
	_, err := url.ParseRequestURI(txt)
	if err != nil {
		return false
	} else {
		return true
	}
}
