package main

import (
	"io"

	"os"
	"fmt"
	"bytes"
)
type UpperWriter struct {
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToLower(data))
}

func main() {
	fmt.Fprintln(&UpperWriter{os.Stdout}, "Hello, World")
}
