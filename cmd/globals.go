package main

type Sampah struct {
	Jenis           string
	MetodeDaurUlang string
	Jumlah          int
	Lokasi          string
	Status          string
}

var (
	DataSampah []Sampah
	version    string = "v0.4"
)
