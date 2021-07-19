package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	str, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	read := string(str)
	intr := startInterpreter(read)
	synt := defaultSyntax()

	for true {
		//checks for a variable
		intr.keyword_check(synt.create_variable).store_with_keyword_check(synt.set_variable).store_with_keyword_check(synt.end_line).end(func(params []string) {
			fmt.Println("ooh yeahh")
		})
	}
}
