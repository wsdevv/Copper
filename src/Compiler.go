package main

import (
	"fmt"
)

var blocks []block

func compile(i *Continue_interpretation, s *syntax, b block) {
	intr := *i
	synt := *s
	bloc := b

	//the amount of memory the function takes up

	for true {

		//checks for a variable, strings take up 8 bytes
		intr = intr.start().bounded_keyword_check(synt.create_variable, synt.end_line).store_with_keyword_check(synt.set_variable).store_with_keyword_check(synt.end_line).end(func(params []string) {
			//TODO IMPORTANT: Get variables to support ints
			fmt.Println("created variable")
			variable_name := params[0]
			variable_value := params[1]

			bloc.memory += 4
			cons, val := process_value(variable_value, synt)
			vari := create_variable(variable_name, "string", val, false, bloc.memory)

			bloc.variables = append(bloc.variables, vari)
			bloc.constants = append(bloc.constants, cons...)
			bloc.execute = append(bloc.execute, vari.compile_data())
		})
		fmt.Println(bloc.execute)
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

	blocks = append(blocks, bloc)
}
