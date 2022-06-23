// Renderers disabled if noim is specified
//

//go:build !im
// +build !im

package main

import (
	"errors"

	"github.com/lmika/goseq/seqdiagram"
)

func PngRenderer(diagram *seqdiagram.Diagram, opts *seqdiagram.ImageOptions, target string) error {
	return errors.New("PNG renderer not available")
}
