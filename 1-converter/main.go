package main

import (
	"errors"
	"fmt"
	"strconv"
)

const usd_eur = 0.88
const usd_rub = 78.48
const eur_rub = usd_rub / usd_eur

type converter = map[string]map[string]float64

func UserScan() (string, string, float64) {
	var from string
	var to string
	var money string
	var money_float64 float64
	fmt.Print("Введите переводимую валюту(USD,EUR,RU) ")
	fmt.Scan(&from)
	err := CurrencyScan(from)
	for err != nil {
		fmt.Println(err)
		fmt.Scan(&from)
		err = CurrencyScan(from)
	}
	fmt.Println("Введите количество денег ")
	for {
		fmt.Scan(&money)
		money1, err := strconv.ParseFloat(money, 64)
		if err == nil && money1 > 0 {
			break
		}
		fmt.Println("Значение меньше 0 или не является числом, введите значение ещё раз")
	}
	money_float64, _ = strconv.ParseFloat(money, 64)
	fmt.Print("Введите целевую валюту")
	switch from {
	case "USD":
		fmt.Println("(RU,EUR) ")
	case "EUR":
		fmt.Println("(RU,USD) ")
	case "RU":
		fmt.Println("(USD,EUR) ")
	}
	fmt.Scan(&to)
	err = CurrencyScan(to)
	for err != nil || to == from {
		fmt.Println("Валюта введена неверно, введите снова")
		fmt.Scan(&to)
		err = CurrencyScan(to)
	}
	return from, to, money_float64
}
func CurrencyScan(value string) error {
	if value != "USD" && value != "EUR" && value != "RU" {
		return errors.New("Валюта введена неверно, введите снова")
	}
	return nil
}

func Converter() {
	var from, to string
	var money float64
	from, to, money = UserScan()
	convert := converter{}
	Convert(from, money, &convert)
	fmt.Println("Переведённые деньги", convert[from][to])
}

func Convert(from string, money float64, convert *converter) {
	switch from {
	case "USD":
		(*convert)["USD"] = map[string]float64{"EUR": money * usd_eur, "RU": money * usd_rub}
	case "RU":
		(*convert)["RU"] = map[string]float64{"EUR": money / usd_rub * usd_eur, "USD": money / usd_rub}
	default:
		(*convert)["EUR"] = map[string]float64{"RU": money * eur_rub, "USD": money / usd_eur}
	}
}

func main() {
	Converter()
}
