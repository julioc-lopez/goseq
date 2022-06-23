package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/howeyc/fsnotify"
	"github.com/lmika/goseq/seqdiagram"
)

// Name of the output file
var flagOut = flag.String("o", "", "Output file")

// The style to use
var flagStyle = flag.String("s", "default", "The style to use")

// Generate an embedded SVG file
var flagEmbedded = flag.Bool("e", false, "Generate an embedded SVG file")

// Setup a watcher to regenerate the file when changed
var flagWatch = flag.Bool("w", false, "Watch for changes")

// Die with error
func die(msg string) {
	fmt.Fprintf(os.Stderr, "goseq: %s\n", msg)
	os.Exit(1)
}

// Construct and build image options based on the current configuration
func buildImageOptions() *seqdiagram.ImageOptions {
	// Work out the style
	style := seqdiagram.DefaultStyle
	if altStyle, hasStyle := seqdiagram.StyleNames[*flagStyle]; hasStyle {
		style = altStyle
	}

	return &seqdiagram.ImageOptions{
		Style:    style,
		Embedded: *flagEmbedded,
	}
}

// Processes a md file
func processMdFile(inFilename, outFilename string, renderer Renderer) error {
	srcFile, err := openSourceFile(inFilename)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	targetFile := io.Discard

	mf := &MarkdownFilter{srcFile, targetFile, func(codeblock string, output io.Writer) error {
		fmt.Fprint(output, codeblock)
		err := processSeqDiagram(strings.NewReader(codeblock), inFilename, "/dev/null", nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "goseq: %s:embedded block - %s\n", inFilename, err.Error())
		}
		return nil
	}}
	return mf.Scan()
}

// Processes a seq file
func processSeqFile(inFilename, outFilename string, renderer Renderer) error {
	srcFile, err := openSourceFile(inFilename)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	return processSeqDiagram(srcFile, inFilename, outFilename, renderer)
}

// Processes a sequence diagram
func processSeqDiagram(infile io.Reader, inFilename, outFilename string, renderer Renderer) error {
	diagram, err := seqdiagram.ParseDiagram(infile, inFilename)
	if err != nil {
		return err
	}

	// Image options
	imageOptions := buildImageOptions()

	// If there's a process instruction, use it as the target of the diagram
	// TODO: be a little smarter with the process instructions
	for _, pr := range diagram.ProcessingInstructions {
		if pr.Prefix == "goseq" {
			outFilename = pr.Value
		}
	}

	if renderer == nil {
		renderer, err = chooseRendererBaseOnOutfile(outFilename)
		if err != nil {
			return err
		}
	}

	err = renderer(diagram, imageOptions, outFilename)
	if err != nil {
		return err
	}

	return nil
}

// Processes a file.  This switches based on the file extension
func processFile(inFilename, outFilename string, renderer Renderer) error {
	ext := filepath.Ext(inFilename)
	if ext == ".md" {
		return processMdFile(inFilename, outFilename, renderer)
	} else {
		return processSeqFile(inFilename, outFilename, renderer)
	}
}

// Setup a watch process which will regenerate the files
func watchAndProcess(inFiles []string, outFile string, renderer Renderer) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	for _, inFile := range inFiles {
		if err := watcher.Watch(inFile); err != nil {
			log.Printf("Could not set notify watch on %s: %v", inFile, err)
		}
	}

	for event := range watcher.Event {
		if event.IsModify() {
			if err := processFile(event.Name, outFile, renderer); err != nil {
				log.Println(event.Name, "-", err.Error())

				continue
			}

			log.Println("Generating", event.Name, "->", outFile)
		}
	}

	return nil
}

func main() {
	var err error

	renderer := SvgRenderer
	outFile := ""

	flag.Parse()

	// Select a suitable renderer (based on the suffix of the output file, if there is one)
	if *flagOut != "" {
		renderer, err = chooseRendererBaseOnOutfile(*flagOut)
		if err != nil {
			die(err.Error())
		}
		outFile = *flagOut
	}

	// Process each file (or stdin)
	if flag.NArg() == 0 {
		err := processFile("-", outFile, renderer)
		if err != nil {
			die("stdin - " + err.Error())
		}
	} else {
		if *flagWatch {
			if err := watchAndProcess(flag.Args(), outFile, renderer); err != nil {
				log.Fatal("Could not watch and process:", err)
			}
		} else {
			for _, inFile := range flag.Args() {
				err := processFile(inFile, outFile, renderer)
				if err != nil {
					die(inFile + " - " + err.Error())
				}
			}
		}
	}
}
