package main

func addBinary(a string, b string) string {

	if len(b) > len(a) {
		a, b = b, a
	}
	x, carry := []byte(a), uint8('0')
	for index := 1; index <= len(x); index++ {
		if index > len(b) {
			if carry == '1' {
				if x[len(x)-index] == '1' {
					x[len(x)-index] = '0'
				} else {
					x[len(x)-index] = '1'
					carry = '0'
				}
				continue
			}
			carry = '0'
			break
		}

		if x[len(x)-index]+b[len(b)-index] == '1'*2 {
			if carry == '1' {
				continue
			}
			carry = '1'
			x[len(x)-index] = '0'
		} else if x[len(x)-index]+b[len(b)-index] == '0'*2 {
			x[len(x)-index] = carry
			carry = '0'
		} else {
			if carry == '1' {
				x[len(x)-index] = '0'
			} else {
				x[len(x)-index] = '1'
			}
		}

	}
	if carry > '0' {
		x = append([]byte{carry}, x...)
	}
	return string(x)
}
