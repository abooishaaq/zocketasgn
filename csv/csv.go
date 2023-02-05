package main

import (
	"fmt"
	"log"
	"os"
)

type csv struct {
	file  *os.File
	Comma rune
	Data  [][]string
}

func NewCSV(file *os.File) *csv {
	return &csv{
		file:  file,
		Comma: ',',
		Data:  [][]string{},
	}
}

func (c *csv) Read() ([]string, error) {
	var (
		line []string
		prev []byte
	)

	for {
		buf := make([]byte, 1)
		_, err := c.file.Read(buf)
		if err != nil {
			return nil, err
		}

		if buf[0] == byte(c.Comma) {
			line = append(line, string(prev))
			prev = []byte{}
		} else if buf[0] == '\n' {
			line = append(line, string(prev))
			break
		} else {
			prev = append(prev, buf[0])
		}
	}

	return line, nil
}

func (c *csv) ReadAll() ([][]string, error) {
	for {
		line, err := c.Read()
		if err != nil {
			return c.Data, err
		}
		c.Data = append(c.Data, line)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No file specified")
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := NewCSV(file)

	data, err := reader.ReadAll()
	if err != nil {
		if err.Error() != "EOF" {
			log.Fatal(err)
		}
	}

	for _, each := range data {
		for _, value := range each {
			fmt.Print(value, "\t")
		}
		fmt.Println()
	}
}
