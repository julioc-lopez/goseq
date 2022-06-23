package main

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

func openSourceFile(filename string) (io.ReadCloser, error) {
	if (filename == "") || (filename == "-") {
		return io.NopCloser(os.Stdin), nil
	}

	return os.Open(filename)
}

func chooseRendererBaseOnOutfile(filename string) (Renderer, error) {
	switch filepath.Ext(filename) {
	case ".png":
		return PngRenderer, nil
	case ".svg":
		return SvgRenderer, nil
	}

	return nil, errors.New("Unsupported extension: " + filename)
}
