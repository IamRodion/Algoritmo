package main

import (
	"github.com/IamRodion/algoritmo/pkg/models"
	"github.com/charmbracelet/huh"
)

func calcularFigura() {
	var figura string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("¿Cuál es la figura a calcular?").
				Options(
					huh.NewOption("Rectángulo", "Rectángulo"),
					huh.NewOption("Cuadrado", "Cuadrado"),
					huh.NewOption("Circulo", "Circulo"),
					huh.NewOption("Triangulo", "Triangulo"),
				).
				Value(&figura),
		),
	)

	if err := form.Run(); err != nil {
		panic("Error al ejecutar el formulario: " + err.Error())
	}

	switch figura {
	case "Rectángulo":
		rect, err := models.NewRectangulo()
		if err != nil {
			panic("Error")
		}
		rect.Calcular()
	case "Cuadrado":
		cuad, err := models.NewCuadrado()
		if err != nil {
			panic("Error")
		}
		cuad.Calcular()
	case "Circulo":
		circ, err := models.NewCirculo()
		if err != nil {
			panic("Error")
		}
		circ.Calcular()
	case "Triangulo":
		tri, err := models.NewTriangulo()
		if err != nil {
			panic("Error")
		}
		tri.Calcular()
	}
}

func calcularSolido() {
	var solido string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("¿Cuál es el solido a calcular?").
				Options(
					huh.NewOption("Cubo", "Cubo"),
					huh.NewOption("Cilindro", "Cilindro"),
					huh.NewOption("Cono", "Cono"),
					huh.NewOption("Esfera", "Esfera"),
				).
				Value(&solido),
		),
	)
	if err := form.Run(); err != nil {
		panic("Error al ejecutar el formulario: " + err.Error())
	}

	switch solido {
	case "Cubo":
		cubo, err := models.NewCubo()
		if err != nil {
			panic("Error")
		}
		cubo.Calcular()
	case "Cilindro":
		cili, err := models.NewCilindro()
		if err != nil {
			panic("Error")
		}
		cili.Calcular()
	case "Cono":
		cono, err := models.NewCono()
		if err != nil {
			panic("Error")
		}
		cono.Calcular()
	case "Esfera":
		esfe, err := models.NewEsfera()
		if err != nil {
			panic("Error")
		}
		esfe.Calcular()

	}

}

func main() {
	var eleccion string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("¿Que quiere calcular?").
				Options(
					huh.NewOption("Figura (2D)", "Figura"),
					huh.NewOption("Solido (3D)", "Solido"),
				).
				Value(&eleccion),
		),
	)
	if err := form.Run(); err != nil {
		panic("Error al ejecutar el formulario: " + err.Error())
	}

	switch eleccion {
	case "Figura":
		calcularFigura()
	case "Solido":
		calcularSolido()
	}
}
