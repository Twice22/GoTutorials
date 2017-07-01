package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0/9.0) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }


/* celsiusFlag need to satisfy flag.Value interface, whose declaration is:

type Value interface {
	String() string
	Set(string) error
}

*/

// *celsiusFlag satisfies the flag.Value interface
type celsiusFlag struct { Celsius }

// to satisfy the flag.Value interface we only need to define a Set method
// as celsiusFlag already embedded a Celsius type that have a String method!!
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		// need to convert to Celsius before assigning the value
		f.Celsius = Celsius(value)
		// need to return an error (here we set f.Celsius corectly so return nil)
		return nil
	case "F", "°F":
		// need to convert to Fahrenheit before being able to use it in CToF
		f.Celsius = FToC(Fahrenheit(value))
		// need to return an error (here we set f.Celsius correctly so return nil)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// the flag argument must have a quantity and a unit, e.g., "100C".

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}

	// func Var(value Value, name string, usage string)
	// the call to Var adds the flag to the 
	// applications's set of command-line flags (flag.CommandLine)
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
