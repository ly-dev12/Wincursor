package main

import (
	"fmt"

	"github.com/wincursor/cursor"
)

func main() {
	var points cursor.POINT
	done := false
	for !done {
		cursor.Getcoords(&points)

		fmt.Println(points.X, points.Y)
	}

}
