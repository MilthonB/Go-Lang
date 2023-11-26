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

func Saludador(names []string) (map[string]string, error) {

	mensajes := make(map[string]string)
	for _, name := range names {
		mensaje, err := Saludar(name)
		if err != nil {
			return nil, err
		}

		mensajes[name] = mensaje
	}

	return mensajes, nil

}

func formatoRandom() string {

	formato := []string{
		"Hola, %v. Bienvenido",
		"Encantado de conocerte, %v",
		"Hola, %v un placer",
	}

	return formato[rand.Intn(len(formato))]

}
