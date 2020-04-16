// Task1 project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Please, enter numbers separated by spaces num")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	s := strings.Fields(input)
	fmt.Println("Before: ", s)
	a := specialSort(converToInt(s))
	fmt.Println("After: ", a)
}

func converToInt(ss []string) []int {
	var is []int
	for i := len(ss) - 1; i >= 0; i-- {
		r, err := strconv.Atoi(string(ss[i]))
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error, please input valid numbers")
			os.Exit(0)
		} else {
			is = append(is, r)
		}
	}
	return is

}

func specialSort(si []int) []int {

	for i := len(si) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if myCompare(si[j], si[j+1]) {
				si[j], si[j+1] = si[j+1], si[j]
			}
		}
	}
	return si
}

func myCompare(first int, second int) bool {
	var resAns bool
	resAns = false
	res := []byte(strconv.Itoa(first))
	res1 := []byte(strconv.Itoa(second))
	resFS := append(res, res1...)
	resSF := append(res1, res...)
	iSF, _ := strconv.Atoi(string(resSF))
	iFS, _ := strconv.Atoi(string(resFS))
	if iSF > iFS {
		resAns = true
	}

	_ = resAns

	return resAns
}
