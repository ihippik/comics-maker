package main

import "testing"

func Test_isValidUrl(t *testing.T) {
	type args struct {
		txt string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success_file",
			args: args{
				txt: "http://template.com/img.png",
			},
			want: true,
		},
		{
			name: "success_https",
			args: args{
				txt: "https://template.com/img.png",
			},
			want: true,
		},
		{
			name: "success_path",
			args: args{
				txt: "http://template.com/image",
			},
			want: true,
		},
		{
			name: "fail_file",
			args: args{
				txt: "image.png",
			},
			want: false,
		},
		{
			name: "fail_without_scheme",
			args: args{
				txt: "emplate.com/image",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidUrl(tt.args.txt); got != tt.want {
				t.Errorf("isValidUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
