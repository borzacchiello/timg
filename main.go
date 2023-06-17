package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var data []byte
	var err error

	ascii := flag.Bool("ascii", false, "ASCII output")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"USAGE: %s [-ascii] <file>\n"+
				"  -ascii: ascii only output\n"+
				"    file: name of the image file to display (or '-' to read from stdin)\n", os.Args[0])
	}

	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	fname := flag.Arg(0)
	fmt.Printf("fname: %s\n", fname)
	if fname == "-" {
		data, err = io.ReadAll(os.Stdin)
	} else {
		data, err = os.ReadFile(fname)
	}
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	img, err := LoadImage(data)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	if *ascii {
		err = DisplayImage(img)
	} else {
		err = DisplayImageRGB(img)
	}
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
}
