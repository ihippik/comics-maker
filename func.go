package main

import (
	"image"
	"net/http"
	"net/url"

	"github.com/fogleman/gg"
	"github.com/sirupsen/logrus"
)

// makeImg main function that does all the work.
func makeImg(fileCfg, fontFile, imgFile, resultFile string) error {
	cfg, err := initConfig(fileCfg)
	if err != nil {
		logrus.Fatalln(err)
	}
	if err = cfg.Validate(); err != nil {
		logrus.Fatalln(err)
	}
	cfg.SetCommonValues()

	var img image.Image
	if isValidUrl(imgFile) {
		response, err := http.Get(imgFile)
		if err != nil {
			logrus.WithError(err).WithField("url", imgFile).Fatalln("can`t load img")
		}
		defer func() {
			if err = response.Body.Close(); err != nil {
				logrus.Fatalln(err)
			}
		}()

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

// drawText writes text on the template using the block settings.
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

// drawDebugRect draws a red rectangle around a text block.
func drawDebugRect(dc *gg.Context, x1, y1, x2, y2 float64) {
	dc.DrawLine(x1, y1, x2, y1)
	dc.DrawLine(x1, y1, x1, y2)
	dc.DrawLine(x1, y2, x2, y2)
	dc.DrawLine(x2, y1, x2, y2)
	dc.SetRGB(50, 255, 255)
	dc.SetLineWidth(1)
	dc.Stroke()
}

// isValidUrl check url.
func isValidUrl(txt string) bool {
	_, err := url.ParseRequestURI(txt)
	return err == nil
}
