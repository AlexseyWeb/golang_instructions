package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "rmgOk.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	LstdFlags := log.Ldate | log.Lshortfile
	iLog := log.New(f, "Lnum ", LstdFlags)
	iLog.Println("Mastering of GO")

	iLog.SetFlags(log.Lshortfile | log.LstdFlags)
	iLog.Println("Another log entry!")
}
