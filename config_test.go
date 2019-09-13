package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigApp_Validate(t *testing.T) {
	type fields struct {
		Config struct {
			Debug     bool
			Size      float64
			Spacing   float64
			TextAlign string `yaml:"textAlign"`
			Blocks    []Block
		}
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "success",
			fields: fields{
				Config: struct {
					Debug     bool
					Size      float64
					Spacing   float64
					TextAlign string `yaml:"textAlign"`
					Blocks    []Block
				}{
					Debug:     false,
					Size:      10,
					Spacing:   10,
					TextAlign: "left",
					Blocks: []Block{
						{
							Size:      10,
							Spacing:   20,
							TextAlign: "left",
							X1:        10,
							Y1:        10,
							X2:        20,
							Y2:        20,
							Text:      "hello world",
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "success_not_set_block_text_align",
			fields: fields{
				Config: struct {
					Debug     bool
					Size      float64
					Spacing   float64
					TextAlign string `yaml:"textAlign"`
					Blocks    []Block
				}{
					Debug:     false,
					Size:      10,
					Spacing:   10,
					TextAlign: "left",
					Blocks: []Block{
						{
							Size:    10,
							Spacing: 20,
							X1:      10,
							Y1:      10,
							X2:      20,
							Y2:      20,
							Text:    "hello world",
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "invalid_block_hcoords",
			fields: fields{
				Config: struct {
					Debug     bool
					Size      float64
					Spacing   float64
					TextAlign string `yaml:"textAlign"`
					Blocks    []Block
				}{
					Debug:     false,
					Size:      10,
					Spacing:   10,
					TextAlign: "left",
					Blocks: []Block{
						{
							Size:      10,
							Spacing:   20,
							TextAlign: "left",
							X1:        10,
							Y1:        10,
							X2:        10,
							Y2:        20,
							Text:      "hello world",
						},
					},
				},
			},
			wantErr: invalidHorizontalCoordsErr,
		},
		{
			name: "invalid_block_vcoords",
			fields: fields{
				Config: struct {
					Debug     bool
					Size      float64
					Spacing   float64
					TextAlign string `yaml:"textAlign"`
					Blocks    []Block
				}{
					Debug:     false,
					Size:      10,
					Spacing:   10,
					TextAlign: "left",
					Blocks: []Block{
						{
							Size:      10,
							Spacing:   20,
							TextAlign: "left",
							X1:        10,
							Y1:        10,
							X2:        20,
							Y2:        10,
							Text:      "hello world",
						},
					},
				},
			},
			wantErr: invalidVerticalCoordsErr,
		},
		{
			name: "invalid_common_text_align",
			fields: fields{
				Config: struct {
					Debug     bool
					Size      float64
					Spacing   float64
					TextAlign string `yaml:"textAlign"`
					Blocks    []Block
				}{
					Debug:     false,
					Size:      10,
					Spacing:   10,
					TextAlign: "invalid",
					Blocks: []Block{
						{
							Size:    10,
							Spacing: 20,
							X1:      10,
							Y1:      10,
							X2:      20,
							Y2:      20,
							Text:    "hello world",
						},
					},
				},
			},
			wantErr: invalidTextAlignErr,
		},
		{
			name: "empty_common_text_align",
			fields: fields{
				Config: struct {
					Debug     bool
					Size      float64
					Spacing   float64
					TextAlign string `yaml:"textAlign"`
					Blocks    []Block
				}{
					Debug:   false,
					Size:    10,
					Spacing: 10,
					Blocks: []Block{
						{
							Size:    10,
							Spacing: 20,
							X1:      10,
							Y1:      10,
							X2:      20,
							Y2:      20,
							Text:    "hello world",
						},
					},
				},
			},
			wantErr: textAlignNotSetErr,
		},
		{
			name: "invalid_text_align",
			fields: fields{
				Config: struct {
					Debug     bool
					Size      float64
					Spacing   float64
					TextAlign string `yaml:"textAlign"`
					Blocks    []Block
				}{
					Debug:     false,
					Size:      10,
					Spacing:   10,
					TextAlign: "left",
					Blocks: []Block{
						{
							Size:      10,
							Spacing:   20,
							TextAlign: "invalid",
							X1:        10,
							Y1:        10,
							X2:        20,
							Y2:        20,
							Text:      "hello world",
						},
					},
				},
			},
			wantErr: invalidTextAlignErr,
		},
		{
			name: "common_text_size_not_set",
			fields: fields{
				Config: struct {
					Debug     bool
					Size      float64
					Spacing   float64
					TextAlign string `yaml:"textAlign"`
					Blocks    []Block
				}{
					Debug:     false,
					Spacing:   10,
					TextAlign: "left",
					Blocks: []Block{
						{
							Size:    10,
							Spacing: 20,
							X1:      10,
							Y1:      10,
							X2:      20,
							Y2:      20,
							Text:    "hello world",
						},
					},
				},
			},
			wantErr: fontSizeNotSetErr,
		},
		{
			name: "common_spacing_not_set",
			fields: fields{
				Config: struct {
					Debug     bool
					Size      float64
					Spacing   float64
					TextAlign string `yaml:"textAlign"`
					Blocks    []Block
				}{
					Debug:     false,
					Size:      10,
					TextAlign: "left",
					Blocks: []Block{
						{
							Size:    10,
							Spacing: 20,
							X1:      10,
							Y1:      10,
							X2:      20,
							Y2:      20,
							Text:    "hello world",
						},
					},
				},
			},
			wantErr: spacingNotSetErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConfigApp{
				Config: tt.fields.Config,
			}
			err := c.Validate()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
