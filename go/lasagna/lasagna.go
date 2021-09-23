package lasagna

const  OvenTime = 40

func RemainingOvenTime(minutes int) int {
	return 40 - minutes
}

func PreparationTime(minutes int) int {
	return 2 * minutes
}

func ElapsedTime(minutes1, minutes2 int) int {
	return PreparationTime(minutes1) + minutes2
}