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
func interpreterInit(text string) Continue_interpretation {
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

//starts interperating and sets the last index to the current index in case the line is incorrect
func (c Continue_interpretation) start() Continue_interpretation {
	i := c.self
	i.lastIndex = i.currentIndex
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
	i := c.self
	found := false
	if c.confirm == true {

		var result strings.Builder
		for a := 0; a < len(end_reading_at); a++ {
			for true {

				//if end_reading_at is in the line then say it is found and stop the loop
				//If it is not found before the token check has ended, break and set the currendIndex to lastIndex
				if i.currentIndex+len(end_reading_at[a]) > len(i.line) {
					break
				} else if end_reading_at[a] == string(i.line[i.currentIndex:i.currentIndex+len(end_reading_at[a])]) {
					found = true
					break
				} else {
					fmt.Println("looking for: ", end_reading_at[a], " found: ", string(i.line[i.currentIndex:i.currentIndex+len(end_reading_at[a])]), "len: ", len(string(i.line[i.currentIndex:i.currentIndex+len(end_reading_at[a])])), " loc: ", i.currentIndex)

					result.WriteString(string(i.line[i.currentIndex]))
					i.currentIndex += 1
				}
			}
			if found == true {
				break
			}
		}
		//adds the collected token to the token list
		i.elements = append(i.elements, result.String())

		//creates a new Continue_interpretation with the edited values
		return Continue_interpretation{
			confirm: found,
			self:    i,
		}
	}

	return Continue_interpretation{
		confirm: false,
		self:    i,
	}
}

//Checks for a keyword until finds it
func (c Continue_interpretation) keyword_check(keyword []string) Continue_interpretation {
	//tests for string(s)
	//multipule keys can do the same thing in the default syntax for this programming language
	//ex $ and func both create a function, $ is just the shortened version
	i := c.self
	found := false
	if c.confirm == true {

		for a := 0; a < len(keyword); a++ {

			for true {

				//if end_reading_at is in the line then say it is found and stop the loop
				//If it is not found before the token check has ended, break and set the currendIndex to lastIndex
				if i.currentIndex+len(keyword[a]) > len(i.line) {
					break
				} else if keyword[a] == string(i.line[i.currentIndex:i.currentIndex+len(keyword[a])]) {
					found = true
					break
				} else {
					//fmt.Println("kw looking for: ", keyword[a], " found: ", string(i.line[i.currentIndex:i.currentIndex+len(keyword[a])]), "len: ", len(string(i.line[i.currentIndex:i.currentIndex+len(keyword[a])])))
					i.currentIndex += 1
				}
			}
			if found == true {

				break
			}
		}

		fmt.Println(i.line[i.currentIndex])
		//creates a new Continue_interpretation with the edited values
		return Continue_interpretation{
			confirm: found,
			self:    i,
		}
	}

	return Continue_interpretation{
		confirm: false,
		self:    i,
	}
}

//stops checking for keyword if the interpereter finds a certain string
func (c Continue_interpretation) bounded_keyword_check(keyword []string, stop_checking_at []string) Continue_interpretation {
	//tests for string(s)
	//multipule keys can do the same thing in the default syntax for this programming language
	//ex $ and func both create a function, $ is just the shortened version
	i := c.self
	found := false
	if c.confirm == true {

		for a := 0; a < len(keyword); a++ {
			for true {
				//stopper right here
				for x := 0; x < len(stop_checking_at); x += 1 {

					if string(i.line[i.currentIndex:i.currentIndex+len(stop_checking_at[x])]) == stop_checking_at[x] {
						fmt.Println(i.line[i.currentIndex : i.currentIndex+len(stop_checking_at[x])])
						return Continue_interpretation{
							confirm: false,
							self:    i,
						}

					}
				}

				//if end_reading_at is in the line then say it is found and stop the loop
				//If it is not found before the token check has ended, break and set the currendIndex to lastIndex
				if i.currentIndex+len(keyword[a]) > len(i.line) {
					fmt.Println("oke2")
					break
				} else if keyword[a] == string(i.line[i.currentIndex:i.currentIndex+len(keyword[a])]) {
					fmt.Println("oke")
					found = true
					break
				} else {
					//fmt.Println("kw looking for: ", keyword[a], " found: ", string(i.line[i.currentIndex:i.currentIndex+len(keyword[a])]), "indx ", i.currentIndex)
					i.currentIndex += 1
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
		self:    i,
	}
}

//ends the interpreter/parser/whatever and executes a function
func (c Continue_interpretation) end(supply func(params []string)) Continue_interpretation {
	i := c.self
	if c.confirm == true {

		i.currentIndex += 1
		fmt.Println(string(i.line[i.currentIndex]))
		supply(c.self.elements)
	} else {
		i.currentIndex = i.lastIndex
	}

	//fmt.Println(string(i.line[i.currentIndex]))
	//fmt.Println(i.currentIndex)
	return Continue_interpretation{
		confirm: true,
		self:    i,
	}
}

type block struct {
	name      string
	ref       string
	execute   string
	constants []constant
	variables []variable
}
type variable struct {
	name  string
	ref   string
	value constant
}
type constant struct {
	name string
	data string
}
