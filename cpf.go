package main

import (
	"os"
	"strconv"
)

func strings2ints(s []string) []int {
	tmp := make([]int, len(s))

	for i, v := range s {
		num, _ := strconv.Atoi(v)
		tmp[i] = num
	}
	return tmp
}

var cache map[int]int

func ints2slice(s int) []int {
	numberLen := len(strconv.Itoa(s))
	tmp := make([]int, numberLen)
	for c := 0; c < numberLen; c++ {
		tmp[numberLen-c-1] = s % 10
		s /= 10
	}

	return tmp
}

func sum(s []int) (sum int) {
	sum = 0

	for _, v := range s {
		sum += v
	}
	return
}

func ints2byte(s []int) []byte {
	tmp := make([]byte, len(s)+1)
	for i, v := range s {
		tmp[i] = byte(v + '0')
	}
	return tmp
}

func genCpfs(start int, end int, file string) {
	handle, _ := os.OpenFile(file, os.O_WRONLY, os.ModeAppend)

	// var thisList strings.Builder

	for n := start; n <= end; n++ {
		nums := ints2slice(n)
		value := 0
		for i, n := range nums {
			value += (10 - i) * n
		}

		digit1 := (11 - (value % 11)) % 10

		value2 := value + (sum(nums) + (digit1 * 2))
		digit2 := (11 - (value2 % 11)) % 10

		nums = append(nums, digit1)
		nums = append(nums, digit2)
		handle.Write(ints2byte(nums))
	}
}

func main() {
	// cpfsTotal := int64(10e11)
	// nCpus := runtime.NumCPU()

	// cache = make(map[int]int)
	genCpfs(100000000, 200000000, "cpfs.txt")
	// nums := strings2ints(strings.Split(CPF[0:len(CPF)-2], ""))

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
