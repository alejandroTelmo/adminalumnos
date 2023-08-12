package alumnos

import "fmt"



type Alumno struct{
	Nombre string
	Apellido string
	FechaNacimiento string
	Pago bool
}

var Alumnos []Alumno

func ListarAlumnos(){
	for i := 0; i < len(Alumnos); i++ {
		fmt.Println("Nombre:", Alumnos[i].Nombre)
		fmt.Println("Apellido:", Alumnos[i].Apellido)
		if Alumnos[i].Pago {
			fmt.Println("Pago: Sí")
		} else {
			fmt.Println("Pago: No")
		}
		fmt.Println()
	}
}

func agregarAlumno(alumno Alumno){
	Alumnos = append(Alumnos, alumno)
}
func ContarAlumnos()int{
	return len(Alumnos)
}


