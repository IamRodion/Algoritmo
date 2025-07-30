package models

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
)

type Cubo struct {
	Lado float32
}

func NewCubo() (*Cubo, error) {
	var ladoStr string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("¿Cuánto mide un lado del cubo en cm?").
				Value(&ladoStr),
		),
	)

	if err := form.Run(); err != nil {
		return nil, err
	}

	lado, err := strconv.ParseFloat(ladoStr, 32)
	if err != nil {
		return nil, fmt.Errorf("tamaño inválido: %w", err)
	}

	return &Cubo{
		Lado: float32(lado),
	}, nil
}

func (cubo Cubo) Calcular() {
	areaSuperficial := 6 * cubo.Lado * cubo.Lado
	volumen := cubo.Lado * cubo.Lado * cubo.Lado
	fmt.Printf("Los datos del cubo son:\nLado: %.1fcm\nÁrea superficial: %.1fcm²\nVolumen: %.1fcm³\n", cubo.Lado, areaSuperficial, volumen)
}
