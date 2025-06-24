package main

import (
	"errors"
	"math"
	"reflect"
)

type Hitung2d interface {
	luas() float64
	keliling() float64
}

type Hitung3d interface {
	volume() float64
}

type Hitung interface {
	Hitung2d
	Hitung3d
}

type Bola struct {
	d float64
}

//JariJari output is a radius
func (b *Bola) JariJari() (float64, error) {
	if reflect.ValueOf(b.d).Kind() != reflect.Float64{
		return 0, errors.New("wrong input")
	}
	return b.d / 2, nil
}

//Volume output is a volume
func (b *Bola) Volume() (float64, error) {
	r, err := b.JariJari()
	if err != nil{
		return 0, errors.New("wrong input")
	}
	return ((4 / 3) * math.Pi * r * r * r), nil
}

//Luas output is an area
func (b *Bola) Luas() (float64, error) {
	r, err := b.JariJari()
	if err != nil {
		return 0, errors.New("wrong input")
	}
	return (4 * math.Pi * r * r), nil
}

//Keliling output is a perimeter
func (b *Bola) Keliling() (float64, error) {
	r, err := b.JariJari()
	if err != nil {
		return 0, errors.New("wrong input")
	}
	return (2 * math.Pi * r), nil
}
