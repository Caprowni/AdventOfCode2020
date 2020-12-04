package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getTrees(across int, down int){
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	trees := 0
	a := 0
	d := 0 
	for i :=0; i < len(lines)-1; i++ {
		a += across	
		d += down
		if d >= len(lines){
			break
		}
		if lines[d][a%len(lines[0])] == '#'{
			trees++
		}
	}
	fmt.Println(trees)
}

func main() {
	getTrees(1,1)
	getTrees(3,1)
	getTrees(5,1)
	getTrees(7,1)
	getTrees(1,2)
}

