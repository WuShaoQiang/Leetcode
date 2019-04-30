package reverseBits

func reverseBits(num uint32) uint32 {
	var res uint32 = 0

	for i := 0; i < 32; i++ {
		x := num & 1
		res <<= 1
		if x == 1 {
			res += 1
		}
		num >>= 1
	}
	return res
}

// func reverseBits(num uint32) uint32 {
// 	var result uint32

// 	l := 32
// 	for l > 0 {
// 		result = result << 1
// 		result += num & 1
// 		num = num >> 1
// 		l--
// 	}
// 	return result
// }

// func reverseBits(num uint32) uint32 {
// 	if num == 0 {
// 		return num
// 	}

// 	l := 32
// 	result := make([]uint32, 32)

// 	for num > 0 {
// 		result[l-1] = num % 2
// 		num = num / 2
// 		l--
// 	}

// 	var total uint32
// 	var tmp uint32
// 	tmp = 1
// 	total = 0
// 	for i := 0; i < 32; i++ {
// 		total += result[i] * tmp
// 		tmp = tmp * 2
// 	}
// 	return total
// }
