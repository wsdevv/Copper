package main

import (
	"math"
	"math/rand"
	"time"
)

//This file holds the methods to store variables and functions

func create_data_name(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	final := ""
	for i := 0; i < length; i += 1 {

		final += string(chars[int(math.Abs(float64((int64(rand.Intn((len(chars))))*time.Now().UnixNano()/int64(1000000))%int64(len(chars)))))])
	}
	return final
}

type block struct {
	name              string
	ref               string
	execute           string
	execute_begenning string
	execute_end       string
	constants         []constant
	variables         []variable
}
type variable struct {
	name  string
	ref   string
	typ   string
	value string
}
type constant struct {
	name string
	data string
	col  bool
}

func create_block(name string, keep_name bool) block {
	if keep_name {
		return block{
			name:              name,
			ref:               name,
			execute_begenning: "\n   push ebp\n   mov esp, ebp\n",
			execute:           "",
			execute_end:       "\n   mov ebp, esp\n   pop ebp\n   ret ;",
		}
	}
	return block{
		name:        create_data_name(int(math.Abs(float64((int64(rand.Intn(16))*time.Now().UnixNano()/int64(1000000))%16)) + 5)),
		ref:         name,
		execute:     "   push ebp\n   mov esp, ebp\n",
		execute_end: "\n   mov ebp, esp\n   pop ebp\n   ret ;",
	}
}
func create_variable(name string, typ string, value string, keep_name bool) variable {
	if keep_name {
		return variable{
			name:  name,
			ref:   name,
			typ:   typ,
			value: value,
		}
	}
	return variable{
		name:  name,
		ref:   create_data_name(int(math.Abs(float64((int64(rand.Intn(16))*time.Now().UnixNano()/int64(1000000))%16)) + 5)),
		typ:   typ,
		value: value,
	}
}

//randomed named constant and named constant, helps in creating lengths
func create_constant_nn(data string, col bool) constant {
	return constant{
		name: create_data_name(int(math.Abs(float64((int64(rand.Intn(16))*time.Now().UnixNano()/int64(1000000))%16)) + 5)),
		data: data,
		col:  col,
	}
}
func create_constant_n(name string, data string, col bool) constant {
	return constant{
		name: name,
		data: data,
		col:  col,
	}
}

func process_value(value string, synt syntax) ([]constant, string) {
	var constants []constant

	data_val := ""
	val := ""
	creating_string := false

	for i := 0; i < len(value); i++ {
		if creating_string == true {
			data_val += string(value[i])
		}
		for x := 0; x < len(synt.declare_string); x++ {
			//creating a string variable
			if string(value[i:i+len(synt.declare_string[x])]) == synt.declare_string[x] {
				if creating_string == false {
					creating_string = true
					data_val += "\""
				} else {
					creating_string = false
					data_val += "\""
					//IMPORTANT: Only supports strings NEED to implament ints, etc
					constants = append(constants, create_constant_nn("\n  DB "+data_val, true))

				}
			}
		}

	}
	//adds the constant name to variable value and adds the constant length to constants
	var constants_len []constant
	for i := 0; i < len(constants); i += 1 {

		//gets the length of a constant
		constants_len = append(constants_len, create_constant_n(constants[i].name+".length", " equ $-"+constants[i].name, false))
		val += "." + constants[i].name
	}

	constants = append(constants, constants_len...)
	return constants, val
}
