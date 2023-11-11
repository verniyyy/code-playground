package src

import (
	"context"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_codeRunner_writeFile(t *testing.T) {
	type args struct {
		src  io.Reader
		path string
	}
	tests := []struct {
		name    string
		_       codeRunner
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "create file and write on src",
			args: args{
				src:  newIOReaderFromString("foo\nbar"),
				path: "test.tmp",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := codeRunner{}
			if err := c.writeFile(tt.args.src, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("codeRunner.writeFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := os.Remove(tt.args.path); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func newIOReaderFromString(s string) io.Reader {
	return strings.NewReader(s)
}

func Test_codeRunner_RunGo(t *testing.T) {
	type args struct {
		ctx context.Context
		src io.Reader
	}
	tests := []struct {
		name    string
		_       codeRunner
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "can running Go code",
			args: args{
				ctx: context.TODO(),
				src: newIOReaderFromString(`
					package main

					import "fmt"

					func main() {
						fmt.Println("Hello world!")
					}
				`),
			},
			want:    "Hello world!\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := codeRunner{}
			got, err := cr.RunGo(tt.args.ctx, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("codeRunner.RunGo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("codeRunner.RunGo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_codeRunner_RunRuby(t *testing.T) {
	type args struct {
		ctx context.Context
		src io.Reader
	}
	tests := []struct {
		name    string
		_       codeRunner
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "can running ruby code",
			args: args{
				ctx: context.TODO(),
				src: newIOReaderFromString(`
					puts "Hello world!"
				`),
			},
			want:    "Hello world!\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := codeRunner{}
			got, err := cr.RunRuby(tt.args.ctx, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("codeRunner.RunRuby() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("codeRunner.RunRuby() = %v, want %v", got, tt.want)
			}
		})
	}
}
