package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var data []byte
	var err error

	if len(os.Args) < 2 {
		data, err = io.ReadAll(os.Stdin)
	} else {
		fname := os.Args[1]
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

	err = DisplayImage(img)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
}
