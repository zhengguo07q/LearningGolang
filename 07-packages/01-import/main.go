package main

import (
	api "01-import/service"
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Max(73.15, 92.46))
	api.HandleReq()
}
