package models

import (
	"bufio"
	"os"
)

type FileReader struct {
	scanner *bufio.Scanner
	file    *os.File
}

func (fr *FileReader) Close() {
	err := fr.file.Close()
	if err != nil {
		panic(err)
	}
}

func NewFilerReader(relativeLoc string) *FileReader {
	file, err := os.Open(relativeLoc)
	if err != nil {
		panic(err)
	}
	return &FileReader{bufio.NewScanner(file), file}
}

func (fr *FileReader) NextLine() *string {
	if fr.scanner.Scan() {
		line := fr.scanner.Text()
		return &line
	} else {
		if fr.scanner.Err() != nil {
			panic(fr.scanner.Err())
		}
		return nil
	}
}
