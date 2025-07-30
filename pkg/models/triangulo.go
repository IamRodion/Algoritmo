package models

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
)

type Triangulo struct {
	LadoA  float32
	LadoB  float32
	Base   float32
	Altura float32
}

func NewTriangulo() (*Triangulo, error) {
	var ladoAStr, ladoBStr, baseStr, alturaStr string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("¿Cuál es el lado A del triángulo en cm?").
				Value(&ladoAStr),
			huh.NewInput().
				Title("¿Cuál es el lado B del triángulo en cm?").
				Value(&ladoBStr),
			huh.NewInput().
				Title("¿Cuál es la base del triángulo en cm?").
				Value(&baseStr),
			huh.NewInput().
				Title("¿Cuál es la altura del triángulo en cm?").
				Value(&alturaStr),
		),
	)

	if err := form.Run(); err != nil {
		return nil, err
	}

	ladoA, err := strconv.ParseFloat(ladoAStr, 32)
	if err != nil {
		return nil, fmt.Errorf("lado A inválido: %w", err)
	}
	ladoB, err := strconv.ParseFloat(ladoBStr, 32)
	if err != nil {
		return nil, fmt.Errorf("lado B inválido: %w", err)
	}
	base, err := strconv.ParseFloat(baseStr, 32)
	if err != nil {
		return nil, fmt.Errorf("base inválida: %w", err)
	}
	altura, err := strconv.ParseFloat(alturaStr, 32)
	if err != nil {
		return nil, fmt.Errorf("altura inválida: %w", err)
	}

	return &Triangulo{
		LadoA:  float32(ladoA),
		LadoB:  float32(ladoB),
		Base:   float32(base),
		Altura: float32(altura),
	}, nil
}

func (t Triangulo) Calcular() {
	perimetro := t.LadoA + t.LadoB + t.Base
	area := (t.Base * t.Altura) / 2
	fmt.Printf("Los datos del triángulo son:\nLado A: %.1fcm\nLado B: %.1fcm\nBase: %.1fcm\nAltura: %.1fcm\nPerímetro: %.1fcm\nÁrea: %.1fcm²\n",
		t.LadoA, t.LadoB, t.Base, t.Altura, perimetro, area)
}
