package main

import (
	"fmt"
	"math/rand"
)

type Latitude float64
type Longitude float64

func (lat Latitude) IsValid() bool {
	return lat >= -90 && lat <= 90
}

func (long Longitude) IsValid() bool {
	return long >= -180 && long <= 180
}

func main() {
	lat := Latitude(rand.Intn(200)-100)
	long := Longitude(rand.Intn(400)-200)
	fmt.Printf("Is latitude %f valid? %t\n", lat, lat.IsValid())
	fmt.Printf("Is longitude %f valid? %t\n", long, long.IsValid())
}

