package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"io"
)

func main() {
	name := "/etc/passwd"
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer f.Close()
	fmt.Println(f)
	fReader := bufio.NewReader(f)
	for {
		fStr, fErr := fReader.ReadString('\n')
		fmt.Println(fStr)
		if fErr == io.EOF{
			return
		}
	}
	if err := f.Close(); err != nil {
		fmt.Println(err)
	}
}
