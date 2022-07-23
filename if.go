package functional

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
func Unless(cond bool, action func() error) error {
	return If(!cond, action)
}
