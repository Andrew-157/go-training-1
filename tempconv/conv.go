package tempconv

// CToF converts Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - kelvinToCelsiusConversionFactor) }

// CToK converts Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + kelvinToCelsiusConversionFactor) }

// KToF converts Kelvin to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

// FToK converts Fahrenheit to Kelvin.
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }
