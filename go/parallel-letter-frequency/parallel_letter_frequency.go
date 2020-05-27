package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(input []string) (output FreqMap){
	comm := make(chan FreqMap)
	defer close(comm)

	// Submit the jobs
	for _, s := range input {
		go func(s string) {
			comm <- Frequency(s)
		}(s)
	}

	// Start with first map
	output = <-comm

	// Reduce subsequenct maps into this
	for i := 1; i < len(input); i++ {
		for k, v := range <-comm {
			output[k] += v
		}
	}

	return output

}
