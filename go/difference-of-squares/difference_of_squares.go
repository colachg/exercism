package diffsquares


func SquareOfSum(n int) int {
	//for i := 1; i <= number; i++ {
	//	sum += i
	//}

	sum := (n * (n + 1))/2
	return sum * sum
}

func SumOfSquares(n int) int{
	//var sum int
	//for i := 1; i <= number; i++ {
	//	sum += i * i
	//}
	return (n * (n + 1) * (2 * n + 1)) / 6
}

func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}
