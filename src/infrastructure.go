package src

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"os/exec"
)

// NewCodeRunner ...
func NewCodeRunner() CodeRunner {
	return codeRunner{}
}

// codeRunner ...
type codeRunner struct{}

// RunGo ...
func (cr codeRunner) RunGo(ctx context.Context, src io.Reader) (string, error) {
	filePath := "/tmp/main.go"
	cr.writeFile(src, filePath)

	cmd := exec.Command("go", "run", filePath)
	stdout, err := cmd.Output()
	if err != nil {
		log.Println(stdout)
		return "", err
	}

	e := os.Remove(filePath)
	if e != nil {
		log.Println(e)
	}

	return fmt.Sprint(string(stdout)), nil
}

// RunGo ...
func (cr codeRunner) RunRuby(ctx context.Context, src io.Reader) (string, error) {
	filePath := "/tmp/main.rb"
	cr.writeFile(src, filePath)

	cmd := exec.Command("ruby", filePath)
	stdout, err := cmd.Output()
	if err != nil {
		log.Println(stdout)
		return "", err
	}

	e := os.Remove(filePath)
	if e != nil {
		log.Println(e)
	}

	return fmt.Sprint(string(stdout)), nil
}

// writeFile ...
func (codeRunner) writeFile(src io.Reader, path string) error {
	// open output file
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		str, err := url.PathUnescape(string(buf[:n]))
		if err != nil {
			return err
		}
		// write a chunk
		if _, err := f.Write([]byte(str)); err != nil {
			return err
		}
	}
	return nil
}
