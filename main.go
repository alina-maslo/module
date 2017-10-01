package main

import (
	"os"
	"fmt"
	//"flag"
	//"strings"
)

func MsgPrintln(msg string) {
	fmt.Println("\033[1;92m" + msg + "\033[0m")
}

func main() {
	MsgPrintln("test test test")
	fmt.Println("%#v", os.Args)
}
