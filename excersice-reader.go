package main

import "github.com/golang-id/tour/reader"

type MyReader struct{}

func (m MyReader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		bytes[i] = 65
	}
	return len(bytes), nil
}

func main() {
	reader.Validate(MyReader{})
}
