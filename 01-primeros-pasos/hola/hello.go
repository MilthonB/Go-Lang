package hello

import (
	"errors"
)

func Saludar(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empy name")
	}

	mensaje := " Hola, mi nombre es: " + name

	return mensaje, nil

}
