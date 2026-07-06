package main

import "fmt"

func luasLingkaran(nums int) int {
	phi := 22 / 7
	var search = nums * nums * phi
	return search
}

func kelilingLingkaran(nums float64) float64 {
	phi := 3.14
	var keliling = 2 * phi * nums
	return keliling

}

func main() {
	luas := luasLingkaran(7)
	keliling := kelilingLingkaran(7)
	fmt.Println("Luas Lingkaran :", luas)
	fmt.Println("Keliling Lingkaran :", keliling)
}
