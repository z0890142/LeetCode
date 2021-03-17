package RomantoInteger

func romanToInt(s string) int {
	rtoi := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	prev := 0
	for index := len(s) - 1; index >= 0; index-- {
		curr := rtoi[s[index]]
		if curr < prev {
			result -= curr
		} else {
			result += curr
		}
		prev = curr

	}

	return result
}

// type RomanNumerial int

// const (
// 	I RomanNumerial = 1
// 	V RomanNumerial = 5
// 	X RomanNumerial = 10
// 	L RomanNumerial = 50
// 	C RomanNumerial = 100
// 	D RomanNumerial = 500
// 	M RomanNumerial = 1000
// )

// func (r RomanNumerial) Int() int {
// 	return int(r)
// }

// var romanNumerials = map[RomanNumerial]string{
// 	I: "I",
// 	V: "V",
// 	X: "X",
// 	L: "L",
// 	C: "C",
// 	D: "D",
// 	M: "M",
// }

// var romanNumerialLookup = map[string]RomanNumerial{
// 	"I": I,
// 	"V": V,
// 	"X": X,
// 	"L": L,
// 	"C": C,
// 	"D": D,
// 	"M": M,
// }

// var orderedNumerials = []RomanNumerial{I, V, X, L, C, D, M}

// func romanToInt(s string) int {
// 	var value int

// 	length := len(s)
// 	priorNumerial := RomanNumerial(0)
// 	for i := length - 1; i >= 0; i-- {
// 		// Retreive integer value of roman numerial
// 		numerial := romanNumerialLookup[string(s[i])]

// 		// If current numerial is smaller than prior, subtract the value
// 		if numerial.Int() < priorNumerial.Int() {
// 			value -= numerial.Int()
// 		} else {
// 			value += numerial.Int()
// 		}
// 		priorNumerial = numerial
// 	}

// 	return value
// }
