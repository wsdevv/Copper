package main


import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syntax"
)

func main() {
	str, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	read := string(str)
	intr := interpreterInit(read)
	synt := defaultSyntax()
	result := "section .text:\n global _main \n"


	file, err := os.Create("./test/test.asm")
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()
	file.WriteString(result)

}
