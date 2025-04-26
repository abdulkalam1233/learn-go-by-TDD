package hello_world

const (
	french  = "French"
	spanish = "Spanish"

	englishGreeting = "Hello, "
	frenchGreeting  = "Bonjour, "
	spanishGreeting = "Hola, "
)

func Hello(name string, country string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(country) + name

}

func greetingPrefix(country string) (prefix string) {
	switch country {
	case spanish:
		prefix = spanishGreeting
	case french:
		prefix = frenchGreeting
	default:
		prefix = englishGreeting
	}
	return
}
