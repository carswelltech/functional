package functional

/*
	If accepts a boolean argument and one or two function arguments.  If the bool argument is true, then it evaluates the first function, if the bool argument is false AND two function arguments were supplied, then it calls the second function argument.
*/
func If(cond bool, action ...func() error) error {
	switch {
	case len(action) == 0:
		return nil
	case cond:
		return action[0]()
	case len(action) > 1:
		return action[1]()
	default:
		return nil
	}
}

/*
	Unless accepts a boolean argument and a functional argument, and executes the function argument only if the boolean is false.
*/
func Unless(cond bool, action func() error) error {
	return If(!cond, action)
}
