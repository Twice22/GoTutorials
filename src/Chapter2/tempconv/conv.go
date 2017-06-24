package tempconv

// conversion functions here
func CToF(c Celsius) Fahrenheit { return c * 9 / 5 + 32}
func FToC(f Fahrenheit) Celsius { return (f - 32) * 5 / 9 }

// Note: all other types const func declared in another files
// having the same package name (package tempconv here) are visible
// in all other files of the same package (it's behave as if all the
// fct definition... are in the same file)