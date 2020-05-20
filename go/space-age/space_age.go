package space


type Planet string

func Age(seconds float64, planet Planet)  float64 {
	earthYear := 31557600.0
	switch planet {
	case "Mercury":
		period := earthYear * 0.2408467
		return seconds / period
	case "Venus":
		period := earthYear * 0.61519726
		return seconds / period
	case "Mars":
		period := earthYear * 1.8808158
		return seconds / period
	case "Jupiter":
		period := earthYear * 11.862615
		return seconds / period
	case "Saturn":
		period := earthYear * 29.447498
		return seconds / period
	case "Uranus":
		period := earthYear * 84.016846
		return seconds / period
	case "Neptune":
		period := earthYear * 164.79132
		return seconds / period
	default:
		return seconds / earthYear
	}
}
