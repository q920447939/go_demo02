package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	file, _ := os.OpenFile("C:/test.log", os.O_CREATE|os.O_WRONLY, 0755)
	file.WriteString("hello golang")
	file.Close()

	var _ io.Writer = (*bytes.Buffer)(nil)
}
