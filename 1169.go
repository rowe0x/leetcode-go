package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

func invalidTransactions(transactions []string) []string {
	m := make(map[string][]string)
	for _, transaction := range transactions {
		name := strings.SplitN(transaction, ",", 1)[0]
		m[name] = append(m[name], transaction)
	}
	for k := range m {
		sort.Slice(m[k], func(i, j int) bool {
			iTime := strings.SplitN(m[k][i], ",", 3)[1]
			jTime := strings.SplitN(m[k][j], ",", 3)[1]
			return iTime < jTime
		})
	}
	var res []string
	for _, v := range m {
		count := len(v)
		if count == 1 {
			amount := strings.Split(v[0], ",")[2]
			if i, _ := strconv.Atoi(amount); i > 1000 {
				res = append(res, v[0])
			}
		} else {
			preTransaction := v[0]
			t := strings.Split(preTransaction, ",")
			preTransactionTime, _ := strconv.Atoi(t[2])
			preTransactionCity := t[3]
			var preTransactionValid = true
			var transactionTime int
			var transactionCity string
			var transactionAmount int
			for _, transaction := range transactions {
				t := strings.Split(transaction, ",")
				transactionAmount, _ = strconv.Atoi(t[1])
				transactionTime, _ = strconv.Atoi(t[2])
				transactionCity = t[3]
				if transactionCity != preTransactionCity && transactionTime - preTransactionTime <= 30 {
					if preTransactionValid {
						res = append(res, preTransaction)
					}
					preTransactionValid = false
					res = append(res, transaction)
				} else if transactionAmount > 1000 {
					preTransactionValid = false
					res = append(res, transaction)
				}
				preTransaction = transaction
				preTransactionTime = transactionTime
				preTransactionCity = transactionCity
			}
		}
	}
	return res
}
