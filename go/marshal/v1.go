package main

import (
	"encoding/json"
	"fmt"
)

/*
we now have three types of wheels. The wheels have some common attribute and special attributes for their own.
the attributes of wheel:
	- common attributes:
		id, radius, type
	- special attributes for SolidWheel:
		material
	- special attributes for SpokeWheel:
		spoke_count, inner_radius
	- special attributes for TireWheel:
		file
*/

// in this implementation, I used the following structure to contain every possible attribute for different wheels.
type Wheel struct {
	Id          string `json:"id"`
	Radius      int    `json:"radius"`
	Type        string `json:"type"`
	Material    string `json:"material"`
	SpokeCount  int    `json:"spoken_count"`
	InnerRadius int    `json:"inner_radius"`
	Fill        string `json:"fill_type"`
}

func main() {
	input := `[{"id":"1234","radius":17,"type":"solid","material":"wood"},{"id":"5432","radius":20,"type":"spoked","spoken_count":30,"inner_radius":19},{"id":"7898","raidus":20,"type":"tire","inner_radius":17,"fill_type":"inflated"}]`

	var wheels []Wheel

	err := json.Unmarshal([]byte(input), &wheels)

	if err == nil {
		for _, w := range wheels {
			fmt.Printf("type: %s, ", w.Type)
			switch w.Type {
			case "solid":
				fmt.Printf("material: %s\n", w.Material)
				break
			case "spoked":
				fmt.Printf("spoken_count: %d, inner_radius: %d\n", w.SpokeCount, w.InnerRadius)
				break
			case "tire":
				fmt.Printf("fill_type: %s\n", w.Fill)
				break
			}
		}
	}
}
