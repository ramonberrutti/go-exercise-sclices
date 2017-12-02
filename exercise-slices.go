package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		arr[y] = make([]uint8, dx)
		for x := 0; x < dy; x++ {
			arr[y][x] = uint8(x^y)
		}
	}
	return arr
}

func main() {
	pic.Show(Pic)
}
