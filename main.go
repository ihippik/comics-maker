package main

import (
	"bufio"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/urfave/cli"
)

var (
	dpi        = float64(72)
	ctx        = new(freetype.Context)
	utf8Font   = new(truetype.Font)
	black      = color.RGBA{0, 0, 0, 255}
	background *image.RGBA
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
			Value:       "wqy-zenhei.ttf",
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

	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		return err
	}

	utf8Font, err = freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}

	fontForeGroundColor := image.NewUniform(black)

	reader, err := os.Open(imgFile)
	if err != nil {
		logrus.Error(err)
		return err
	}

	tmpl, _, err := image.DecodeConfig(reader)
	if err != nil {
		logrus.Error(err)
		return err
	}
	_, err = reader.Seek(0, 0)
	if err != nil {
		logrus.Error(err)
		return err
	}

	background = image.NewRGBA(image.Rect(0, 0, tmpl.Width, tmpl.Height))

	templateImg, err := png.Decode(reader)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer reader.Close()

	draw.Draw(background, background.Bounds(), templateImg, image.ZP, draw.Src)

	ctx = freetype.NewContext()
	ctx.SetDPI(dpi)
	ctx.SetFont(utf8Font)
	ctx.SetClip(background.Bounds())
	ctx.SetDst(background)
	ctx.SetSrc(fontForeGroundColor)

	for i, block := range cfg.Config.Blocks {
		pt := freetype.Pt(block.X, block.Y+int(ctx.PointToFixed(block.Size)>>6))
		for _, str := range block.Strings {
			err = block.Validate(tmpl.Width, tmpl.Height)
			if err != nil {
				logrus.WithField("num_block", i).Warningln(err.Error())
				continue
			}

			ctx.SetFontSize(block.Size)
			_, err := ctx.DrawString(str, pt)
			if err != nil {
				logrus.Fatalln(err)
			}
			pt.Y += ctx.PointToFixed(block.Size * block.Spacing)
		}
	}

	outFile, err := os.Create(resultFile)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer outFile.Close()
	buff := bufio.NewWriter(outFile)

	err = png.Encode(buff, background)
	if err != nil {
		logrus.Fatalln(err)
	}

	err = buff.Flush()
	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.Infoln("save to", resultFile)
	return nil
}
