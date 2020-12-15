package reader

import (
	"bufio"
	"log"
	"os"
)

type InputReader struct {
	filename   string
	lineparser func([]byte) interface{}
}

func NewInputReader(filename string, lineparser func([]byte) interface{}) *InputReader {
	return &InputReader{
		filename:   filename,
		lineparser: lineparser,
	}
}

func (i *InputReader) ParseInput() []interface{} {
	file, err := os.Open(i.filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var output []interface{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineBytes := scanner.Bytes()
		output = append(output, i.lineparser(lineBytes))
	}
	return output
}
