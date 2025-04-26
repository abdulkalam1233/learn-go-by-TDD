package hello_world

const englishGreetingPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		return englishGreetingPrefix + "World"
	}
	return englishGreetingPrefix + name
}
