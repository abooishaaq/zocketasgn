package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

func download(file string, done chan bool) {
	url_parts := strings.Split(file, "/")
	file_name := url_parts[len(url_parts)-1]
	out, err := os.Create(file_name)
	defer out.Close()
	if err != nil {
		panic(err)
	}
	resp, err := http.Get(file)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bufio.NewReader(resp.Body).WriteTo(out)
	done <- true
}

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		panic("No files to download")
	}

	done := make(chan bool)
	for _, file := range files {
		go download(file, done)
	}
	done_count := 0
	for {
		if done_count == len(files) {
			break
		}
		<-done
		done_count++
	}
}
