## Answer to ex7.7

the help message display a °C because it uses the String() method from celsiusFlag
(or more precisely the String() method from Celsius type embedded in celsiusFlag).
Indeed, we have:

`func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }`