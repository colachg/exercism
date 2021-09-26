package cars

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	const rate = 221
	return float64(rate * speed) * successRate(speed)
}
// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

// successRate is used to calculate the ratio of an item being created without
// error for a given speed
func successRate(speed int) float64 {
	switch {
	case speed <=4 && speed >=1:
		return float64(1)
	case speed <=8 && speed >= 5:
		return 0.9
	case speed <=10 && speed >= 9:
		return 0.77
	}
	return float64(0)
}
