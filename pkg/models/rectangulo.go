package models

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
)

type Rectangulo struct {
	Alto  float32
	Ancho float32
}

func NewRectangulo() (*Rectangulo, error) {
	var altoStr, anchoStr string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("¿Cuál es el alto del rectángulo en cm?").
				Value(&altoStr),
			huh.NewInput().
				Title("¿Cuál es el ancho del rectángulo en cm?").
				Value(&anchoStr),
		),
	)

	if err := form.Run(); err != nil {
		return nil, err
	}

	alto, err := strconv.ParseFloat(altoStr, 32)
	if err != nil {
		return nil, fmt.Errorf("alto inválido: %w", err)
	}
	ancho, err := strconv.ParseFloat(anchoStr, 32)
	if err != nil {
		return nil, fmt.Errorf("ancho inválido: %w", err)
	}

	return &Rectangulo{
		Alto:  float32(alto),
		Ancho: float32(ancho),
	}, nil
}

func (rectangulo Rectangulo) Calcular() {
	perimetro := 2 * (rectangulo.Alto + rectangulo.Ancho)
	area := rectangulo.Alto * rectangulo.Ancho
	fmt.Printf("Los datos del rectángulo son:\nAlto: %.1fcm\nAncho: %.1fcm\nPerímetro: %.1fcm\nÁrea: %.1fcm²\n", rectangulo.Alto, rectangulo.Ancho, perimetro, area)
}
