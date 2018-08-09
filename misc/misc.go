package misc

// TerOpt provides an emulation of ternary operator.
func TerOpt(cond bool, val1 interface{}, val2 interface{}) interface{} {
	if cond {
		return val1
	}
	return val2
}
