package models

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
)

type Cuadrado struct {
	Lado float32
}

func NewCuadrado() (*Cuadrado, error) {
	var ladoStr string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("¿Cuanto mide un lado del cuadrado en cm?").
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

	return &Cuadrado{
		Lado: float32(lado),
	}, nil
}

func (cuadrado Cuadrado) Calcular() {
	perimetro := cuadrado.Lado * 4
	area := cuadrado.Lado * cuadrado.Lado
	fmt.Printf("Los datos del cuadrado son:\nLado: %.1fcm\nPerímetro: %.1fcm\nÁrea: %.1fcm²\n", cuadrado.Lado, perimetro, area)
}
