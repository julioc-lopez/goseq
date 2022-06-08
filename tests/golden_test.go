package tests

import (
	"bufio"
	"bytes"
	"flag"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/kylelemons/godebug/diff"
)

var (
	update = flag.Bool("update", false, "update the golden files of this test")
)

const gosecBin = "goseq-test.exe"

func TestGolden(t *testing.T) {
	testBin := buildTestBin(t)

	// process test .seq files
	entries, err := fs.Glob(os.DirFS("."), "testdata/input/*.seq")
	noError(t, err)

	// update golden files?
	if *update {
		for _, e := range entries {
			runNoError(t, testBin, "-o", filepath.Join("testdata/golden", filepath.Base(e)+".svg"), e)
		}

		t.Skip("Re-generated golden files")
	}

	// build results page
	f, err := os.Create("testout.html")
	noError(t, err)

	defer f.Close()

	w := bufio.NewWriter(f)

	const header = `<html>
<head>
  <style>
    table { border: solid thin black; border-collapse: collapse; }
    td { border: solid; }
  </style>
</head>
<body>
`

	n, err := w.WriteString(header)
	noError(t, err)

	if want, got := len(header), n; want != got {
		t.Fatalf("Short write to results buffer, want %v, got %v", want, got)
	}

	for _, e := range entries {
		wantFile := filepath.Join("testdata/golden", filepath.Base(e)+".svg")
		want, err := ioutil.ReadFile(wantFile)
		noError(t, err)

		got := runOut(t, testBin, e)

		if !bytes.Equal(got, want) {
			t.Log(diff.Diff(string(want), string(got)))
			t.Fatalf("%s %q output does not match %q", testBin, e, wantFile)
		}

		w.WriteString("<p>")
		w.WriteString(e)
		w.WriteString("</p>\n<table><tr><td><pre>\n")

		// insert input file
		in, err := os.Open(e)
		noError(t, err)

		w.ReadFrom(in)
		w.WriteString("</pre></td><td>\n")
		// add resulting output
		w.Write(bytes.TrimPrefix(got, []byte("<?xml version=\"1.0\"?>\n")))
		w.WriteString("</td></tr></table>\n")
	}

	w.WriteString("</body>\n</html>\n")
	noError(t, w.Flush())
}

func buildTestBin(t *testing.T) string {
	out := filepath.Join(t.TempDir(), gosecBin)

	runNoError(t, "go", "build", "-o", out, "../")

	return out
}

// noError logs and bails out when err is not nil
func noError(tb testing.TB, err error, message ...interface{}) {
	tb.Helper()

	if err != nil {
		tb.Log("unexpected error:", err)
		tb.Fatal(message...)
	}
}

// run runs `cmd arg...`. The resulting output and stderr are sent to
// the stdout and stderr of this process. An error is returned if there's a
// problem, see exec.Cmd.run() for details.
func run(tb testing.TB, name string, arg ...string) error {
	tb.Helper()

	c := exec.Command(name, arg...)
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout

	tb.Log(str2any(c.Args)...)

	return c.Run()
}

// runNoError runs a command with args. Logs and bails out on error.
func runNoError(tb testing.TB, name string, arg ...string) {
	tb.Helper()

	noError(tb, run(tb, name, arg...))
}

// runOut runs `cmd arg...`. The resulting  stderr is sent to the stderr of this
// process. Stdout is returned. On error, it logs and panics.
func runOut(tb testing.TB, name string, arg ...string) []byte {
	c := exec.Command(name, arg...)
	c.Stderr = os.Stderr

	tb.Log(str2any(c.Args)...)

	out, err := c.Output()
	noError(tb, err, "Running ", name)

	return out
}

// Str2interface converts a slice of strings to a slice of interfaces
func str2any(s []string) []any {
	anys := make([]any, 0, len(s))

	for _, x := range s {
		anys = append(anys, x)
	}
	return anys
}
