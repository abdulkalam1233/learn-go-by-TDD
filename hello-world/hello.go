package hello_world

const englishGreeting = "Hello, "
const spanishGreeting = "Hola, "
const spanishCountry = "Spanish"

func Hello(name string, country string) string {
	if name == "" {
		return englishGreeting + "World"
	}
	if country == spanishCountry {
		return spanishGreeting + name
	}
	return englishGreeting + name
}
