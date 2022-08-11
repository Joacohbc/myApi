package models

import (
	"fmt"
	"myAPI/src/validations"
	"strconv"
	"strings"
	"time"
)

type People struct {
	// Cédula de la persona
	CI int `json:"ci"`

	// Genero de la persona, Masculino, Femenino
	// Gender string `json:"gender"`

	// Nombre de la persona
	Name string `json:"name"`

	// Segundo nombre de la personas
	SecondName string `json:"second_name"`

	// Apellido de la persona
	Surname string `json:"surname"`

	// Segundo apellido de la persona
	SecondSurname string `json:"second_surname"`

	// Fecha de nacimiento de la persona, en string para luego convertirla
	// a time.Time
	Birthdate string `json:"birthdate"`

	// Fecha de nacimiento de la persona que es la que usa le servidor
	// de tipo time.Time
	BirthdateTime time.Time `json:"birthdate_time"`

	// Código del país
	// CountryCode string `json:"country_code"`

	// Numero de la personas
	// PhoneNumber string `json:"phone_number"`

	// Email de la personas
	//Email string `json:"email"`
}

// Retorna en string los datos de la Persona (CI, Name, Surname, Birthdate)
func (p *People) String() string {
	return fmt.Sprintf("%d - %s - %s - %s", p.CI, p.Name, p.Surname, p.Birthdate)
}

// Valida la CI de la persona
func ValidCI(ci int) error {
	if len(strconv.Itoa(ci)) != 8 {
		return fmt.Errorf("el campo de la CI debe tener 8 digitos")
	}
	return nil
}

// Valida la CI de la persona
func (p *People) ValidCI() error {
	return ValidCI(p.CI)
}

// Valida el Name de la persona y le aplico un formato correcto
func (p *People) ValidName() error {

	// El nombre del campo del error
	const fieldName = "el primer nombre"

	// Valido el nombre
	if err := validations.ValidText(fieldName, p.Name, validations.NotEmpty, 1, 50, validations.CanContainsSpace); err != nil {
		return err
	}

	// Valido que le nombre solo tenga letras
	if err := validations.ValidOnlyLetters(fieldName, p.Name, validations.CanContainsSpace); err != nil {
		return err
	}

	return nil
}

func formatTextField(s string) string {
	return validations.FirstUpperInEachWord(strings.TrimSpace(s))
}

// Aplica el formato correcto para el Name de la persona
func (p *People) FormatName() {
	p.Name = formatTextField(p.Name)
}

// Valida el SecondName de la persona
func (p *People) ValidSecondName() error {

	// El nombre del campo del error
	const fieldName = "el segundo nombre"

	// Valido el nombre
	if err := validations.ValidText(fieldName, p.SecondName, validations.CanBeEmpty, 0, 50, validations.CanContainsSpace); err != nil {
		return err
	}

	// Valido que le nombre solo tenga letras
	if err := validations.ValidOnlyLetters(fieldName, p.SecondName, validations.CanContainsSpace); err != nil {
		return err
	}

	return nil
}

// Aplica el formato correcto para el SecondName de la persona
func (p *People) FormatSecondName() {
	p.SecondName = formatTextField(p.SecondName)
}

// Valida el Surname de la persona
func (p *People) ValidSurname() error {

	// El nombre del campo del error
	const fieldName = "el primer apellido"

	// Valido el nombre
	if err := validations.ValidText(fieldName, p.Surname, validations.NotEmpty, 1, 50, validations.CanContainsSpace); err != nil {
		return err
	}

	// Valido que le nombre solo tenga letras
	if err := validations.ValidOnlyLetters(fieldName, p.Surname, validations.CanContainsSpace); err != nil {
		return err
	}

	return nil
}

// Aplica el formato correcto para el Surname de la persona
func (p *People) FormatSurname() {
	p.Surname = formatTextField(p.Surname)
}

// Valida Second Surname de la persona
func (p *People) ValidSecondSurname() error {

	// El nombre del campo del error
	const fieldName = "el segundo apellido"

	// Valido el nombre
	if err := validations.ValidText(fieldName, p.SecondSurname, validations.CanBeEmpty, 0, 50, validations.CanContainsSpace); err != nil {
		return err
	}

	// Valido que le nombre solo tenga letras
	if err := validations.ValidOnlyLetters(fieldName, p.SecondSurname, validations.CanContainsSpace); err != nil {
		return err
	}

	return nil
}

// Aplica el formato correcto para el SecondSurname de la persona
func (p *People) FormatSecondSurname() {
	p.SecondSurname = formatTextField(p.SecondSurname)
}

// Valida Birthdate de la persona y sobrescribe el valor de BirthdateTime
func (p *People) ValidBirthdate() error {

	// Referencia: January 2, 15:04:05, 2006
	timeParsed, err := time.Parse("02/01/2006", p.Birthdate)
	if err != nil {
		return fmt.Errorf("el formato de la fecha de nacimiento debe ser dia/mes/año")
	}
	p.BirthdateTime = timeParsed

	if p.BirthdateTime.After(time.Now()) {
		return fmt.Errorf("la fecha de nacimiento debe ser anterior a la fecha actual")
	}

	return nil
}

// Valida todos los campos de la persona
func (p *People) ValidAll() error {

	if err := p.ValidCI(); err != nil {
		return err
	}

	if err := p.ValidName(); err != nil {
		return err
	}

	if err := p.ValidSecondName(); err != nil {
		return err
	}

	if err := p.ValidSurname(); err != nil {
		return err
	}

	if err := p.ValidSecondSurname(); err != nil {
		return err
	}

	if err := p.ValidBirthdate(); err != nil {
		return err
	}

	return nil
}

// Aplica el formato a todos los campos de la persona
func (p *People) FormatAll() {
	p.FormatName()
	p.FormatSecondName()
	p.FormatSurname()
	p.FormatSecondSurname()
}

// Aplica el formato y valida todos los campos de la persona
func (p *People) FormatAndValidAll() error {
	p.FormatAll()
	return p.ValidAll()
}
