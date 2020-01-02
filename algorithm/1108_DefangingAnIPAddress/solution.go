package _1108_DefangingAnIPAddress

func defangIPaddr(address string) string {
	s := ""
	for _, c := range address {
		if c == int32("."[0]) {
			s += "[.]"
		} else {
			s += string(c)
		}
	}
	return s
}
