package MultiplyStrings

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	result := make([]int, len(num1)+len(num2))

	for index1 := len(num1) - 1; index1 >= 0; index1-- {
		for index2 := len(num2) - 1; index2 >= 0; index2-- {
			sum := int(num1[index1]-'0') * int(num2[index2]-'0')
			result[index2+index1+1] += sum
			result[index2+index1] += result[index2+index1+1] / 10
			result[index2+index1+1] = result[index2+index1+1] % 10
		}
	}
	if result[0] == 0 {
		result = result[1:]
	}

	res := make([]byte, len(result))
	for i := 0; i < len(result); i++ {
		res[i] = '0' + byte(result[i])
	}
	return string(res)
}
