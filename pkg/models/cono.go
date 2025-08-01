package models

import (
	"fmt"
	"math"
	"strconv"

	"github.com/charmbracelet/huh"
)

type Cono struct {
	Radio  float32
	Altura float32
}

func NewCono() (*Cono, error) {
	var radioStr, alturaStr string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("¿Cuál es el radio de la base del cono en cm?").
				Value(&radioStr),
			huh.NewInput().
				Title("¿Cuál es la altura del cono en cm?").
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

	return &Cono{
		Radio:  float32(radio),
		Altura: float32(altura),
	}, nil
}

func (c Cono) Calcular() {
	generatriz := math.Sqrt(float64(c.Radio)*float64(c.Radio) + float64(c.Altura)*float64(c.Altura))
	areaBase := math.Pi * float64(c.Radio) * float64(c.Radio)
	areaLateral := math.Pi * float64(c.Radio) * generatriz
	areaTotal := areaBase + areaLateral
	volumen := (areaBase * float64(c.Altura)) / 3

	fmt.Printf(
		"Los datos del cono son:\nRadio: %.1fcm\nAltura: %.1fcm\nGeneratriz: %.1fcm\nÁrea superficial: %.1fcm²\nVolumen: %.1fcm³\n",
		c.Radio, c.Altura, generatriz, areaTotal, volumen,
	)
}
