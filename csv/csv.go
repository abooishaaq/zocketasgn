package main

import (
	"bufio"
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

func (c *csv) ReadAll() ([][]string, error) {
	reader := bufio.NewReader(c.file)

	for {
		line, err := reader.ReadBytes(byte('\n'))

		if err != nil {
			return c.Data, err
		}

		line = line[:len(line)-1]
		data := make([]string, 0)
		prev := []byte{}

		for i := 0; i < len(line); i++ {
			if line[i] == byte(c.Comma) {
				data = append(data, string(prev))
				prev = []byte{}
			} else {
				prev = append(prev, line[i])
			}
		}
		if err != nil {
			return c.Data, err
		}
		c.Data = append(c.Data, data)
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
