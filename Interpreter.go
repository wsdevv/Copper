package main

import (
	"fmt"
	"strings"
)

type Interpreter struct {
	elements     []string
	line         string
	currentIndex int
	lastIndex    int
}

//Says if the current line interpreter should Continue checking for the token/string
type Continue_interpretation struct {
	confirm bool
	self    Interpreter
}

//starts a interpreter by converting a new interpreter into Continue_interpretation struct
func startInterpreter(text string) Continue_interpretation {
	i := Interpreter{
		line:         text,
		currentIndex: 0,
		lastIndex:    0,
	}
	return Continue_interpretation{
		confirm: true,
		self:    i,
	}
}

//test if a line starts with a certain token
//line           = the line the function is grabing from
//end_reading_at = ends reading token at a certain character or string(s)
func (c Continue_interpretation) store_with_keyword_check(end_reading_at []string) Continue_interpretation {

	//tests for string(s)
	//multipule keys can do the same thing in the default syntax for this programming language
	//ex $ and func both create a function, $ is just the shortened version
	if c.confirm == true {
		//clones c's interpreter then Continues
		i := c.self
		found := false
		i.lastIndex = i.currentIndex
		var result strings.Builder
		for a := 0; a < len(end_reading_at); a++ {
			for true {

				//if end_reading_at is in the line then say it is found and stop the loop
				//If it is not found before the token check has ended, break and set the currendIndex to lastIndex
				if end_reading_at[a] == string(i.line[i.currentIndex:len(end_reading_at[a])]) {
					found = true
					break
				} else if 1+i.currentIndex+len(end_reading_at[a]) > len(i.line) {
					i.currentIndex = i.lastIndex
					break
				} else {
					result.WriteString(string(i.line[i.currentIndex]))
					i.currentIndex += 1
				}
			}
			if found == true {
				break
			}
		}
		//adds the collected token to the token list
		fmt.Println(result.String())
		i.elements = append(i.elements, result.String())

		//creates a new Continue_interpretation with the edited values
		return Continue_interpretation{
			confirm: found,
			self:    i,
		}
	}

	return Continue_interpretation{
		confirm: false,
	}
}

//basically the same as store_with_keyword_check but does not store a value within the value list
func (c Continue_interpretation) keyword_check(keyword []string) Continue_interpretation {
	//tests for string(s)
	//multipule keys can do the same thing in the default syntax for this programming language
	//ex $ and func both create a function, $ is just the shortened version
	if c.confirm == true {
		//clones c's interpreter then Continues
		i := c.self
		found := false
		i.lastIndex = i.currentIndex

		for a := 0; a < len(keyword); a++ {
			for true {

				//if end_reading_at is in the line then say it is found and stop the loop
				//If it is not found before the token check has ended, break and set the currendIndex to lastIndex
				if keyword[a] == string(i.line[i.currentIndex:len(keyword[a])]) {
					found = true
					break
				} else if 1+i.currentIndex+len(keyword[a]) > len(i.line) {
					i.currentIndex = i.lastIndex
					break
				}
			}
			if found == true {
				break
			}
		}

		//creates a new Continue_interpretation with the edited values
		return Continue_interpretation{
			confirm: found,
			self:    i,
		}
	}

	return Continue_interpretation{
		confirm: false,
	}
}

//ends the interpreter/parser/whatever
func (c Continue_interpretation) end(supply func(params []string)) {
	if c.confirm == true {
		supply(c.self.elements)
	}
}
