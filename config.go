package main

import (
	"errors"
	"io/ioutil"

	"github.com/fogleman/gg"
	"gopkg.in/yaml.v2"
)

type (
	//ConfigApp struct with app configuration.
	ConfigApp struct {
		Config struct {
			Debug     bool
			Size      float64
			Spacing   float64
			TextAlign string `yaml:"textAlign"`
			Blocks    []Block
		}
	}
	//Block struct with text block settings.
	Block struct {
		Size      float64
		Spacing   float64
		TextAlign string `yaml:"textAlign"`
		X1        float64
		Y1        float64
		X2        float64
		Y2        float64
		Text      string
	}
)

const (
	alignLeft   = "left"
	alignCenter = "center"
	alignRight  = "right"
)

// text errors.
var (
	spacingNotSetErr           = errors.New("spacing must be set")
	fontSizeNotSetErr          = errors.New("font size must be set")
	textAlignNotSetErr         = errors.New("text align must be set")
	invalidTextAlignErr        = errors.New("invalid text align")
	invalidHorizontalCoordsErr = errors.New("invalid horizontal block coordinates")
	invalidVerticalCoordsErr   = errors.New("invalid vertical block coordinates")
)

var alignMap = map[string]gg.Align{alignLeft: gg.AlignLeft, alignCenter: gg.AlignCenter, alignRight: gg.AlignRight}

func initConfig(file string) (ConfigApp, error) {
	var AppConfig ConfigApp
	c, err := ioutil.ReadFile(file)
	if err != nil {
		return AppConfig, err
	}
	if err := yaml.Unmarshal(c, &AppConfig); err != nil {
		return AppConfig, err
	}
	return AppConfig, nil
}

// Validate config file.
func (c *ConfigApp) Validate() error {
	if c.Config.Spacing == 0 {
		return spacingNotSetErr
	}
	if c.Config.Size == 0 {
		return fontSizeNotSetErr
	}
	if len(c.Config.TextAlign) == 0 {
		return textAlignNotSetErr
	}
	_, ok := alignMap[c.Config.TextAlign]
	if !ok {
		return invalidTextAlignErr
	}
	for _, block := range c.Config.Blocks {
		if block.X1 >= block.X2 {
			return invalidHorizontalCoordsErr
		}
		if block.Y1 >= block.Y2 {
			return invalidVerticalCoordsErr
		}
		_, ok := alignMap[block.TextAlign]
		if !ok && len(block.TextAlign) > 0 {
			return invalidTextAlignErr
		}
	}
	return nil
}

// SetCommonValues set common settings for each block.
func (c *ConfigApp) SetCommonValues() {
	for i, block := range c.Config.Blocks {
		if block.Size == 0 {
			c.Config.Blocks[i].Size = c.Config.Size
		}
		if block.Spacing == 0 {
			c.Config.Blocks[i].Spacing = c.Config.Spacing
		}
		if len(block.TextAlign) == 0 {
			c.Config.Blocks[i].TextAlign = c.Config.TextAlign
		}
	}
}
