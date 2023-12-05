package hola

func Hola(name string) string {

	if name == "" {
		return "", error.new("El nombre esta vacios")
	}

	mensaje := "Hola, mi nombre es: " + name

	return mensaje

}
