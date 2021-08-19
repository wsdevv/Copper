package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var blocks []block

func compile(i *Continue_interpretation, s *syntax, bid int) {
	intr := *i
	synt := *s

	//checks for a variable
	intr = intr.start().bounded_keyword_check(synt.create_variable, synt.end_line).store_with_keyword_check(synt.set_variable).store_with_keyword_check(synt.end_line).end(func(params []string) {
		fmt.Println("created variable")
		variable_name := params[0]
		variable_value := params[1]
	})

	//checks for a function
	intr = intr.start().bounded_keyword_check(synt.create_function, synt.end_line).store_with_keyword_check(synt.parameters_start).store_with_keyword_check(synt.parameters_end).keyword_check(synt.chunk_start).store_with_keyword_check(synt.chunk_end).end(func(params []string) {
		fmt.Println("created function")
	})
}

func main() {
	str, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	read := string(str)
	intr := interpreterInit(read)
	synt := defaultSyntax()

	//creates the starting point
	blocks = append(blocks, block{
		name:    "_start",
		execute: "",
	})

	for true {
		compile(&intr, &synt, 0)
		if intr.self.currentIndex >= len(read)-3 {
			break
		}
	}
}
