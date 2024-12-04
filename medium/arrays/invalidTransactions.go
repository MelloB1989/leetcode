package main

import (
	"fmt"
	"strconv"
	"strings"
)

func invalidTransactions(transactions []string) []string {
	trxs := make(map[string][]string)
	invalid := []string{}
	for _, trx := range transactions {
		trx_destructured := makeTransactions(trx)
		trxs[trx_destructured.Name] = append(trxs[trx_destructured.Name], trx)
	}

	for _, trx := range transactions {
		dt := makeTransactions(trx)
		if !isValid(trxs[dt.Name], trx) {
			invalid = append(invalid, trx)
		}
	}
	return invalid
}

func isValid(prev_trx []string, cur_trx string) bool {
	ct := makeTransactions(cur_trx)
	if ct.Amt > 1000 {
		return false
	}

	for _, trx := range prev_trx {
		it := makeTransactions(trx)
		if abs(ct.Time-it.Time) <= 60 && ct.City != it.City {
			return false
		}
	}
	return true
}

type Trx struct {
	Name string
	Time int
	Amt  int
	City string
}

func makeTransactions(trx string) Trx {
	t := strings.Split(trx, ",")
	ti, _ := strconv.Atoi(t[1])
	amt, _ := strconv.Atoi(t[2])
	tr := Trx{
		Name: t[0],
		Time: ti,
		Amt:  amt,
		City: t[3],
	}
	return tr
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	transactions := []string{"alice,20,800,mtv", "alice,50,100,mtv", "alice,51,100,frankfurt"}
	//["alice,20,800,mtv","alice,50,100,mtv","alice,51,100,frankfurt"]
	result := invalidTransactions(transactions)
	fmt.Println(result)
}
