package main

import (
	"fmt"
	"math"
)

func calc_flow_rate(bore float64, stroke float64, r_1 float64, r_2 float64, rpm float64) float64 {

	volume := math.Pow(bore*0.5, 2.0) * math.Pi * stroke

	end_rpm := r_2 / r_1 * rpm

	flow_rate := volume * end_rpm * 2 // double this as the piston pump is double acting.

	return flow_rate * 1e-6 * 60.0

}

func main() {
	r := calc_flow_rate(41.0, 38.0, 340.0, 50.0, 1_400.0)

	fmt.Println("flow rate in litres per hour is", r)
}
