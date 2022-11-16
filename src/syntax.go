package main

/*
 * FILE: syntax.go
 * PURPOSE: holds functions to build the syntax.json file into syntax rules
 */

type SyntaxBlock struct {
	next_op SyntaxBlock,
	properties map[string]string
}

func build_syntax() {
	
}