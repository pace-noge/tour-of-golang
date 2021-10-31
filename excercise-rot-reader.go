package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i, code := range p {
		if 'A' <= code && code <= 'M' || 'a' <= code && code <= 'm' {
			p[i] += 13
		} else if 'N' <= code && code <= 'Z' || 'n' <= code && code <= 'z' {
			p[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
