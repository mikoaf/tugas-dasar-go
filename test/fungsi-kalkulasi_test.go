package main

import (
	"math"
	"testing"
)

//TestJariJari for test JariJari function
func TestJariJari(t *testing.T) {
	b := &Bola{d: 10}
	r, err := b.JariJari()
	if err != nil {
		t.Fatalf("JariJari() error: %v", err)
	}

	want := 5.0
	if r != want {
		t.Errorf("JariJari() = %v; want match for %v", r, want)
	}
}

//TestVolume for test Volume function
func TestVolume(t *testing.T) {
	b := &Bola{d: 2}
	v, err := b.Volume()
	if err != nil {
		t.Fatalf("Volume() error: %v", err)
	}

	want := 3.14
	if math.Abs(v-want) > 0.01 {
		t.Errorf("Volume() = %v; want match for %v", v, want)
	}
}

//TestKeliling for test Keliling function
func TestKeliling(t *testing.T) {
	b := &Bola{d: 5}
	v, err := b.Keliling()
	if err != nil {
		t.Fatalf("Volume() error: %v", err)
	}

	want := 15.71
	if math.Abs(v-want) > 0.01 {
		t.Errorf("Keliling() = %v; want match for %v", v, want)
	}
}

//TestLuas for test Luas function
func TestLuas(t *testing.T) {
	b := &Bola{d: 3}
	v, err := b.Luas()
	if err != nil {
		t.Fatalf("Luas() error: %v", err)
	}

	want := 28.27
	if math.Abs(v-want) > 0.01 {
		t.Errorf("Luas() = %v; want match for %v", v, want)
	}
}
