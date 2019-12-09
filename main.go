package main

import (
	"fmt"
	"math"
)

func toString(arr []uint8) string {
	var res string

	for _, bt := range arr {
		res = fmt.Sprintf("%s%d", res, bt)
	}

	if res[0] == '1' {
		return res
	}
	return res[1:]
}

func sumBits(a, b, carryOver uint8) (uint8, uint8) {
	var sum uint8
	sum = a & b

	if sum == 1 {
		return 0 | carryOver, sum | carryOver
	} else {
		if a|b == 1 && carryOver == 1 {
			return 0, 1
		} else {
			return (a|b) | carryOver, 0
		}
	}
}

func addBinary(a string, b string) string {
	var carryOver uint8
	aIndex := len(a)-1
	bIndex := len(b)-1
	maxSize := int(math.Max(float64(len(a)), float64(len(b))))+1
	resultArr := make([]uint8, maxSize)
	currArrIndex := maxSize-1

	for aIndex >= 0 || bIndex >= 0 {
		var sum byte
		if aIndex < 0 {
			for bIndex >= 0 {
				sum, carryOver = sumBits(uint8(b[bIndex]-'0'), carryOver, 0)
				resultArr[currArrIndex] = sum
				currArrIndex--
				bIndex--
			}
			break
		}

		if bIndex < 0 {
			for aIndex >= 0 {
				sum, carryOver = sumBits(uint8(a[aIndex]-'0'), carryOver, 0)
				resultArr[currArrIndex] = sum
				currArrIndex--
				aIndex--
			}
			break
		}

		sum, carryOver = sumBits(uint8(a[aIndex]-'0'), uint8(b[bIndex] - '0'), carryOver)
		resultArr[currArrIndex] = sum
		currArrIndex--
		aIndex--
		bIndex--
	}

	resultArr[currArrIndex] = carryOver

	return toString(resultArr)
}

func main() {
	fmt.Printf("sum %s\n", addBinary("1010", "11"))
}
