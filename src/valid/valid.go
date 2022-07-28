package valid

import (
	"fmt"
	"strings"
	"unicode"
)

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

// Representa si un campo puede ser o no vació (para uso de funciones)
type Emptiness bool

const (
	NotEmpty   Emptiness = false
	CanBeEmpty Emptiness = true
)

// Representa si un campo puede o no contener espacios (para uso de funciones)
type Spaces bool

const (
	NotSpace         Spaces = false
	CanContainsSpace Spaces = true
)

// Valida el valor del puntero string:
/*
	- Aplica TrimSpaces (borra espacio al principio y al final)
	- Todas las letras en mayúscula
	- Controla largo de string (incluyendo máximo y mínimos como largos validos)
*/
// Retornos:
/*
	Esta vació (si @canBeEmpty es false) - @fieldName no puede estar vació
	Contiene espacios (si @containsSpaces es false) - @fieldName no puede contener espacios
	Largo mínimo no cumplido - @fieldName debe tener como mínimo de @minLen caracteres
	Largo máximo no cumplido - @fieldName debe tener como máximo de @maLen caracteres
*/
func ValidText(fieldName string, value string, canBeEmpty Emptiness, minLen int, maxLen int, containsSpaces Spaces) (string, error) {

	// Borro espacios iniciales y finales
	value = strings.TrimSpace(value)

	// Pongo todas en mayúscula
	value = FirstUpperInEachWord(value)

	// Si no puede contener espacios, valido que no contenga espacios
	if !containsSpaces {
		// La función Fields divide por espacios, entonces si el
		// array tiene mas de 1 de largo, significa que hay mas 2 palabras (por tanto un espacio)
		if len(strings.Fields(value)) > 1 {
			return "", fmt.Errorf("%s no puede contener espacios", fieldName)
		}
	}

	// Si no puede ser vacía, valido que no sea vacía
	if !canBeEmpty {
		if len(value) == 0 {
			return "", fmt.Errorf("%s no puede estar vacio", fieldName)
		}
	}

	// Valido la cantidad caracteres mínimos
	if len(value) < minLen {
		return "", fmt.Errorf("%s debe tener como mínimo de %d caracteres", fieldName, minLen)
	}

	// Valido la cantidad caracteres máximos
	if len(value) > maxLen {
		return "", fmt.Errorf("%s debe tener como maximo de %d caracteres", fieldName, maxLen)
	}

	return value, nil
}

// Valida que todos los caracteres del string sea letras o espacios
// Error: @fieldName debe contener unicamente letras, no puede contener: (caracter erróneo)
func ValidOnlyLetters(fieldName string, value string, containsSpaces Spaces) error {

	if containsSpaces {
		for _, c := range value {
			// Si el caracter no es una letra no un espacio
			if !unicode.IsLetter(c) && !unicode.IsSpace(c) {
				return fmt.Errorf("%s debe contener unicamente letras o espacios, no puede contener: \"%s\"", fieldName, string(c))
			}
		}

	} else {
		for _, c := range value {
			// Si el caracter no es una letra
			if !unicode.IsLetter(c) {
				return fmt.Errorf("%s debe contener unicamente letras, no puede contener: \"%s\"", fieldName, string(c))
			}
		}
	}

	return nil
}

// Valida que todos los caracteres del string sean digitos o espacios
// Error: @fieldName debe contener unicamente digitos, no puede contener: (caracter erróneo)
func ValidOnlyDigits(fieldName string, value string, containsSpaces Spaces) error {

	if containsSpaces {
		for _, c := range value {
			// Si el caracter no es un dígito ni un espacio
			if !unicode.IsLetter(c) && !unicode.IsSpace(c) {
				return fmt.Errorf("%s debe contener unicamente numeros o espacios, no puede contener: \"%s\"", fieldName, string(c))
			}
		}

	} else {
		for _, c := range value {
			// Si el caracter no es un dígito
			if !unicode.IsLetter(c) {
				return fmt.Errorf("%s debe contener unicamente numeros, no puede contener: \"%s\"", fieldName, string(c))
			}
		}
	}

	return nil
}

// Valida que todos los caracteres del string sean letras, digitos o espacios
// Error: @fieldName debe contener unicamente letras o digitos, no puede contener: (caracter erróneo)
func ValidLettersAndDigits(fieldName string, value string, containsSpaces Spaces) error {

	if containsSpaces {
		for _, c := range value {
			// Si el caracter no es un dígito ni una letra ni un espacio
			if !unicode.IsDigit(c) && !unicode.IsLetter(c) && !unicode.IsSpace(c) {
				return fmt.Errorf("%s debe contener unicamente letras o numeros, no puede contener: %s", fieldName, string(c))
			}
		}
	} else {
		for _, c := range value {
			// Si el caracter no es un dígito ni una letra ni un espacio
			if !unicode.IsDigit(c) && !unicode.IsLetter(c) {
				return fmt.Errorf("%s debe contener unicamente letras o numeros, no puede contener: %s", fieldName, string(c))
			}
		}
	}

	return nil
}
