package main

import (
	"fmt"
	"github.com/kr/pretty"
	"math"
)

type Measurement struct {
	value float64 //value of measurement
	units string  //unit of measurement
}

type Rocket struct {
	name            string      //name of the rocket
	mass            Measurement //mass in kg
	area            Measurement //effective area in m ^2
	drag            Measurement //drag coefficient
	max_velocity    Measurement //maximum velocity
	max_altitude    Measurement //maximum velocity
	wind_resistance Measurement //wind resistance
}

type Engine struct {
	//TODO: build table of impulse / thurs for all estes engines
	//http://www2.estesrockets.com/pdf/Estes_Engine_Chart.pdf
	product_number  int
	name            string
	delay           Measurement
	impulse         Measurement // impulse rating of the engine (N-sec)
	thrust          Measurement //thrust rating of the engine (N)
	mass            Measurement // total mass of the engine (kg)
	propellant_mass Measurement // mass of the propellent
	burn_time       Measurement //total burn time of rocket
	max_payload     Measurement
}

//http://www.rocketmime.com/rockets/qref.html
func compute_terminal_conditons(r *Rocket, e *Engine) {
	r.compute_wind_resistance()

	//computer oft-used intermediate terms

	r.compute_max_velocity(e.thrust.value, e.burn_time.value)

	r.compute_max_altitude(e.thrust.value)
}

func (r *Rocket) compute_wind_resistance() {
	rho := 1.2                                                        // 1.2 kg/m^3 of air
	r.wind_resistance.value = 0.5 * rho * r.drag.value * r.area.value //wind resistance factor
}

func (r *Rocket) compute_max_altitude(thrust float64) {
	g := 9.8 // 9.8 m / s^2

	//compute the height of boost is complete
	//yb = [-M / (2*k)]*ln([T - M*g - k*v^2] / [T - M*g])
	yb1 := (-1 * r.mass.value / (2 * r.wind_resistance.value))
	yb21 := thrust - r.mass.value*g - r.wind_resistance.value*math.Pow(r.max_velocity.value, 2)
	yb22 := thrust - r.mass.value*g
	yb := yb1 * math.Log(yb21/yb22)
	fmt.Println("Your motor will empty at an altitude of:", yb, "meters")

	//compute the height once the boosters are off
	//yc = [+M / (2*k)]*ln([M*g + k*v^2] / [M*g])
	yc1 := (r.mass.value / (2 * r.wind_resistance.value))
	yc21 := r.mass.value*g + r.wind_resistance.value*math.Pow(r.max_velocity.value, 2)
	yc22 := r.mass.value * g
	yc := yc1 * math.Log(yc21/yc22)
	fmt.Println("Your post-thrust elevation gain will be:", yc, "meters")

	y_max := yb + yc
	r.max_altitude.value = y_max
}

func (r *Rocket) compute_max_velocity(thrust float64, time float64) {
	//max velocity given a rocket's burn time and thrust
	g := 9.8 // 9.8 m / s^2
	q := math.Sqrt((thrust - r.mass.value*g) / r.wind_resistance.value)
	x := 2 * r.wind_resistance.value * q / r.mass.value
	r.max_velocity.value = q * (1 - math.Exp(-1*x*time)) /
		(1 + math.Exp(-1*x*time))
}

func main() {
	r := Rocket{
		name:         "Audrey Mark 0",
		mass:         Measurement{value: 0.05, units: "kg"},
		area:         Measurement{value: 0.0004, units: "m^2"},
		drag:         Measurement{value: 0.75, units: ""}, //seems reasonable
		max_velocity: Measurement{value: 0, units: "m/s"},
		max_altitude: Measurement{value: 0, units: "m"},
		//https://en.wikipedia.org/wiki/Drag_coefficient

	} //estimated rocket

	e := Engine{
		product_number:  1598,
		name:            "A8-3",
		max_payload:     Measurement{value: 0.085, units: "kg"},
		delay:           Measurement{value: 3, units: "s"},
		impulse:         Measurement{value: 2.5, units: "N-s"},
		thrust:          Measurement{value: 10.7, units: "N"},
		mass:            Measurement{value: 0.0162, units: "kg"},
		propellant_mass: Measurement{value: 0.00312, units: "kg"},
		burn_time:       Measurement{value: 0.5, units: "s"},
	} //example for A8-3

	compute_terminal_conditons(&r, &e)
	fmt.Printf("%# v\n", pretty.Formatter(r))
	fmt.Printf("%# v\n", pretty.Formatter(e))

	fmt.Printf("You're going to reach %v maximum velocity\n", r.max_velocity)
	fmt.Println("Your maximum altitude will be:", r.max_altitude)

}
