package main

import (
	"fmt"
	"strconv"
)

const CPF = "48650757850"

func strings2ints(s []string) []int {
	tmp := make([]int, len(s))

	for i, v := range s {
		num, _ := strconv.Atoi(v)
		tmp[i] = num
	}
	return tmp
}

func ints2slice(s int) []int {
	numberLen := len(strconv.Itoa(s))
	tmp := make([]int, numberLen)
	for c := 0; c < numberLen; c++ {
		tmp[numberLen-c-1] = s % 10
		s /= 10
	}

	return tmp
}

func genCpfs(start int, end int) {
	for n := start; n <= end; n++ {
		fmt.Println(n, ints2slice(n))

		for i

	}
}

func main() {
	genCpfs(100000000, 200000000)
	// nums := strings2ints(strings.Split(CPF[0:len(CPF)-2], ""))

	// value := 0
	// for i, n := range nums {
	// 	mult := (10 - i)
	// 	value += mult * n
	// }

	// digit1 := 11 - (value % 11)
	// digit2 := 0
	// if digit1 < 10 {
	// 	nums = append(nums, digit1)
	// 	value = 0
	// 	for i, n := range nums {
	// 		mult := (11 - i)
	// 		value += mult * n
	// 	}
	// 	digit2 = 11 - (value % 11)
	// 	if digit2 >= 10 {
	// 		digit2 = 0
	// 	}
	// }
	// nums = append(nums, digit2)

	// fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), ""), "[]"))

}
