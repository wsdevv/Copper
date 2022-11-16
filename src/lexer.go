package main

/*
 * FILE: lexer.go
 * PURPOSE: holds the lexer functions
 */
import "strings"
import "fmt"
import "os"

//import "github.com/go-errors/errors"

const (
	LogTypeNone             = 0
	LogTypeAll              = 1
	LogTypeSeperate         = 2
	LogTypeRemoveWhitespace = 3
	LogTypeNoSelf           = 4
)

/*
used in testing
*/
func main() {
	lexer := new_lexer("var no = 0; func hi(boi, gurl) {mmm}")

	lexer = lexer.add_rule(func(lex Lexer) Lexer {
		lex = lex.start().next_string("var", LogTypeSeperate).next_string("=", LogTypeRemoveWhitespace+LogTypeNoSelf).next_string(";", LogTypeNoSelf)

		fmt.Println(string(lex.program[lex.position-1]))
		lex = lex.end()
		return lex
	})

	lexer = lexer.add_rule(func(lex Lexer) Lexer {
		lex = lex.start().next_string("func", LogTypeSeperate).next_string("(", LogTypeRemoveWhitespace+LogTypeNoSelf).next_string(")", LogTypeNoSelf).next_string("{", LogTypeNoSelf).next_string("}", LogTypeNoSelf)

		fmt.Println(string(lex.program[lex.position-1]))
		lex = lex.end()
		return lex
	})

	(&lexer).init()

	lexer.walk(func() {})
}

/**
 *
 * Lexer
 *
 * Purpose: struct used to initialize and continue char-by-char lexing
 *
 * ATTRIBUTES
 * pass_data: the data to pass to a interperater/compiler onece the lexing is finished
 * stop_codes: the code that stops the lexer if a token is not found, default is ";". (stop code can only be 1 char, though it is declared a string)
 * position: the current position
 * move:     the next-line position
 * program:  the code to be compiled
 * continue_lexer: tells to continue a function chain if in progress
 */
type Lexer struct {

	// Feilds for token handling and moving through text
	pass_data      []string
	stop_code      map[string]int
	continue_lexer bool
	position       int
	move           int
	program        string

	// Fields for rule handling
	rules   []func(Lexer) Lexer
	finish  int
	success bool
	check   bool

	// Fields for error handling
	start_called bool
	end_called   bool
}

/*
 * new_lexer
 * Purpose: initializes a lexing object/struct for use
 */
func new_lexer(program string) Lexer {

	stop_code := make(map[string]int)
	stop_code[";"] = 1

	return Lexer{
		pass_data:      []string{},
		stop_code:      stop_code,
		continue_lexer: true,
		position:       0,
		move:           0,
		program:        program,

		rules:   []func(Lexer) Lexer{},
		finish:  0,
		success: false,
		check:   false,

		start_called: false,
		end_called:   true,
	}

}

/*
 *
 */
func (lexer *Lexer) init() {

	// lexer has to have at least one rule for the init function to work
	if len(lexer.rules) < 1 {
		fmt.Println(fmt.Errorf("ERROR: Lexer has to have at least one rule for 'func init()' to be called."))
		os.Exit(-1)
	}

	for rule_index := 0; rule_index < len(lexer.rules); rule_index++ {

		go func(r_index int) {
			
			lex := &lexer
			lex_out := *lexer
			lex_rule := (**lex).rules[r_index]
			

			for true {
				if (**lex).check {
					lex_out = lex_rule(lex_out)
					if lex_out.continue_lexer {
						**lex = lex_out
						(**lex).check = false
						(**lex).success = true
					} else {
						(**lex).finish += 1
						for (**lex).check {
						}
					}
				}
			}

		}(rule_index)

	}
}

/*
 */
func (lexer *Lexer) walk(run func()) {

	// lexer has to have at least one rule for the walk function to work
	if len(lexer.rules) < 1 {
		fmt.Println(fmt.Errorf("ERROR: Lexer has to have at least one rule for 'func walk()' to be called."))
		os.Exit(-1)
	}

	for true {

		if lexer.position >= len(lexer.program) {
			break
		}
		lexer.check = true

		run()

		for lexer.success == false {
			if lexer.success == true {
				break
			}
			if lexer.finish >= len(lexer.rules) {
				return
			}
		}
	}
}

/*
 */
func (lexer Lexer) add_rule(rule_func func(Lexer) Lexer) Lexer {
	lexer.rules = append(lexer.rules, rule_func)
	return lexer
}

/*
 * Lexer method start
 * starts the lexer at position n
 * REQUIRED FOR LEXER TO WORK
 */
func (lexer Lexer) start() Lexer {

	// Prevents start being called twice without calling end between them
	if !lexer.end_called {
		fmt.Println(fmt.Errorf("ERROR: 'func end()' not called before calling 'func start()', try 'myLexer = myLexer.end()'."))
		os.Exit(-1)
	}

	lexer.end_called = false
	lexer.start_called = true

	lexer.continue_lexer = true
	lexer.move = lexer.position
	return lexer
}

/*
 * Lexer method end
 * stops the lexer and checks for success
 * REQUIRED FOR LEXER TO WORK PROPERLY
 */
func (lexer Lexer) end() Lexer {
	// Prevents end being called twice without calling start between them
	if !lexer.start_called {
		fmt.Println(fmt.Errorf("ERROR: 'func start()' not called before calling 'func end()', try 'myLexer = myLexer.start()(...)'."))
		os.Exit(-1)
	}

	if lexer.continue_lexer == false {
		lexer.position = lexer.move
	}

	lexer.start_called = false
	lexer.end_called = true
	lexer.pass_data = []string{}
	return lexer
}

/*
 * lexer method
 * next_string
 * Purpose: checks for the next token and logs previous characters if token is found. If token is not found, stop function chain.
 */
func (lexer Lexer) next_string(get string, log int) Lexer {

	logger := ""
	// checks if the previous operation returned true, If it did, continue the operation
	if lexer.continue_lexer {
		lexer.continue_lexer = false

		/*
		 * time complexity is O(n+x),
		 * where n is the max number of charecters in the line
		 * x is the length of the stop_code list
		 */
		for inc := 0; inc < len(lexer.program); inc++ {

			if lexer.position+1 < len(lexer.program) {

				if lexer.stop_code[lexer.program[lexer.position:lexer.position+1]] == 1 {

					eol_list_item := lexer.program[lexer.position : lexer.position+1]
					// if logging is activated, then add to the pass_data
					if log != LogTypeNone && (get == eol_list_item || get == "EoF") {

						lexer.pass_data = append(lexer.pass_data, format_log(get, logger+get, log)[:]...)
						lexer.continue_lexer = true
						lexer.position += len(eol_list_item)

					}
					return lexer
				}
			}
			// this only executes if EoF is detected
			if lexer.position+len(get) > len(lexer.program) {
				break
			}

			// tests for the token
			// by getting a slice
			// of the program
			// using the current character position
			// plus the tokens length
			if lexer.program[lexer.position:lexer.position+len(get)] == get {
				lexer.continue_lexer = true
				if log != LogTypeNone {
					lexer.pass_data = append(lexer.pass_data, format_log(get, logger+get, log)[:]...)

				}

				//skips past token once found
				//so inf loops are avoided
				lexer.position += len(get)
				break
			}

			// logs data down
			if log != LogTypeNone {
				logger += string(lexer.program[lexer.position])
			}
			// next char
			lexer.position += 1

		}

	}

	return lexer
}

/*
 * lexer method
 * next_string
 * Purpose: checks multiple tokens that might be next and logs previous characters if token is found. If token(s) not found, stope function chain.
 */
func (lexer Lexer) next_list(get []string, log int) Lexer {
	/*
			logger := "";
			// checks if the previous operation returned true, If it did, continue the operation
			if (lexer.continue_lexer) {
					lexer.continue_lexer = false;


					 * time complexity is O(n+x),
					 * where n is the max number of charecters in the line
					 * x is the length of the stop_code list

					for (true) {




							// this only executes if EoF is detected
							if (lexer.position+1>len(lexer.program)) {
								  return lexer;
						}

						// loops through the list of tokens that the lexer/interpereter is searching for
						// TODO: this may need to be changed later to get rid of nested for loops and increase performance
						for current_token:=0;current_token<len(get);current_token++ {
							token := get[current_token];

								// loops through the EoL code list (stop_code) and tries to detect it
							for eoli:=0;eoli<len(lexer.stop_code);eoli++ {

									// End of line list item
									eol_list_item := lexer.stop_code[eoli];

									stop_code_position := lexer.position+len(eol_list_item)


				        	// this only executes if EoL is detected
									if (stop_code_position<len(lexer.program)) {
				               if(lexer.program[lexer.position:stop_code_position] == eol_list_item) {

										   // if logging is activated, then add to the pass_data
											 if (log!=LogTypeNone && (token==eol_list_item || token=="EoF")) {

						  						    lexer.pass_data=append(lexer.pass_data, format_log(token, logger+token, log)[:]...);
												      lexer.continue_lexer=true;
				                      lexer.position+=len(eol_list_item);

								      	};
								       return lexer;
						      	}
								}
							}

							// makes sure token is not larger than program length
							// so errors are avoided
							if (lexer.position+len(token)<len(lexer.program)) {
									// tests for the token
									// by getting a slice
									// of the program
									// using the current character position
									// plus the tokens length
									if (lexer.program[lexer.position:lexer.position+len(token)]==token) {
											lexer.continue_lexer = true;
											if (log!=LogTypeNone) {
										  	lexer.pass_data=append(lexer.pass_data, format_log(token, logger+token, log)[:]...);

											};

											//skips past token once found
											//so inf loops are avoided
											lexer.position+=len(token);
											return lexer;
										}

								  }
						}
						// logs data down
		        if (log!=LogTypeNone) {
			          logger+=string(lexer.program[lexer.position]);
						}
						// next char
						lexer.position+=1;

					}

			}
	*/
	return lexer
}

/*
 * Lexer method
 * unpack
 * Purpose: unpacks a Lexer object and returns attributes
 */
func (lexer Lexer) unpack() (int, string, []string) {
	return lexer.position, lexer.program, lexer.pass_data
}

/*
 * firmat_log
 * Purpose: formats logged data from the lexer
 */
func format_log(get string, log string, log_type int) []string {
	switch log_type {

	case LogTypeNone:
		return []string{""}
	case LogTypeAll:
		return []string{log}
	case LogTypeSeperate:
		return strings.Split(log, " ")
	case LogTypeRemoveWhitespace:
		return []string{strings.TrimSpace(log)}
	case LogTypeSeperate + LogTypeRemoveWhitespace:
		return strings.Split(strings.TrimSpace(log), " ")
	case LogTypeNoSelf:
		return []string{strings.Trim(log, get)}
	case LogTypeSeperate + LogTypeNoSelf:
		log = strings.Trim(log, get)
		return strings.Split(log, " ")
	case LogTypeRemoveWhitespace + LogTypeNoSelf:
		log = strings.Trim(log, get)
		return []string{strings.TrimSpace(log)}
	case LogTypeSeperate + LogTypeRemoveWhitespace + LogTypeNoSelf:
		return strings.Split(strings.Trim(strings.TrimSpace(log), get), " ")
	default:
		return []string{}

	}
}
