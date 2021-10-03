package main

type paintlc_synt struct {
	x                int
	y                int
	z                int
	dec_func         int
	parameters       int
	ret              int
	const_sep        int
	func_mem_usage   int
	func_name        int
	par              int
	variable         int
	variable_value   int
	variable_mem_loc int
	const_dec        int
	const_val        int
	var_name         int
	const_name       int
}

func get_paintlc_defaults() paintlc_synt {
	return paintlc_synt{
		x:              5048,
		y:              5049,
		z:              5050,
		dec_func:       6048,
		parameters:     6049,
		ret:            6050,
		par:            6051,
		const_sep:      6056,
		func_mem_usage: 6057,
		func_name:      6054,

		variable:         7048,
		variable_value:   7049,
		variable_mem_loc: 7050,
		const_dec:        7051,
		const_val:        7056,
		var_name:         7057,
		const_name:       7054,
	}
}
