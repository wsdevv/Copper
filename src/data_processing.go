package main

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
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
	name      string
	ref       string
	memory    int
	id        string
	execute   []string
	constants []constant
	variables []variable
}
type variable struct {
	name     string
	ref      string
	typ      string
	value    string
	location int
	id       int
}
type constant struct {
	name string
	data string
	col  bool
}

func create_block(name string, keep_name bool) block {
	if keep_name {
		return block{
			name: name,
			ref:  name,
		}
	}
	return block{
		name: create_data_name(int(math.Abs(float64((int64(rand.Intn(16))*time.Now().UnixNano()/int64(1000000))%16)) + 5)),
		ref:  name,
	}
}
func create_variable(name string, typ string, value string, keep_name bool, location int) variable {
	if keep_name {
		return variable{
			name:     name,
			ref:      name,
			typ:      typ,
			value:    value,
			location: location,
		}
	}
	return variable{
		name:     name,
		ref:      create_data_name(int(math.Abs(float64((int64(rand.Intn(16))*time.Now().UnixNano()/int64(1000000))%16)) + 5)),
		id:       int((int64(int64(rand.Intn(1024)) * time.Now().UnixNano() / int64(1000000)))),
		typ:      typ,
		value:    value,
		location: location,
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

//returns as compiled paintL
func (v variable) compile_data() string {
	defaults := get_paintlc_defaults()
	return strings.Join([]string{string(rune(defaults.x)), v.ref, string(rune(defaults.var_name)), strconv.Itoa(v.location), string(rune(defaults.variable_mem_loc)), v.value, string(rune(defaults.variable_value)), string(rune(defaults.variable))}, "")
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
