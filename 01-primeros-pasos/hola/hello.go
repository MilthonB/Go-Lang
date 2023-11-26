package hello

import (
	"errors"
	"fmt"
	"math/rand"
)

func Saludar(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empy name")
	}

	mensaje := fmt.Sprintf(formatoRandom(), name)

	return mensaje, nil

}

func formatoRandom() string {

	formato := []string{
		"Hola, %v. Bienvenido",
		"Encantado de conocerte, %v",
		"Hola, %v un placer",
	}

	return formato[rand.Intn(len(formato))]

}
