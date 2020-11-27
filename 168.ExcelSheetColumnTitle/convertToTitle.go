package ExcelSheetColumnTitle

func convertToTitle(n int) string {
	var result string
	for n != 0 {
		rem := (n - 1) % 26
		n = (n - 1) / 26
		result = string(byte(rem+'A')) + result
	}

	return result
}
