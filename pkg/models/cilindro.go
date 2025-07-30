package models

import (
	"fmt"
	"math"
	"strconv"

	"github.com/charmbracelet/huh"
)

type Cilindro struct {
	Radio  float32
	Altura float32
}

func NewCilindro() (*Cilindro, error) {
	var radioStr, alturaStr string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("¿Cuál es el radio de la base del cilindro en cm?").
				Value(&radioStr),
			huh.NewInput().
				Title("¿Cuál es la altura del cilindro en cm?").
				Value(&alturaStr),
		),
	)

	if err := form.Run(); err != nil {
		return nil, err
	}

	radio, err := strconv.ParseFloat(radioStr, 32)
	if err != nil {
		return nil, fmt.Errorf("radio inválido: %w", err)
	}
	altura, err := strconv.ParseFloat(alturaStr, 32)
	if err != nil {
		return nil, fmt.Errorf("altura inválida: %w", err)
	}

	return &Cilindro{
		Radio:  float32(radio),
		Altura: float32(altura),
	}, nil
}

func (c Cilindro) Calcular() {
	areaBase := math.Pi * float64(c.Radio) * float64(c.Radio)
	areaLateral := 2 * math.Pi * float64(c.Radio) * float64(c.Altura)
	areaTotal := 2*areaBase + areaLateral
	volumen := areaBase * float64(c.Altura)

	fmt.Printf(
		"Los datos del cilindro son:\nRadio: %.1fcm\nAltura: %.1fcm\nÁrea superficial: %.1fcm²\nVolumen: %.1fcm³\n",
		c.Radio, c.Altura, areaTotal, volumen,
	)
}
