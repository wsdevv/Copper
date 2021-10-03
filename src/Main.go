package main

//WINDOWS DIST
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

	//creates the starting point
	compile(&intr, &synt, create_block("_main", true))

	fmt.Println(blocks)

	file, err := os.Create("./test/test.asm")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(result)

}
