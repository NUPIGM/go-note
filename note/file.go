package main

import (
	"bufio"
	"os"
)

func main() {
	f, _ := os.OpenFile("text.txt", os.O_CREATE|os.O_RDWR, 0777)
	// defer f.Close()
	// f.Write([]byte("3333"))

	bufio.NewReader(f)

}
