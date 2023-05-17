package main

import (
	"fmt"
)

func main() {
	var n, m int
	scan, err := fmt.Scan(&n, &m)
	if err != nil {
		return
	}
	fmt.Println(scan)
}
