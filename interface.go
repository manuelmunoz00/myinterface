package myinterface

import "fmt"

// Perro struct
type Perro struct {
}

// Pescado struct
type Pescado struct {
}

// Ave struct
type Ave struct {
}

//Los structs anteriores deben implementar el metodo mover para cada uno que es requisito para la interfaz de tipo animal
func (p Perro) mover() string {
	return "Soy un perro y estoy caminando"
}

func (p Pescado) mover() string {
	return "Sour un pescado y estoy nadando"
}

func (p Ave) mover() string {
	return "Soy un ave y estoy volando"
}

//Interfaz de tipo animal
type animal interface {
	mover() string
}

//MoverAnimal Función que recibe un animal
func MoverAnimal(a animal) {
	fmt.Println(a.mover())
}
