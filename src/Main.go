package main

//WINDOWS DIST
import (
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

	//another step to compiling the code
	for i := 0; i < len(blocks); i += 1 {
		result += "\n" + blocks[i].name + ":\n"
		result += blocks[i].execute_begenning
		result += blocks[i].execute
		result += blocks[i].execute_end
		for u := 0; u < len(blocks[i].constants); u += 1 {

			//checks if the constant is setting something that has to include a ":" to declare it
			if blocks[i].constants[u].col == true {
				result += "\n." + blocks[i].constants[u].name + ":" + blocks[i].constants[u].data
			} else {
				result += "\n." + blocks[i].constants[u].name + blocks[i].constants[u].data
			}
		}
	}

	file, err := os.Create("./test/test.asm")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(result)

}
