package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Temp struct {
	min      float64
	max      float64
	sum      float64
	quantity int
}

// func onebrc() {

// }

func main() {
	file, err := os.Open("measurements.txt")

	var result = make(map[string]Temp)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine()

		if len(line) > 0 {
			details := strings.Split(string(line), ";")
			if _, ok := result[details[0]]; !ok {
				f, err := strconv.ParseFloat(details[1], 64)

				if err != nil {
					panic(err)
				}

				result[details[0]] = Temp{
					min:      f,
					max:      f,
					sum:      f,
					quantity: 1,
				}
			} else {
				f, err := strconv.ParseFloat(details[1], 64)

				if err != nil {
					panic(err)
				}

				result[details[0]] = Temp{
					min:      min(result[details[0]].min, f),
					max:      max(result[details[0]].max, f),
					sum:      result[details[0]].sum + f,
					quantity: result[details[0]].quantity + 1,
				}
			}
		}
		if err != nil {
			break
		}
	}
	fmt.Println(result)

}
