package pkg2

func returnString() string {
	return "abc"
}

func errorShouldBeTheLastType() (error, string) {
	return nil, "abc"
}
