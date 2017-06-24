package tempconv

import "fmt"

// defined two types (Celsius and Fahrenheit)
// In a sense, Celsius and Fahrenheit are synonyms of float64 type
// yet we cannot perfom operations between Celsius and Fahrenheit (like c + f)
type Celsius float64
type Fahrenheit float64

// define constantes of type Celsius
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

// define conversions function from Celsius to Fahrenheit and Fahrenheit to Celsius
// Fahrenheit() and Celsius() are not fct calls. These are conversions to return the
// appropriate types
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius( (f - 32) * 5 / 9) }

// overwrite string method for Celsius type
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }