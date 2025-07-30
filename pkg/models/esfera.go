package models

import (
	"fmt"
	"math"
	"strconv"

	"github.com/charmbracelet/huh"
)

type Esfera struct {
	Radio float32
}

func NewEsfera() (*Esfera, error) {
	var radioStr string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("¿Cuál es el radio de la esfera en cm?").
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

	return &Esfera{
		Radio: float32(radio),
	}, nil
}

func (e Esfera) Calcular() {
	areaSuperficial := 4 * math.Pi * float64(e.Radio) * float64(e.Radio)
	volumen := (4.0 / 3.0) * math.Pi * math.Pow(float64(e.Radio), 3)
	fmt.Printf(
		"Los datos de la esfera son:\nRadio: %.1fcm\nÁrea superficial: %.1fcm²\nVolumen: %.1fcm³\n",
		e.Radio, areaSuperficial, volumen,
	)
}
