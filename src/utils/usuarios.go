package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type People struct {
	CI            int       `json:"ci"`
	Name          string    `json:"name"`
	SecondName    string    `json:"second_name"`
	Surname       string    `json:"surname"`
	SecondSurname string    `json:"second_surname"`
	Birthdate     string    `json:"birthdate"`
	BirthdateTime time.Time `json:"birthdate_time"`
}

// Retorna en string los datos de la Persona (CI, Name, Surname, Birthdate)
func (p *People) String() string {
	return fmt.Sprintf("%d - %s - %s - %s", p.CI, p.Name, p.Surname, p.Birthdate)
}

// Valida la CI de la persona
func (p *People) ValidCI() error {
	if len(strconv.Itoa(p.CI)) != 8 {
		return fmt.Errorf("el campo de la CI debe tener 8 digitos")
	}
	return nil
}

// Valida Name y Second Name de la persona
func (p *People) ValidNames() error {
	p.Name = strings.TrimSpace(p.Name)
	if p.Name == "" || len(p.Name) > 50 {
		return fmt.Errorf("el campo del nombre no puede estar vacio ni ser de mas de 50 caracteres")
	}

	p.SecondName = strings.TrimSpace(p.SecondName)
	if len(p.SecondName) > 50 {
		return fmt.Errorf("el campo del segundo nombre no puede tener mas de 50 caracteres")
	}

	return nil
}

// Valida Surname y Second Surname de la persona
func (p *People) ValidSurnames() error {
	p.Surname = strings.TrimSpace(p.Surname)
	if len(p.Surname) == 0 || len(p.Surname) > 50 {
		return fmt.Errorf("el campo del apellido no puede estar vacio ni ser de mas de 50 caracteres")
	}

	p.SecondSurname = strings.TrimSpace(p.SecondSurname)
	if len(p.SecondSurname) > 50 {
		return fmt.Errorf("el campo del segundo apellido no puede estar vacio ni ser de mas de 50 caracteres")
	}

	return nil
}

// Valida Birthdate de la persona
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

	if err := p.ValidNames(); err != nil {
		return err
	}

	if err := p.ValidSurnames(); err != nil {
		return err
	}

	if err := p.ValidBirthdate(); err != nil {
		return err
	}

	return nil
}
