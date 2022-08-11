package validations

import "strings"

// Retorna el string con la primera letra mayúscula
func FirstUpper(s string) string {
	// Si solo tiene una letra, retorno esa única letra en mayúscula
	if len(s) == 1 {
		return strings.ToUpper(string(s[0]))
	} else {
		// Si tiene mas de una letra, retorno la primera mayúscula y el resto en minúscula
		return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
	}
}

// Retorna el string con la primera letra mayúscula de cada palabra (usa strings.Fields para separar)
func FirstUpperInEachWord(s string) string {
	// Separo el string en palabras
	words := strings.Fields(s)

	// A cada palabra le pongo la primera letra en mayúscula
	for i, v := range words {
		words[i] = FirstUpper(v)
	}

	// Retorno todas las palabras unidas por un espacio
	return strings.Join(words, " ")
}
