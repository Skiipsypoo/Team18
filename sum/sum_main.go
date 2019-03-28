package main

import "os"
import "fmt"
import L "./sum"
import "strconv"

func main() {
	x, err := strconv.Atoi(os.Args[1])
	if err != nil{panic(err)}
	y, err := strconv.Atoi(os.Args[2])
	if err != nil{panic(err)}
z := int8(x)
v := int8(y)

fmt.Println(z)
fmt.Println(v)

k := L.SumInt8(z, v)
fmt.Println(k)
}

