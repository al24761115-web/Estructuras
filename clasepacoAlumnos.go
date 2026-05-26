package main

import "fmt"

type Subject struct {
	Name  string
	Grade float64
}

type Student struct {
	Name     string
	Subjects map[string]Subject
}

func (s *Student) AddSubject(name string, grade float64) {
	if s.Subjects == nil {
		s.Subjects = make(map[string]Subject)
	}
	s.Subjects[name] = Subject{
		Name:  name,
		Grade: grade,
	}
}

func (s Student) Average() float64 {
	total := 0.0
	for _, sub := range s.Subjects {
		total += sub.Grade
	}
	if len(s.Subjects) == 0 {
		return 0
	} else {
		result := total / float64(len(s.Subjects))
		return result
	}
}

func (s Student) IsPassing() bool {
	average := s.Average()
	if average >= 70 {
		return true
	}
	return false
}

func (s Student) FindSubject(name string) {
	found := false
	for _, sub := range s.Subjects {
		if sub.Name == name {
			fmt.Println("Clase encontrada", sub.Name, "Calificacion:", sub.Grade)
			found = true
		}
	}
	if found == false {
		fmt.Println("Error la clase ", name, "no existe")
	}
}

func (s *Student) RemoveSubject(name string) {
	found := false
	for _, sub := range s.Subjects {
		if sub.Name == name {
			found = true
		}
	}
	if found == true {
		delete(s.Subjects, name)
		fmt.Println("Clase eliminada:", name)
	} else {
		fmt.Println("Error la clase ", name, "no existe")
	}
}

func (s *Student) UpdateGrade(name string, grade float64) {
	found := false
	for _, sub := range s.Subjects {
		if sub.Name == name {
			found = true
		}
	}
	if found == true {
		s.Subjects[name] = Subject{Name: name, Grade: grade}
		fmt.Println("Calificacion actualizada")
	} else {
		fmt.Println("Error, la clase ", name, "no existe")
	}
}

func (s Student) TopSubject() {
	topGrade := 0.0
	topName := ""
	for _, sub := range s.Subjects {
		if sub.Grade > topGrade {
			topGrade = sub.Grade
			topName = sub.Name
		}
	}
	fmt.Println("Clase con calificacion mas alta:", topName, "tienes", topGrade)
}

func (s Student) Print() {
	fmt.Println("Alumno:", s.Name)
	for _, sub := range s.Subjects {
		fmt.Println("-", sub.Name, ":", sub.Grade)
	}
	fmt.Println("Tu promedio es:", s.Average())
	if s.IsPassing() {
		fmt.Println("Estas Aprobado")
	} else {
		fmt.Println("Estas Reprobado")
	}
}

func main() {
	student := Student{Name: "Carlos"}
	student.AddSubject("Matematicas", 90)
	student.AddSubject("Programacion", 95)
	student.AddSubject("Fisica", 80)

	student.Print()

	fmt.Print("Buscar clase: ")
	var buscar string
	fmt.Scan(&buscar)
	student.FindSubject(buscar)

	fmt.Print("Eliminar clase: ")
	var eliminar string
	fmt.Scan(&eliminar)
	student.RemoveSubject(eliminar)
	student.Print()

	fmt.Print("Actualizar calificacion, escribe el nombre de la clase: ")
	var actualizar string
	fmt.Scan(&actualizar)
	fmt.Print("Nueva calificacion: ")
	var nuevaCalificacion float64
	fmt.Scan(&nuevaCalificacion)
	student.UpdateGrade(actualizar, nuevaCalificacion)
	student.Print()

	fmt.Print("Quieres ver la clase con mayor calificacion?: ")
	var respuesta string
	fmt.Scan(&respuesta)
	if respuesta == "si" {
		student.TopSubject()
	} else {
		fmt.Println(" Terminamos ")
	}
}
