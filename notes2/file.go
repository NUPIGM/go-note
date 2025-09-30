package notes2

import (
	"bufio"
	"fmt"
	"os"
)

func File() {
	f, _ := os.OpenFile("text.txt", os.O_CREATE|os.O_RDWR, 0777)
	// defer f.Close()
	// f.Write([]byte("3333"))

	bufio.NewReader(f)

}
func Status() {
	info, _ := os.Stat("./file.go")
	fmt.Println(info.Size(), info.Name(), info.Mode())
}
