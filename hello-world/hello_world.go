package hello

const testVersion = 2

// Given a string, will return "Hello, [string]!"
func HelloWorld(s string) string {
	if s != "" {
		return "Hello, " + s + "!"
	}

	return "Hello, World!"
}
