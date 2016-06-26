package oop

import (
	"fmt"
)

type Part struct {
	Name        string
	Description string
	NeedsSpare  bool
}

type Parts []Part

func (parts Parts) Spares() (spares Parts) {
	for _, part := range parts {
		if part.NeedsSpare {
			fmt.Println(part, "is appended!")
			spares = append(spares, part)
		}
	}
	fmt.Println(spares)
	return spares
}

type Bicycle struct {
	Size string
	Parts
}

var (
	RoadBikeParts = Parts{
		{"chain", "10-speed", true},
		{"tire_size", "23", true},
		{"tape_color", "red", true},
	}

	MountainBikeParts = Parts{
		{"chain", "10-speed", true},
		{"tire_size", "21", true},
		{"front_shock", "Manitou", false},
		{"rear_shock", "Fox", true},
	}

	RecumbentBikeParts = Parts{
		{"chain", "9-speed", true},
		{"tire_size", "28", true},
		{"flag", "tail and orange", true},
	}
)

func Call_bicycle() {
	roadBike := Bicycle{Size: "L", Parts: RoadBikeParts}
	mountainBike := Bicycle{Size: "L", Parts: MountainBikeParts}
	recumbentBike := Bicycle{Size: "L", Parts: RecumbentBikeParts}

	//roadBike.Spares()
	fmt.Println(roadBike)
	fmt.Println(mountainBike)
	fmt.Println(recumbentBike)
}
