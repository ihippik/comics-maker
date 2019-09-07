package main

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type (
	ConfigApp struct {
		Config struct {
			Debug   bool
			Size    float64
			Spacing float64
			Blocks  []Block `yaml:"blocks"`
		}
	}

	Block struct {
		Size    float64
		Spacing float64
		X1      float64
		Y1      float64
		X2      float64
		Y2      float64
		Text    string
	}
)

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
	for i, block := range c.Config.Blocks {
		if block.Size == 0 {
			c.Config.Blocks[i].Size = c.Config.Size
		}
		if block.Spacing == 0 {
			c.Config.Blocks[i].Spacing = c.Config.Spacing
		}
	}
	return nil
}
