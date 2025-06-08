package main

import "fmt"

func UserScan() (string, string, int) {
	var from string
	var to string
	var money int
	fmt.Scan(&from, &to, &money)
	return from, to, money
}
func Convert(from, to string, money int) int {
	return 0
}

func main() {
	const usd_eur = 0.88
	const usd_rub = 77.5
	const eur_rub = usd_rub / usd_eur
}
