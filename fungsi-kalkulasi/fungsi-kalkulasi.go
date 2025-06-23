package main

import(
	"fmt"
	"math"
)

type hitung2d interface{
	luas() float64
	keliling() float64
}

type hitung3d interface{
	volume() float64
}

type hitung interface{
	hitung2d
	hitung3d
}

type bola struct{
	d float64
}

func (b *bola) jariJari() float64{
	return b.d / 2
}

func (b *bola) volume() float64 {
	return ((4/3) * math.Pi * math.Pow(b.jariJari(), 3))
}

func (b *bola) luas() float64 {
	return (4 * math.Pi * math.Pow(b.jariJari(), 2))
}

func (b *bola) keliling() float64 {
	return (2 * math.Pi * b.jariJari())
}

func main(){
	var bangunRuang hitung = &bola{5}

	fmt.Printf("Luas: %.2f cm^2\n", bangunRuang.luas())
	fmt.Printf("Keliling: %.2f cm\n", bangunRuang.keliling())
	fmt.Printf("Volume: %.2f cm^3\n", bangunRuang.volume())
	fmt.Printf("Jari-jari: %.2f cm\n", bangunRuang.(*bola).jariJari())
}