package hello

const (
	// Languages
	spanish = "Spanish"
	french = "French"
	sanskrit = "Sanskrit"
	german = "German"

	// Prefixes
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix = "Bonjour, "
	sanskritPrefix = "नमस्कारः, "
	germanPrefix = "begrüßen, "
) 

func Hello(name, language string) string {	
	if name == "" {
		name = "World"
	}

	return greetPrefix(language) + name
}

// In Go, public functions starts with a capital letter and private functions starts with a small letter

func greetPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	case sanskrit:
		prefix = sanskritPrefix
	case german:
		prefix = germanPrefix
	default:
		prefix = englishPrefix
	}
	return
}
