package main

import "fmt"

func luasLingkaran(nums float64) float64 {
	phi := 3.14
	var search = nums * nums * phi
	return search
}

func kelilingLingkaran(nums float64) float64 {
	phi := 3.14
	var keliling = 2 * phi * nums
	return keliling

}

func main() {
	var nums float64
	fmt.Printf("\n-------Hitung Luas dan Keliling Lingkaran-------\n\nInput Panjang Jari- Jari : ")
	fmt.Scanf("%g", &nums)
	fmt.Printf("\nLuas Lingkaran : %1.f \n", luasLingkaran(nums))
	fmt.Printf("Keliling Lingkaran : %1.f \n\n--------Succes--------\n\n\n", kelilingLingkaran(nums))
}
