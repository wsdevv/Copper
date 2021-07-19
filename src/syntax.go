package main

//initializes the syntax rules
type syntax struct {
	end_line                    []string
	create_variable             []string
	set_variable                []string
	create_pointer              []string
	get_pointer                 []string
	create_function             []string
	parameters_start            []string
	parameters_end              []string
	chunk_start                 []string
	chunk_end                   []string
	string_start                []string
	string_end                  []string
	array_start                 []string
	array_end                   []string
	create_import               []string
	change_rules                []string
	one_character_no_space_rule bool
}

func defaultSyntax() syntax {
	return syntax{
		end_line:                    []string{"\n", ";"},
		create_variable:             []string{"!", "var"},
		create_pointer:              []string{"@"},
		get_pointer:                 []string{"^"},
		create_function:             []string{"$", "func"},
		parameters_start:            []string{"("},
		parameters_end:              []string{")"},
		chunk_start:                 []string{"{", "start"},
		chunk_end:                   []string{"}", "end"},
		string_start:                []string{"\"", "'"},
		string_end:                  []string{"\"", "'"},
		array_start:                 []string{"["},
		array_end:                   []string{"]"},
		create_import:               []string{"&", "import"},
		change_rules:                []string{"#"},
		set_variable:                []string{"="},
		one_character_no_space_rule: true,
	}
}
