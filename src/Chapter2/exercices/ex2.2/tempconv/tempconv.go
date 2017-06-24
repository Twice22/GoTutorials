// Package tempconv performs Celsius, Fahrenheit and Kelvin conversions
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

type Feet float64
type Meter float64

type Kg float64
type Pound float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC 	  Celsius = 0
	BoilingC 	  Celsius = 100
)

// String methods
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (f Feet) String() string { return fmt.Sprintf("%g f", f) }
func (m Meter) String() string { return fmt.Sprintf("%g m", m) }

func (k Kg) String() string { return fmt.Sprintf("%g kg", k) }
func (p Pound) String() string { return fmt.Sprintf("%g lb", p) }


// conversion functions here
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func FToM(f Feet) Meter { return Meter(0.3048 * f) }
func MToF(m Meter) Feet { return Feet(m / 0.3048) }

func PToK(p Pound) Kg { return Kg(0.453592 * p) }
func KToP(k Kg) Pound { return Pound(k / 0.453592) }