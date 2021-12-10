package main

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func openSourceFile(filename string) (io.ReadCloser, error) {
	if (filename == "") || (filename == "-") {
		return ioutil.NopCloser(os.Stdin), nil
	}

	return os.Open(filename)
}

func chooseRendererBaseOnOutfile(filename string) (Renderer, error) {
	ext := filepath.Ext(filename)
	if ext == ".png" {
		return PngRenderer, nil
	} else if ext == ".svg" {
		return SvgRenderer, nil
	}

	return nil, errors.New("Unsupported extension: " + filename)
}
