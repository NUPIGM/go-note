package notes2

import (
	"errors"
	"log"
	"os"
)

var (
	INFO *log.Logger
	WARN *log.Logger
	ERR  *log.Logger
)

func In() {
	file, err := os.OpenFile("./file/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		log.Panic(err)
	}
	INFO = log.New(file, "INFO:", log.LstdFlags|log.Llongfile)
	WARN = log.New(file, "WARN:", log.LstdFlags|log.Llongfile)
	ERR = log.New(file, "ERR:", log.LstdFlags|log.Llongfile)
}
func Log() {
	defer func() {
		err := recover()
		log.Println("捕捉到了：", err)
	}()
	err := errors.New("错误")
	INFO.Println(err)
	// WARN.Panicln(err)
	// ERR.Fatalln(err)
}
