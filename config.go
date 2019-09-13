package main

import (
	"errors"
	"github.com/fogleman/gg"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type (
	ConfigApp struct {
		Config struct {
			Debug     bool
			Size      float64
			Spacing   float64
			TextAlign string  `yaml:"textAlign"`
			Blocks    []Block `yaml:"blocks"`
		}
	}

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

func (c *ConfigApp) Validate() error {
	if c.Config.Spacing == 0 {
		return errors.New("spacing must be set")
	}
	if c.Config.Size == 0 {
		return errors.New("font size must be set")
	}
	if len(c.Config.TextAlign) == 0 {
		return errors.New("text align must be set")
	}
	_, ok := alignMap[c.Config.TextAlign]
	if !ok {
		return errors.New("invalid text align")
	}
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
	return nil
}
