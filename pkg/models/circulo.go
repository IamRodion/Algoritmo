package models

import (
	"fmt"
	"math"
	"strconv"

	"github.com/charmbracelet/huh"
)

type Circulo struct {
	Radio float32
}

func NewCirculo() (*Circulo, error) {
	var radioStr string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("¿Cuál es el Radio del circulo en cm?").
				Value(&radioStr),
		),
	)

	if err := form.Run(); err != nil {
		return nil, err
	}

	radio, err := strconv.ParseFloat(radioStr, 32)
	if err != nil {
		return nil, fmt.Errorf("radio inválido: %w", err)
	}

	return &Circulo{
		Radio: float32(radio),
	}, nil
}

func (circulo Circulo) Calcular() {
	perimetro := 2 * math.Pi * circulo.Radio
	area := math.Pi * (circulo.Radio * circulo.Radio)
	fmt.Printf("Los datos del circulo son:\nRadio: %.1fcm\nPerímetro: %.1fcm\nÁrea: %.1fcm²\n", circulo.Radio, perimetro, area)
}
