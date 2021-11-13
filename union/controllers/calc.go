package controllers

import (
	"fmt"
	"strings"
)

func ValorMontadores(kit string) float64 {
	return (toFloat(kit) * 10) / 100
}

func ValorComissao(kit string) float64 {
	return (toFloat(kit) * 0.05) / 100
}

func CashFormat(v float64) string {
	cash := fmt.Sprintf("%.2f", v)
	return strings.Replace(cash, ".", ",", 1)
}