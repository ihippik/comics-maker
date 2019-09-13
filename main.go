package main

import (
	"os"

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
		logrus.Fatal(err)
	}
}
