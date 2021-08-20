package main

import (
	"fmt"
	"strconv"
)

var blocks []block

func compile(i *Continue_interpretation, s *syntax, b block) {
	intr := *i
	synt := *s
	bloc := b

	//the amount of memory the function takes up
	var memory int64
	memory = 0

	for true {

		//checks for a variable
		intr = intr.start().bounded_keyword_check(synt.create_variable, synt.end_line).store_with_keyword_check(synt.set_variable).store_with_keyword_check(synt.end_line).end(func(params []string) {
			fmt.Println("created variable")
			variable_name := params[0]
			variable_value := params[1]
			memory += 4

			cons, val := process_value(variable_value, synt)
			bloc.variables = append(bloc.variables, create_variable(variable_name, "string", val, false))
			bloc.constants = append(bloc.constants, cons...)

			bloc.execute += "\n   mov DWORD [ebp-" + strconv.FormatInt(memory, 10) + "], " + val
		})

		// you can only create a function if the current function is start
		if bloc.name == "_start" {

			//checks for a function
			intr = intr.start().bounded_keyword_check(synt.create_function, synt.end_line).store_with_keyword_check(synt.parameters_start).store_with_keyword_check(synt.parameters_end).keyword_check(synt.chunk_start).store_with_keyword_check(synt.chunk_end).end(func(params []string) {
				fmt.Println("created function")
			})
		}

		if intr.self.currentIndex >= len(intr.self.line)-3 {
			break
		}

	}
	bloc.execute_begenning += "   sub esp, " + strconv.FormatInt(memory, 10)

	blocks = append(blocks, bloc)
}
